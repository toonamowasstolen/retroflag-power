package app

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
	"github.com/toonamowasstolen/retroflag-power/internal/events"
	"github.com/toonamowasstolen/retroflag-power/internal/executor"
	"github.com/toonamowasstolen/retroflag-power/internal/planner"
	"github.com/toonamowasstolen/retroflag-power/internal/power"
	"github.com/toonamowasstolen/retroflag-power/internal/status"
)

func TestNewHasPlanner(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)

	got := New(logger, config.Default()).Planner()

	if got == nil {
		t.Fatal("Planner() = nil, want app-owned planner")
	}

	if plan := got.NewDryRunPlan("app lifecycle"); plan.Action != planner.ActionNoop {
		t.Fatalf("Planner().NewDryRunPlan() Action = %q, want %q", plan.Action, planner.ActionNoop)
	}
}

func TestNewStartsWithStartingStatus(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)

	got := New(logger, config.Default()).Status()
	want := status.New(config.Default(), status.StateStarting)

	if got != want {
		t.Fatalf("Status() = %#v, want %#v", got, want)
	}
}

func TestNewHasNoPreparedPlan(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)

	if plan, ok := New(logger, config.Default()).Plan(); ok {
		t.Fatalf("Plan() = %#v, true; want no prepared plan", plan)
	}
}

func TestNewHasNoExecutionSummary(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)

	if summary, ok := New(logger, config.Default()).ExecutionSummary(); ok {
		t.Fatalf("ExecutionSummary() = %#v, true; want no execution summary", summary)
	}
}

func TestNewHasNoExecutionStatus(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)

	got := New(logger, config.Default()).ExecutionStatus()
	want := ExecutionStatus{}

	if got != want {
		t.Fatalf("ExecutionStatus() = %#v, want %#v", got, want)
	}
}

func TestNewHasNoRuntimeSnapshotSummaries(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)

	got := New(logger, config.Default()).RuntimeSnapshot()
	want := RuntimeSnapshot{
		Status: status.New(config.Default(), status.StateStarting),
	}

	if got != want {
		t.Fatalf("RuntimeSnapshot() = %#v, want %#v", got, want)
	}
}

func TestRuntimeSnapshotSummaryBeforeStartup(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)

	app := New(logger, config.Default())
	got := app.RuntimeSnapshot().Summary()
	want := RuntimeSnapshotSummary{
		State: status.StateStarting,
	}

	if got != want {
		t.Fatalf("RuntimeSnapshot().Summary() before startup = %#v, want %#v", got, want)
	}

	const wantString = "state=starting plan_present=false execution_complete=false execution_success=false execution_error_captured=false dry_run_noop_only=false"
	if gotString := got.String(); gotString != wantString {
		t.Fatalf("RuntimeSnapshot().Summary().String() before startup = %q, want %q", gotString, wantString)
	}
}

func TestRuntimeSummaryBeforeStartupMatchesRuntimeSnapshotSummary(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())

	got := app.RuntimeSummary()
	want := app.RuntimeSnapshot().Summary()

	if got != want {
		t.Fatalf("RuntimeSummary() before startup = %#v, want RuntimeSnapshot().Summary() %#v", got, want)
	}
}

func TestRuntimeDiagnosticBeforeStartupMatchesRuntimeSummary(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())

	got := app.RuntimeDiagnostic()
	wantSummary := app.RuntimeSummary()

	if got.Summary != wantSummary {
		t.Fatalf("RuntimeDiagnostic().Summary before startup = %#v, want RuntimeSummary() %#v", got.Summary, wantSummary)
	}
	if gotString, wantString := got.String(), wantSummary.String(); gotString != wantString {
		t.Fatalf("RuntimeDiagnostic().String() before startup = %q, want RuntimeSummary().String() %q", gotString, wantString)
	}
}

func TestStartupDiagnosticBeforeStartupIsUnavailable(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())

	if diagnostic, ok := app.StartupDiagnostic(); ok {
		t.Fatalf("StartupDiagnostic() before startup = %#v, true; want no startup diagnostic", diagnostic)
	}
}

func TestStartupResultBeforeStartupIsZero(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())

	got := app.StartupResult()
	want := StartupResult{}

	if got != want {
		t.Fatalf("StartupResult() before startup = %#v, want %#v", got, want)
	}
	if app.StartupSucceeded() != got.Succeeded {
		t.Fatalf("StartupSucceeded() before startup = %t, want StartupResult().Succeeded %t", app.StartupSucceeded(), got.Succeeded)
	}
}

func TestStartupSucceededBeforeStartupIsFalse(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())

	if app.StartupSucceeded() {
		t.Fatal("StartupSucceeded() before startup = true, want false")
	}
}

func TestRunLogsLifecycle(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	New(logger, config.Default()).Run(ctx)

	const want = `retroflag-powerd 0.1.0-dev starting dry_run=true
retroflag-powerd ready
shutdown signal received
retroflag-powerd stopped
`
	if got := output.String(); got != want {
		t.Fatalf("Run() logs:\n%q\nwant:\n%q", got, want)
	}
}

func TestRunStopsWithStoppedStatus(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	app := New(logger, config.Default())
	app.Run(ctx)

	got := app.Status()
	want := status.New(config.Default(), status.StateStopped)

	if got != want {
		t.Fatalf("Status() = %#v, want %#v", got, want)
	}
}

func TestRuntimeSnapshotSummaryAfterStartup(t *testing.T) {
	logged := make(chan string)
	checked := make(chan struct{})
	logger := log.New(&checkingWriter{logged: logged, checked: checked}, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	app := New(logger, config.Default())
	done := make(chan struct{})

	go func() {
		app.Run(ctx)
		close(done)
	}()

	assertLogAndStatus(t, logged, checked, "starting dry_run=true", status.StateStarting, app)
	assertLogAndStatus(t, logged, checked, "ready", status.StateReady, app)

	got := app.RuntimeSnapshot().Summary()
	want := RuntimeSnapshotSummary{
		State:                  status.StateReady,
		HasPlan:                true,
		ExecutionComplete:      true,
		ExecutionSucceeded:     true,
		ExecutionErrorCaptured: false,
		DryRunNoopOnly:         true,
	}
	if got != want {
		t.Fatalf("RuntimeSnapshot().Summary() after startup = %#v, want %#v", got, want)
	}
	const wantString = "state=ready plan_present=true execution_complete=true execution_success=true execution_error_captured=false dry_run_noop_only=true"
	if gotString := got.String(); gotString != wantString {
		t.Fatalf("RuntimeSnapshot().Summary().String() after startup = %q, want %q", gotString, wantString)
	}

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done
}

func TestStartupResultAfterStartupIsSuccessful(t *testing.T) {
	logged := make(chan string)
	checked := make(chan struct{})
	logger := log.New(&checkingWriter{logged: logged, checked: checked}, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	app := New(logger, config.Default())
	done := make(chan struct{})

	go func() {
		app.Run(ctx)
		close(done)
	}()

	assertLogAndStatus(t, logged, checked, "starting dry_run=true", status.StateStarting, app)
	assertLogAndStatus(t, logged, checked, "ready", status.StateReady, app)

	got := app.StartupResult()
	want := StartupResult{
		Completed: true,
		Succeeded: true,
	}
	if got != want {
		t.Fatalf("StartupResult() after startup = %#v, want %#v", got, want)
	}
	if app.StartupSucceeded() != got.Succeeded {
		t.Fatalf("StartupSucceeded() after startup = %t, want StartupResult().Succeeded %t", app.StartupSucceeded(), got.Succeeded)
	}

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done
}

func TestStartupDiagnosticAfterStartupMatchesRuntimeDiagnostic(t *testing.T) {
	logged := make(chan string)
	checked := make(chan struct{})
	logger := log.New(&checkingWriter{logged: logged, checked: checked}, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	app := New(logger, config.Default())
	done := make(chan struct{})

	go func() {
		app.Run(ctx)
		close(done)
	}()

	assertLogAndStatus(t, logged, checked, "starting dry_run=true", status.StateStarting, app)
	assertLogAndStatus(t, logged, checked, "ready", status.StateReady, app)

	got, ok := app.StartupDiagnostic()
	if !ok {
		t.Fatal("StartupDiagnostic() after startup reports no startup diagnostic")
	}
	want := app.RuntimeDiagnostic()
	if got != want {
		t.Fatalf("StartupDiagnostic() after startup = %#v, want RuntimeDiagnostic() %#v", got, want)
	}

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done
}

func TestStartupSucceededAfterStartupIsTrue(t *testing.T) {
	logged := make(chan string)
	checked := make(chan struct{})
	logger := log.New(&checkingWriter{logged: logged, checked: checked}, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	app := New(logger, config.Default())
	done := make(chan struct{})

	go func() {
		app.Run(ctx)
		close(done)
	}()

	assertLogAndStatus(t, logged, checked, "starting dry_run=true", status.StateStarting, app)
	assertLogAndStatus(t, logged, checked, "ready", status.StateReady, app)

	if !app.StartupSucceeded() {
		t.Fatal("StartupSucceeded() after startup = false, want true")
	}

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done
}

func TestRuntimeSummaryAfterStartupMatchesRuntimeSnapshotSummary(t *testing.T) {
	logged := make(chan string)
	checked := make(chan struct{})
	logger := log.New(&checkingWriter{logged: logged, checked: checked}, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	app := New(logger, config.Default())
	done := make(chan struct{})

	go func() {
		app.Run(ctx)
		close(done)
	}()

	assertLogAndStatus(t, logged, checked, "starting dry_run=true", status.StateStarting, app)
	assertLogAndStatus(t, logged, checked, "ready", status.StateReady, app)

	got := app.RuntimeSummary()
	want := app.RuntimeSnapshot().Summary()

	if got != want {
		t.Fatalf("RuntimeSummary() after startup = %#v, want RuntimeSnapshot().Summary() %#v", got, want)
	}

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done
}

func TestRuntimeDiagnosticAfterStartupMatchesRuntimeSummary(t *testing.T) {
	logged := make(chan string)
	checked := make(chan struct{})
	logger := log.New(&checkingWriter{logged: logged, checked: checked}, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	app := New(logger, config.Default())
	done := make(chan struct{})

	go func() {
		app.Run(ctx)
		close(done)
	}()

	assertLogAndStatus(t, logged, checked, "starting dry_run=true", status.StateStarting, app)
	assertLogAndStatus(t, logged, checked, "ready", status.StateReady, app)

	got := app.RuntimeDiagnostic()
	wantSummary := app.RuntimeSummary()

	if got.Summary != wantSummary {
		t.Fatalf("RuntimeDiagnostic().Summary after startup = %#v, want RuntimeSummary() %#v", got.Summary, wantSummary)
	}
	if gotString, wantString := got.String(), wantSummary.String(); gotString != wantString {
		t.Fatalf("RuntimeDiagnostic().String() after startup = %q, want RuntimeSummary().String() %q", gotString, wantString)
	}

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done
}

func TestRuntimeSnapshotSummaryAfterShutdown(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	app := New(logger, config.Default())
	app.Run(ctx)

	got := app.RuntimeSnapshot().Summary()
	want := RuntimeSnapshotSummary{
		State:                  status.StateStopped,
		HasPlan:                true,
		ExecutionComplete:      true,
		ExecutionSucceeded:     true,
		ExecutionErrorCaptured: false,
		DryRunNoopOnly:         true,
	}
	if got != want {
		t.Fatalf("RuntimeSnapshot().Summary() after shutdown = %#v, want %#v", got, want)
	}

	const wantString = "state=stopped plan_present=true execution_complete=true execution_success=true execution_error_captured=false dry_run_noop_only=true"
	if gotString := got.String(); gotString != wantString {
		t.Fatalf("RuntimeSnapshot().Summary().String() after shutdown = %q, want %q", gotString, wantString)
	}
}

func TestRuntimeSummaryAfterShutdownMatchesRuntimeSnapshotSummary(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	app := New(logger, config.Default())
	app.Run(ctx)

	got := app.RuntimeSummary()
	want := app.RuntimeSnapshot().Summary()

	if got != want {
		t.Fatalf("RuntimeSummary() after shutdown = %#v, want RuntimeSnapshot().Summary() %#v", got, want)
	}
}

func TestRuntimeDiagnosticAfterShutdownMatchesRuntimeSummary(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	app := New(logger, config.Default())
	app.Run(ctx)

	got := app.RuntimeDiagnostic()
	wantSummary := app.RuntimeSummary()

	if got.Summary != wantSummary {
		t.Fatalf("RuntimeDiagnostic().Summary after shutdown = %#v, want RuntimeSummary() %#v", got.Summary, wantSummary)
	}
	if got.Summary.State != status.StateStopped {
		t.Fatalf("RuntimeDiagnostic().Summary.State after shutdown = %q, want %q", got.Summary.State, status.StateStopped)
	}
	if gotString, wantString := got.String(), wantSummary.String(); gotString != wantString {
		t.Fatalf("RuntimeDiagnostic().String() after shutdown = %q, want RuntimeSummary().String() %q", gotString, wantString)
	}
}

func TestStartupDiagnosticAfterShutdownKeepsStartupSnapshot(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	app := New(logger, config.Default())
	app.Run(ctx)

	got, ok := app.StartupDiagnostic()
	if !ok {
		t.Fatal("StartupDiagnostic() after shutdown reports no startup diagnostic")
	}

	want := RuntimeDiagnostic{
		Summary: RuntimeSnapshotSummary{
			State:                  status.StateReady,
			HasPlan:                true,
			ExecutionComplete:      true,
			ExecutionSucceeded:     true,
			ExecutionErrorCaptured: false,
			DryRunNoopOnly:         true,
		},
	}
	if got != want {
		t.Fatalf("StartupDiagnostic() after shutdown = %#v, want startup snapshot %#v", got, want)
	}

	current := app.RuntimeDiagnostic()
	if current.Summary.State != status.StateStopped {
		t.Fatalf("RuntimeDiagnostic().Summary.State after shutdown = %q, want %q", current.Summary.State, status.StateStopped)
	}
	if got == current {
		t.Fatalf("StartupDiagnostic() after shutdown = current RuntimeDiagnostic() %#v; want preserved startup snapshot", got)
	}
}

func TestStartupResultAfterShutdownRemainsSuccessful(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	app := New(logger, config.Default())
	app.Run(ctx)

	got := app.StartupResult()
	want := StartupResult{
		Completed: true,
		Succeeded: true,
	}

	if got != want {
		t.Fatalf("StartupResult() after shutdown = %#v, want %#v", got, want)
	}
	if app.StartupSucceeded() != got.Succeeded {
		t.Fatalf("StartupSucceeded() after shutdown = %t, want StartupResult().Succeeded %t", app.StartupSucceeded(), got.Succeeded)
	}
}

func TestStartupSucceededAfterShutdownRemainsTrue(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	app := New(logger, config.Default())
	app.Run(ctx)

	if !app.StartupSucceeded() {
		t.Fatal("StartupSucceeded() after shutdown = false, want true")
	}
}

func TestRunPreparesDryRunPlanAndReachesLifecycleStatuses(t *testing.T) {
	logged := make(chan string)
	checked := make(chan struct{})
	logger := log.New(&checkingWriter{logged: logged, checked: checked}, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	app := New(logger, config.Default())
	done := make(chan struct{})

	go func() {
		app.Run(ctx)
		close(done)
	}()

	assertLogAndStatus(t, logged, checked, "starting dry_run=true", status.StateStarting, app)
	assertLogAndStatus(t, logged, checked, "ready", status.StateReady, app)

	plan, ok := app.Plan()
	if !ok {
		t.Fatal("Plan() reports no prepared plan after startup")
	}
	summary := plan.Summary()
	wantSummary := planner.PlanSummary{
		DryRun:      true,
		ActionCount: 1,
		NoopOnly:    true,
	}
	if summary != wantSummary {
		t.Fatalf("Plan().Summary() = %#v, want %#v", summary, wantSummary)
	}
	if plan.Reason == "" {
		t.Fatal("Plan().Reason is empty, want startup reason")
	}

	executionSummary, ok := app.ExecutionSummary()
	if !ok {
		t.Fatal("ExecutionSummary() reports no execution after startup")
	}
	wantExecutionSummary := executor.ResultSummary{
		DryRun:         true,
		NoopOnly:       true,
		ActionsHandled: 1,
		Succeeded:      true,
	}
	if executionSummary != wantExecutionSummary {
		t.Fatalf("ExecutionSummary() = %#v, want %#v", executionSummary, wantExecutionSummary)
	}

	executionStatus := app.ExecutionStatus()
	wantExecutionStatus := ExecutionStatus{
		Completed:     true,
		ErrorCaptured: false,
		ErrorMessage:  "",
	}
	if executionStatus != wantExecutionStatus {
		t.Fatalf("ExecutionStatus() = %#v, want %#v", executionStatus, wantExecutionStatus)
	}

	readySnapshot := app.RuntimeSnapshot()
	wantReadySnapshot := RuntimeSnapshot{
		Status:              status.New(config.Default(), status.StateReady),
		HasPlanSummary:      true,
		PlanSummary:         wantSummary,
		ExecutionStatus:     wantExecutionStatus,
		HasExecutionSummary: true,
		ExecutionSummary:    wantExecutionSummary,
	}
	if readySnapshot != wantReadySnapshot {
		t.Fatalf("RuntimeSnapshot() after startup = %#v, want %#v", readySnapshot, wantReadySnapshot)
	}

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done

	stoppedSnapshot := app.RuntimeSnapshot()
	wantStoppedSnapshot := wantReadySnapshot
	wantStoppedSnapshot.Status = status.New(config.Default(), status.StateStopped)
	if stoppedSnapshot != wantStoppedSnapshot {
		t.Fatalf("RuntimeSnapshot() after shutdown = %#v, want %#v", stoppedSnapshot, wantStoppedSnapshot)
	}
}

func TestProcessPowerIntentPlansAndExecutesDryRunNoop(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())

	got, err := app.ProcessPowerIntent(power.IntentPowerButtonPressed)
	if err != nil {
		t.Fatalf("ProcessPowerIntent() error = %v, want nil", err)
	}

	wantResult := executor.Result{
		DryRun:         true,
		NoopOnly:       true,
		ActionsHandled: 1,
		Succeeded:      true,
	}
	if got != wantResult {
		t.Fatalf("ProcessPowerIntent() result = %#v, want %#v", got, wantResult)
	}

	plan, ok := app.Plan()
	if !ok {
		t.Fatal("Plan() reports no prepared plan after power intent")
	}
	wantPlan := planner.Plan{
		Action:      planner.ActionNoop,
		Reason:      "dry-run power intent: power_button_pressed",
		PowerIntent: power.IntentPowerButtonPressed,
	}
	if plan.Action != wantPlan.Action || plan.Reason != wantPlan.Reason || plan.PowerIntent != wantPlan.PowerIntent {
		t.Fatalf("Plan() = %#v, want deterministic dry-run power intent plan matching %#v", plan, wantPlan)
	}

	summary := plan.Summary()
	wantPlanSummary := planner.PlanSummary{
		DryRun:      true,
		ActionCount: 1,
		NoopOnly:    true,
	}
	if summary != wantPlanSummary {
		t.Fatalf("Plan().Summary() = %#v, want %#v", summary, wantPlanSummary)
	}

	executionSummary, ok := app.ExecutionSummary()
	if !ok {
		t.Fatal("ExecutionSummary() reports no execution after power intent")
	}
	wantExecutionSummary := executor.ResultSummary{
		DryRun:         true,
		NoopOnly:       true,
		ActionsHandled: 1,
		Succeeded:      true,
	}
	if executionSummary != wantExecutionSummary {
		t.Fatalf("ExecutionSummary() = %#v, want %#v", executionSummary, wantExecutionSummary)
	}

	wantEvents := []events.Event{
		{
			Type:    events.TypePowerIntentReceived,
			Message: "power intent received intent=power_button_pressed",
		},
		{
			Type:    events.TypeDryRunPlanPrepared,
			Message: "dry-run plan prepared intent=power_button_pressed action=noop",
		},
		{
			Type:    events.TypeNoopExecutionCompleted,
			Message: "noop execution completed intent=power_button_pressed actions_handled=1",
		},
	}
	if gotEvents := app.Events(); !equalEvents(gotEvents, wantEvents) {
		t.Fatalf("Events() after power intent = %#v, want %#v", gotEvents, wantEvents)
	}
}

func TestProcessPowerIntentUsesExplicitNoopPolicy(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	cfg := config.Default()
	cfg.PowerButtonAction = config.PowerButtonActionNoop
	app := New(logger, cfg)

	got, err := app.ProcessPowerIntent(power.IntentPowerButtonPressed)
	if err != nil {
		t.Fatalf("ProcessPowerIntent() error = %v, want nil", err)
	}

	if !got.DryRun || !got.NoopOnly || !got.Succeeded || got.ActionsHandled != 1 {
		t.Fatalf("ProcessPowerIntent() result = %#v, want successful dry-run noop-only result", got)
	}

	plan, ok := app.Plan()
	if !ok {
		t.Fatal("Plan() reports no prepared plan after explicit noop policy")
	}
	if plan.Action != planner.ActionNoop {
		t.Fatalf("Plan().Action = %q, want %q", plan.Action, planner.ActionNoop)
	}
}

func TestProcessPowerIntentRejectsUnsupportedPolicyClearly(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	cfg := config.Default()
	cfg.PowerButtonAction = "shutdown"
	app := New(logger, cfg)

	got, err := app.ProcessPowerIntent(power.IntentPowerButtonPressed)
	if err == nil {
		t.Fatal("ProcessPowerIntent() error = nil, want unsupported policy error")
	}

	const wantError = `unsupported power_button_action "shutdown" (supported: noop)`
	if err.Error() != wantError {
		t.Fatalf("ProcessPowerIntent() error = %q, want %q", err.Error(), wantError)
	}

	if got != (executor.Result{}) {
		t.Fatalf("ProcessPowerIntent() result = %#v, want zero result with no executor action", got)
	}

	if _, ok := app.Plan(); ok {
		t.Fatal("Plan() reports prepared plan for unsupported policy, want no plan")
	}

	gotExecutionStatus := app.ExecutionStatus()
	wantExecutionStatus := ExecutionStatus{
		Completed:     true,
		ErrorCaptured: true,
		ErrorMessage:  wantError,
	}
	if gotExecutionStatus != wantExecutionStatus {
		t.Fatalf("ExecutionStatus() = %#v, want %#v", gotExecutionStatus, wantExecutionStatus)
	}

	wantEvents := []events.Event{
		{
			Type:    events.TypePowerIntentReceived,
			Message: "power intent received intent=power_button_pressed",
		},
	}
	if gotEvents := app.Events(); !equalEvents(gotEvents, wantEvents) {
		t.Fatalf("Events() after unsupported policy = %#v, want %#v", gotEvents, wantEvents)
	}
}

func TestPowerIntentEventsAreDeterministic(t *testing.T) {
	var firstOutput bytes.Buffer
	first := New(log.New(&firstOutput, "", 0), config.Default())
	if _, err := first.ProcessPowerIntent(power.IntentPowerButtonPressed); err != nil {
		t.Fatalf("first ProcessPowerIntent() error = %v, want nil", err)
	}

	var secondOutput bytes.Buffer
	second := New(log.New(&secondOutput, "", 0), config.Default())
	if _, err := second.ProcessPowerIntent(power.IntentPowerButtonPressed); err != nil {
		t.Fatalf("second ProcessPowerIntent() error = %v, want nil", err)
	}

	if got, want := second.Events(), first.Events(); !equalEvents(got, want) {
		t.Fatalf("second Events() = %#v, want first Events() %#v", got, want)
	}
}

func TestEventsReturnsSnapshot(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())
	if _, err := app.ProcessPowerIntent(power.IntentPowerButtonPressed); err != nil {
		t.Fatalf("ProcessPowerIntent() error = %v, want nil", err)
	}

	snapshot := app.Events()
	snapshot[0] = events.Event{
		Type:    "changed",
		Message: "changed",
	}

	got := app.Events()[0]
	want := events.Event{
		Type:    events.TypePowerIntentReceived,
		Message: "power intent received intent=power_button_pressed",
	}
	if got != want {
		t.Fatalf("Events()[0] = %#v after snapshot mutation, want %#v", got, want)
	}
}

func TestExecutionStatusReportsCapturedExecutionError(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	app := New(logger, config.Default())
	app.hasExecution = true
	app.executionErr = executor.ErrUnsupportedPlan

	got := app.ExecutionStatus()
	want := ExecutionStatus{
		Completed:     true,
		ErrorCaptured: true,
		ErrorMessage:  executor.ErrUnsupportedPlan.Error(),
	}

	if got != want {
		t.Fatalf("ExecutionStatus() = %#v, want %#v", got, want)
	}
}

func TestPlanReturnsSnapshot(t *testing.T) {
	var output bytes.Buffer
	logger := log.New(&output, "", 0)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	app := New(logger, config.Default())
	app.Run(ctx)

	snapshot, ok := app.Plan()
	if !ok {
		t.Fatal("Plan() reports no prepared plan after startup")
	}
	snapshot.Action = "changed"
	snapshot.Reason = "changed"

	got, ok := app.Plan()
	if !ok {
		t.Fatal("Plan() reports no prepared plan after snapshot mutation")
	}
	if got.Action != planner.ActionNoop {
		t.Fatalf("Plan().Action = %q after snapshot mutation, want %q", got.Action, planner.ActionNoop)
	}
	if got.Reason != "daemon startup" {
		t.Fatalf("Plan().Reason = %q after snapshot mutation, want %q", got.Reason, "daemon startup")
	}
}

type checkingWriter struct {
	logged  chan<- string
	checked <-chan struct{}
}

func (w *checkingWriter) Write(p []byte) (int, error) {
	w.logged <- string(p)
	<-w.checked
	return len(p), nil
}

func assertLogAndStatus(
	t *testing.T,
	logged <-chan string,
	checked chan<- struct{},
	logPart string,
	wantState status.State,
	app *App,
) {
	t.Helper()

	if got := <-logged; !strings.Contains(got, logPart) {
		t.Fatalf("log = %q, want it to contain %q", got, logPart)
	}
	if got := app.Status().State; got != wantState {
		t.Fatalf("Status().State = %q after %q log, want %q", got, logPart, wantState)
	}
	checked <- struct{}{}
}

func equalEvents(a []events.Event, b []events.Event) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
