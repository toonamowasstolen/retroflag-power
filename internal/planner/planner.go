package planner

type Action string

const ActionNoop Action = "noop"

type Plan struct {
	Action Action
	Reason string
}

func NewDryRunPlan(reason string) Plan {
	return Plan{
		Action: ActionNoop,
		Reason: reason,
	}
}
