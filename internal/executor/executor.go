package executor

import (
	"errors"

	"github.com/toonamowasstolen/retroflag-power/internal/planner"
)

var ErrUnsupportedPlan = errors.New("unsupported plan")

type Executor struct{}

type Result struct {
	DryRun         bool
	NoopOnly       bool
	ActionsHandled int
}

func New() *Executor {
	return &Executor{}
}

func (e *Executor) Execute(plan planner.Plan) (Result, error) {
	summary := plan.Summary()
	result := Result{
		DryRun:   summary.DryRun,
		NoopOnly: summary.NoopOnly,
	}

	if !summary.DryRun || !summary.NoopOnly {
		return result, ErrUnsupportedPlan
	}

	result.ActionsHandled = summary.ActionCount
	return result, nil
}
