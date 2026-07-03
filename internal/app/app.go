package app

import (
	"context"
	"log"
)

type App struct {
	logger *log.Logger
}

func New(logger *log.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Run(ctx context.Context) {
	a.logger.Println("retroflag-powerd starting")
	a.logger.Println("retroflag-powerd ready")

	<-ctx.Done()

	a.logger.Println("shutdown signal received")
	a.logger.Println("retroflag-powerd stopped")
}
