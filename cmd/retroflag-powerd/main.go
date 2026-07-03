package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/toonamowasstolen/retroflag-power/internal/app"
	"github.com/toonamowasstolen/retroflag-power/internal/logging"
	"github.com/toonamowasstolen/retroflag-power/internal/version"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(version.String())
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app.New(logging.New()).Run(ctx)
}
