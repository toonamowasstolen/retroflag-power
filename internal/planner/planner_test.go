package planner

import "testing"

func TestPlannerCreatesDryRunPlan(t *testing.T) {
	const reason = "app-owned planner"

	plan := New().NewDryRunPlan(reason)

	if plan.Action != ActionNoop {
		t.Fatalf("NewDryRunPlan() Action = %q, want %q", plan.Action, ActionNoop)
	}

	if plan.Reason != reason {
		t.Fatalf("NewDryRunPlan() Reason = %q, want %q", plan.Reason, reason)
	}
}

func TestNewDryRunPlan(t *testing.T) {
	const reason = "dry-run planning"

	plan := NewDryRunPlan(reason)

	if plan.Action != ActionNoop {
		t.Fatalf("NewDryRunPlan() Action = %q, want %q", plan.Action, ActionNoop)
	}

	if plan.Reason != reason {
		t.Fatalf("NewDryRunPlan() Reason = %q, want %q", plan.Reason, reason)
	}
}

func TestDryRunPlanSummary(t *testing.T) {
	plan := NewDryRunPlan("dry-run planning")

	got := plan.Summary()
	want := PlanSummary{
		DryRun:      true,
		ActionCount: 1,
		NoopOnly:    true,
	}

	if got != want {
		t.Fatalf("Summary() = %#v, want %#v", got, want)
	}
}

func TestZeroPlanSummary(t *testing.T) {
	got := (Plan{}).Summary()
	want := PlanSummary{
		DryRun:      false,
		ActionCount: 0,
		NoopOnly:    false,
	}

	if got != want {
		t.Fatalf("Summary() = %#v, want %#v", got, want)
	}
}
