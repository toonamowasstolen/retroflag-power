package app

import (
	"context"
	"log"

	"github.com/toonamowasstolen/retroflag-power/internal/version"
)

type App struct {
	logger *log.Logger
}

func New(logger *log.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Run(ctx context.Context) {
	a.logger.Printf("%s starting", version.String())
	a.logger.Println("retroflag-powerd ready")

	<-ctx.Done()

	a.logger.Println("shutdown signal received")
	a.logger.Println("retroflag-powerd stopped")
}
