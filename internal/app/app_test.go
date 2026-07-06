package app

import (
	"bytes"
	"context"
	"log"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
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
