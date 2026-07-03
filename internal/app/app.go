package app

import (
	"context"
	"fmt"
	"log"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
	"github.com/toonamowasstolen/retroflag-power/internal/events"
)

type App struct {
	logger *log.Logger
	config config.Config
}

func New(logger *log.Logger, cfg config.Config) *App {
	return &App{
		logger: logger,
		config: cfg,
	}
}

func (a *App) Run(ctx context.Context) {
	a.logEvent(events.Event{
		Type:    events.TypeDaemonStarting,
		Message: fmt.Sprintf("%s %s starting dry_run=%t", a.config.AppName, a.config.Version, a.config.DryRun),
	})
	a.logEvent(events.Event{
		Type:    events.TypeDaemonReady,
		Message: fmt.Sprintf("%s ready", a.config.AppName),
	})

	<-ctx.Done()

	a.logEvent(events.Event{
		Type:    events.TypeShutdownSignalReceived,
		Message: "shutdown signal received",
	})
	a.logEvent(events.Event{
		Type:    events.TypeDaemonStopped,
		Message: fmt.Sprintf("%s stopped", a.config.AppName),
	})
}

func (a *App) logEvent(event events.Event) {
	a.logger.Println(event.Message)
}
