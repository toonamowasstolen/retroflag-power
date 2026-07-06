package status

import "github.com/toonamowasstolen/retroflag-power/internal/config"

type State string

const (
	StateStarting State = "starting"
	StateReady    State = "ready"
	StateStopping State = "stopping"
	StateStopped  State = "stopped"
)

type Status struct {
	AppName string
	Version string
	DryRun  bool
	State   State
}

func New(cfg config.Config, state State) Status {
	return Status{
		AppName: cfg.AppName,
		Version: cfg.Version,
		DryRun:  cfg.DryRun,
		State:   state,
	}
}
