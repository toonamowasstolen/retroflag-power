package app

import (
	"context"
	"log"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
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
	a.logger.Printf("%s %s starting dry_run=%t", a.config.AppName, a.config.Version, a.config.DryRun)
	a.logger.Printf("%s ready", a.config.AppName)

	<-ctx.Done()

	a.logger.Println("shutdown signal received")
	a.logger.Printf("%s stopped", a.config.AppName)
}
