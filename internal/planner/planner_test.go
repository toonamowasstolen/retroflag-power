package planner

import "testing"

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
