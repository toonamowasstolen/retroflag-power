package app

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
	"github.com/toonamowasstolen/retroflag-power/internal/executor"
	"github.com/toonamowasstolen/retroflag-power/internal/planner"
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

	cancel()
	assertLogAndStatus(t, logged, checked, "shutdown signal received", status.StateStopping, app)
	assertLogAndStatus(t, logged, checked, "stopped", status.StateStopped, app)
	<-done
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
