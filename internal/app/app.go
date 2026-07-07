package app

import (
	"context"
	"fmt"
	"log"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
	"github.com/toonamowasstolen/retroflag-power/internal/events"
	"github.com/toonamowasstolen/retroflag-power/internal/executor"
	"github.com/toonamowasstolen/retroflag-power/internal/input"
	"github.com/toonamowasstolen/retroflag-power/internal/planner"
	"github.com/toonamowasstolen/retroflag-power/internal/power"
	"github.com/toonamowasstolen/retroflag-power/internal/status"
)

type App struct {
	logger               *log.Logger
	config               config.Config
	planner              *planner.Planner
	executor             *executor.Executor
	plan                 planner.Plan
	hasPlan              bool
	executionResult      executor.Result
	executionErr         error
	hasExecution         bool
	startupResult        StartupResult
	startupDiagnostic    RuntimeDiagnostic
	hasStartupDiagnostic bool
	status               status.Status
	events               []events.Event
}

type ExecutionStatus struct {
	Completed     bool
	ErrorCaptured bool
	ErrorMessage  string
}

type RuntimeSnapshot struct {
	Status              status.Status
	HasPlanSummary      bool
	PlanSummary         planner.PlanSummary
	ExecutionStatus     ExecutionStatus
	HasExecutionSummary bool
	ExecutionSummary    executor.ResultSummary
}

type RuntimeSnapshotSummary struct {
	State                  status.State
	HasPlan                bool
	ExecutionComplete      bool
	ExecutionSucceeded     bool
	ExecutionErrorCaptured bool
	DryRunNoopOnly         bool
}

type RuntimeDiagnostic struct {
	Summary RuntimeSnapshotSummary
}

type StartupResult struct {
	Completed bool
	Succeeded bool
}

func (s RuntimeSnapshotSummary) String() string {
	return fmt.Sprintf(
		"state=%s plan_present=%t execution_complete=%t execution_success=%t execution_error_captured=%t dry_run_noop_only=%t",
		s.State,
		s.HasPlan,
		s.ExecutionComplete,
		s.ExecutionSucceeded,
		s.ExecutionErrorCaptured,
		s.DryRunNoopOnly,
	)
}

func (d RuntimeDiagnostic) String() string {
	return d.Summary.String()
}

func New(logger *log.Logger, cfg config.Config) *App {
	return &App{
		logger:   logger,
		config:   cfg,
		planner:  planner.New(),
		executor: executor.New(),
		status:   status.New(cfg, status.StateStarting),
	}
}

func (a *App) Run(ctx context.Context) {
	a.setStatus(status.StateStarting)
	a.logEvent(events.Event{
		Type:    events.TypeDaemonStarting,
		Message: fmt.Sprintf("%s %s starting dry_run=%t", a.config.AppName, a.config.Version, a.config.DryRun),
	})

	a.plan = a.planner.NewDryRunPlan("daemon startup")
	a.hasPlan = true
	a.executionResult, a.executionErr = a.executor.Execute(a.plan)
	a.hasExecution = true

	a.setStatus(status.StateReady)
	a.startupResult = StartupResult{
		Completed: true,
		Succeeded: true,
	}
	a.startupDiagnostic = a.RuntimeDiagnostic()
	a.hasStartupDiagnostic = true
	a.logEvent(events.Event{
		Type:    events.TypeDaemonReady,
		Message: fmt.Sprintf("%s ready", a.config.AppName),
	})

	<-ctx.Done()

	a.setStatus(status.StateStopping)
	a.logEvent(events.Event{
		Type:    events.TypeShutdownSignalReceived,
		Message: "shutdown signal received",
	})

	a.setStatus(status.StateStopped)
	a.logEvent(events.Event{
		Type:    events.TypeDaemonStopped,
		Message: fmt.Sprintf("%s stopped", a.config.AppName),
	})
}

func (a *App) Status() status.Status {
	return a.status
}

func (a *App) Planner() *planner.Planner {
	return a.planner
}

func (a *App) Plan() (planner.Plan, bool) {
	return a.plan, a.hasPlan
}

func (a *App) PlanSummary() (planner.PlanSummary, bool) {
	if !a.hasPlan {
		return planner.PlanSummary{}, false
	}

	return a.plan.Summary(), true
}

func (a *App) ProcessPowerIntent(intent power.Intent) (executor.Result, error) {
	a.logEvent(events.Event{
		Type:    events.TypePowerIntentReceived,
		Message: fmt.Sprintf("power intent received intent=%s", intent),
	})

	if err := a.config.ValidatePowerButtonAction(); err != nil {
		a.executionResult = executor.Result{}
		a.executionErr = err
		a.hasExecution = true
		return a.executionResult, a.executionErr
	}

	a.plan = a.planner.NewDryRunPowerIntentPlan(intent, planner.Action(a.config.PowerButtonAction))
	a.hasPlan = true
	a.logEvent(events.Event{
		Type:    events.TypeDryRunPlanPrepared,
		Message: fmt.Sprintf("dry-run plan prepared intent=%s action=%s", intent, a.plan.Action),
	})

	a.executionResult, a.executionErr = a.executor.Execute(a.plan)
	a.hasExecution = true
	if a.executionErr == nil && a.executionResult.DryRun && a.executionResult.NoopOnly {
		a.logEvent(events.Event{
			Type:    events.TypeNoopExecutionCompleted,
			Message: fmt.Sprintf("noop execution completed intent=%s actions_handled=%d", intent, a.executionResult.ActionsHandled),
		})
	}

	return a.executionResult, a.executionErr
}

func (a *App) ProcessInputEvent(event input.Event) (executor.Result, error) {
	switch event.Type {
	case input.EventTypePowerButtonPressed:
		return a.ProcessPowerIntent(power.IntentPowerButtonPressed)
	default:
		return executor.Result{}, fmt.Errorf("unsupported input event %q", event.Type)
	}
}

func (a *App) ProcessNextInputEvent(ctx context.Context, observer input.Observer) (executor.Result, error) {
	event, err := observer.NextEvent(ctx)
	if err != nil {
		return executor.Result{}, err
	}

	return a.ProcessInputEvent(event)
}

func (a *App) Events() []events.Event {
	return append([]events.Event(nil), a.events...)
}

func (a *App) ExecutionSummary() (executor.ResultSummary, bool) {
	if !a.hasExecution {
		return executor.ResultSummary{}, false
	}

	return a.executionResult.Summary(), true
}

func (a *App) RuntimeSnapshot() RuntimeSnapshot {
	planSummary, hasPlanSummary := a.PlanSummary()
	executionSummary, hasExecutionSummary := a.ExecutionSummary()

	return RuntimeSnapshot{
		Status:              a.Status(),
		HasPlanSummary:      hasPlanSummary,
		PlanSummary:         planSummary,
		ExecutionStatus:     a.ExecutionStatus(),
		HasExecutionSummary: hasExecutionSummary,
		ExecutionSummary:    executionSummary,
	}
}

func (a *App) RuntimeSummary() RuntimeSnapshotSummary {
	return a.RuntimeSnapshot().Summary()
}

func (a *App) RuntimeDiagnostic() RuntimeDiagnostic {
	return RuntimeDiagnostic{
		Summary: a.RuntimeSummary(),
	}
}

func (a *App) StartupDiagnostic() (RuntimeDiagnostic, bool) {
	if !a.hasStartupDiagnostic {
		return RuntimeDiagnostic{}, false
	}

	return a.startupDiagnostic, true
}

func (a *App) StartupResult() StartupResult {
	return a.startupResult
}

func (a *App) StartupSucceeded() bool {
	return a.StartupResult().Succeeded
}

func (s RuntimeSnapshot) Summary() RuntimeSnapshotSummary {
	summary := RuntimeSnapshotSummary{
		State:                  s.Status.State,
		HasPlan:                s.HasPlanSummary,
		ExecutionComplete:      s.ExecutionStatus.Completed,
		ExecutionErrorCaptured: s.ExecutionStatus.ErrorCaptured,
	}

	if s.HasExecutionSummary {
		summary.ExecutionSucceeded = s.ExecutionSummary.Succeeded
	}

	if s.HasPlanSummary && s.HasExecutionSummary {
		summary.DryRunNoopOnly = s.PlanSummary.DryRun &&
			s.PlanSummary.NoopOnly &&
			s.ExecutionSummary.DryRun &&
			s.ExecutionSummary.NoopOnly
	}

	return summary
}

func (a *App) ExecutionStatus() ExecutionStatus {
	executionStatus := ExecutionStatus{
		Completed:     a.hasExecution,
		ErrorCaptured: a.executionErr != nil,
	}
	if a.executionErr != nil {
		executionStatus.ErrorMessage = a.executionErr.Error()
	}

	return executionStatus
}

func (a *App) setStatus(state status.State) {
	a.status = status.New(a.config, state)
}

func (a *App) logEvent(event events.Event) {
	a.events = append(a.events, event)
	a.logger.Println(event.Message)
}
