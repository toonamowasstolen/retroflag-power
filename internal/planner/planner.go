package planner

type Action string

const ActionNoop Action = "noop"

type Plan struct {
	Action Action
	Reason string
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
	}
}
