package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/toonamowasstolen/retroflag-power/internal/app"
	"github.com/toonamowasstolen/retroflag-power/internal/logging"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app.New(logging.New()).Run(ctx)
}
