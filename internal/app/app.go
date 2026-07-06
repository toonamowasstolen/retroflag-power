package app

import (
	"context"
	"fmt"
	"log"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
	"github.com/toonamowasstolen/retroflag-power/internal/events"
	"github.com/toonamowasstolen/retroflag-power/internal/planner"
	"github.com/toonamowasstolen/retroflag-power/internal/status"
)

type App struct {
	logger  *log.Logger
	config  config.Config
	planner *planner.Planner
	status  status.Status
}

func New(logger *log.Logger, cfg config.Config) *App {
	return &App{
		logger:  logger,
		config:  cfg,
		planner: planner.New(),
		status:  status.New(cfg, status.StateStarting),
	}
}

func (a *App) Run(ctx context.Context) {
	a.setStatus(status.StateStarting)
	a.logEvent(events.Event{
		Type:    events.TypeDaemonStarting,
		Message: fmt.Sprintf("%s %s starting dry_run=%t", a.config.AppName, a.config.Version, a.config.DryRun),
	})

	a.setStatus(status.StateReady)
	a.logEvent(events.Event{
		Type:    events.TypeDaemonReady,
		Message: fmt.Sprintf("%s ready", a.config.AppName),
	})

	<-ctx.Done()

	a.setStatus(status.StateStopping)
	a.logEvent(events.Event{
		Type:    events.TypeShutdownSignalReceived,
		Message: "shutdown signal received",
	})

	a.setStatus(status.StateStopped)
	a.logEvent(events.Event{
		Type:    events.TypeDaemonStopped,
		Message: fmt.Sprintf("%s stopped", a.config.AppName),
	})
}

func (a *App) Status() status.Status {
	return a.status
}

func (a *App) Planner() *planner.Planner {
	return a.planner
}

func (a *App) setStatus(state status.State) {
	a.status = status.New(a.config, state)
}

func (a *App) logEvent(event events.Event) {
	a.logger.Println(event.Message)
}
