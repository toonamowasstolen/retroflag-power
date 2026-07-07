package executor

import (
	"errors"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/planner"
	"github.com/toonamowasstolen/retroflag-power/internal/power"
)

func TestExecuteDryRunNoopPlan(t *testing.T) {
	plan := planner.NewDryRunPlan("executor test")

	got, err := New().Execute(plan)
	if err != nil {
		t.Fatalf("Execute() error = %v, want nil", err)
	}

	want := Result{
		DryRun:         true,
		NoopOnly:       true,
		ActionsHandled: 1,
		Succeeded:      true,
	}
	if got != want {
		t.Fatalf("Execute() = %#v, want %#v", got, want)
	}
}

func TestExecuteDryRunPowerIntentPlanRemainsNoopOnly(t *testing.T) {
	plan := planner.NewDryRunPowerIntentPlan(power.IntentPowerButtonPressed)

	got, err := New().Execute(plan)
	if err != nil {
		t.Fatalf("Execute() error = %v, want nil", err)
	}

	want := Result{
		DryRun:         true,
		NoopOnly:       true,
		ActionsHandled: 1,
		Succeeded:      true,
	}
	if got != want {
		t.Fatalf("Execute() = %#v, want %#v", got, want)
	}
}

func TestExecuteUnsupportedPlan(t *testing.T) {
	plan := planner.Plan{
		Action: "shutdown",
		Reason: "real action",
	}

	got, err := New().Execute(plan)
	if !errors.Is(err, ErrUnsupportedPlan) {
		t.Fatalf("Execute() error = %v, want %v", err, ErrUnsupportedPlan)
	}

	want := Result{
		DryRun:      false,
		NoopOnly:    false,
		Unsupported: true,
	}
	if got != want {
		t.Fatalf("Execute() = %#v, want %#v", got, want)
	}
}

func TestSuccessfulResultSummary(t *testing.T) {
	result := Result{
		DryRun:         true,
		NoopOnly:       true,
		ActionsHandled: 1,
		Succeeded:      true,
	}

	got := result.Summary()
	want := ResultSummary{
		DryRun:         true,
		NoopOnly:       true,
		ActionsHandled: 1,
		Succeeded:      true,
	}

	if got != want {
		t.Fatalf("Summary() = %#v, want %#v", got, want)
	}
}

func TestUnsupportedResultSummary(t *testing.T) {
	result := Result{
		Unsupported: true,
	}

	got := result.Summary()
	want := ResultSummary{
		Unsupported: true,
	}

	if got != want {
		t.Fatalf("Summary() = %#v, want %#v", got, want)
	}
}
