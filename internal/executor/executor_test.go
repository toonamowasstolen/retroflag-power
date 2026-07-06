package executor

import (
	"errors"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/planner"
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
		DryRun:   false,
		NoopOnly: false,
	}
	if got != want {
		t.Fatalf("Execute() = %#v, want %#v", got, want)
	}
}
