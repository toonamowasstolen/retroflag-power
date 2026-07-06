package planner

type Action string

const ActionNoop Action = "noop"

type Plan struct {
	Action Action
	Reason string
	dryRun bool
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

func NewDryRunPlan(reason string) Plan {
	return Plan{
		Action: ActionNoop,
		Reason: reason,
		dryRun: true,
	}
}
