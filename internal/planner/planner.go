package planner

import "github.com/toonamowasstolen/retroflag-power/internal/power"

type Action string

const ActionNoop Action = "noop"

type Plan struct {
	Action      Action
	Reason      string
	PowerIntent power.Intent
	dryRun      bool
}

type PlanSummary struct {
	DryRun      bool
	ActionCount int
	NoopOnly    bool
}

func (p Plan) Summary() PlanSummary {
	actionCount := 0
	if p.Action != "" {
		actionCount = 1
	}

	return PlanSummary{
		DryRun:      p.dryRun,
		ActionCount: actionCount,
		NoopOnly:    actionCount == 1 && p.Action == ActionNoop,
	}
}

type Planner struct{}

func New() *Planner {
	return &Planner{}
}

func (p *Planner) NewDryRunPlan(reason string) Plan {
	return NewDryRunPlan(reason)
}

func (p *Planner) NewDryRunPowerIntentPlan(intent power.Intent, action Action) Plan {
	return NewDryRunPowerIntentPlan(intent, action)
}

func NewDryRunPlan(reason string) Plan {
	return Plan{
		Action: ActionNoop,
		Reason: reason,
		dryRun: true,
	}
}

func NewDryRunPowerIntentPlan(intent power.Intent, action Action) Plan {
	return Plan{
		Action:      action,
		Reason:      "dry-run power intent: " + intent.String(),
		PowerIntent: intent,
		dryRun:      true,
	}
}
