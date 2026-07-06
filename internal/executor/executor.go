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
	Succeeded      bool
	Unsupported    bool
}

type ResultSummary struct {
	DryRun         bool
	NoopOnly       bool
	ActionsHandled int
	Succeeded      bool
	Unsupported    bool
}

func (r Result) Summary() ResultSummary {
	return ResultSummary{
		DryRun:         r.DryRun,
		NoopOnly:       r.NoopOnly,
		ActionsHandled: r.ActionsHandled,
		Succeeded:      r.Succeeded,
		Unsupported:    r.Unsupported,
	}
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
		result.Unsupported = true
		return result, ErrUnsupportedPlan
	}

	result.ActionsHandled = summary.ActionCount
	result.Succeeded = true
	return result, nil
}
