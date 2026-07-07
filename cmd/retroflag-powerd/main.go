package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/toonamowasstolen/retroflag-power/internal/app"
	"github.com/toonamowasstolen/retroflag-power/internal/config"
	"github.com/toonamowasstolen/retroflag-power/internal/events"
	"github.com/toonamowasstolen/retroflag-power/internal/input"
	"github.com/toonamowasstolen/retroflag-power/internal/power"
	"github.com/toonamowasstolen/retroflag-power/internal/version"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	os.Exit(run(ctx, os.Args[1:], os.Stdout, os.Stderr))
}

func run(ctx context.Context, args []string, stdout io.Writer, stderr io.Writer) int {
	cfg := config.Default()
	flags := flag.NewFlagSet("retroflag-powerd", flag.ContinueOnError)
	flags.SetOutput(stderr)

	showVersion := flags.Bool("version", false, "print version and exit")
	dryRunPowerButton := flags.Bool("dry-run-power-button", false, "process the dry-run power button intent and exit")
	fakePowerButtonObserver := flags.Bool("fake-power-button-observer", false, "emit a fake power button observer event and exit")
	powerButtonAction := flags.String("power-button-action", cfg.PowerButtonAction, "dry-run power button action policy")

	if err := flags.Parse(args); err != nil {
		return 2
	}

	if *showVersion {
		fmt.Fprintln(stdout, version.String())
		return 0
	}

	cfg.PowerButtonAction = *powerButtonAction
	if *dryRunPowerButton {
		if err := runDryRunPowerButton(ctx, cfg, stdout, stderr); err != nil {
			fmt.Fprintf(stderr, "dry-run power button failed: %v\n", err)
			return 1
		}
		return 0
	}
	if *fakePowerButtonObserver {
		if err := runFakePowerButtonObserver(ctx, cfg, stdout, stderr); err != nil {
			fmt.Fprintf(stderr, "fake power button observer failed: %v\n", err)
			return 1
		}
		return 0
	}

	app.New(log.New(stderr, "", log.LstdFlags), cfg).Run(ctx)
	return 0
}

func runDryRunPowerButton(ctx context.Context, cfg config.Config, stdout io.Writer, stderr io.Writer) error {
	return runAppAndProcess(ctx, cfg, stderr, func(daemon *app.App) error {
		result, err := daemon.ProcessPowerIntent(power.IntentPowerButtonPressed)
		summary := result.Summary()
		fmt.Fprintf(
			stdout,
			"dry_run_power_button intent=%s processed=%t execution_success=%t dry_run=%t noop_only=%t actions_handled=%d real_shutdown=false hardware_action=false\n",
			power.IntentPowerButtonPressed,
			err == nil,
			summary.Succeeded,
			summary.DryRun,
			summary.NoopOnly,
			summary.ActionsHandled,
		)

		return err
	})
}

func runFakePowerButtonObserver(ctx context.Context, cfg config.Config, stdout io.Writer, stderr io.Writer) error {
	return runAppAndProcess(ctx, cfg, stderr, func(daemon *app.App) error {
		result, err := daemon.ProcessNextInputEvent(ctx, input.NewFakePowerButtonObserver())
		summary := result.Summary()
		fmt.Fprintf(
			stdout,
			"fake_power_button_observer event=%s processed=%t execution_success=%t dry_run=%t noop_only=%t actions_handled=%d real_shutdown=false hardware_action=false\n",
			input.EventTypePowerButtonPressed,
			err == nil,
			summary.Succeeded,
			summary.DryRun,
			summary.NoopOnly,
			summary.ActionsHandled,
		)
		printEventBreadcrumbs(stdout, daemon.Events())

		return err
	})
}

func runAppAndProcess(ctx context.Context, cfg config.Config, stderr io.Writer, process func(*app.App) error) error {
	appCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	ready := make(chan struct{})
	done := make(chan struct{})
	logger := log.New(&readySignalWriter{
		dst:   stderr,
		ready: ready,
	}, "", log.LstdFlags)
	daemon := app.New(logger, cfg)

	go func() {
		daemon.Run(appCtx)
		close(done)
	}()

	select {
	case <-ready:
	case <-ctx.Done():
		cancel()
		<-done
		return ctx.Err()
	}

	err := process(daemon)

	cancel()
	<-done

	return err
}

func printEventBreadcrumbs(stdout io.Writer, appEvents []events.Event) {
	for i, event := range appEvents {
		fmt.Fprintf(stdout, "event_breadcrumb index=%d type=%s message=%q\n", i, event.Type, event.Message)
	}
}

type readySignalWriter struct {
	dst   io.Writer
	ready chan<- struct{}
	once  sync.Once
}

func (w *readySignalWriter) Write(p []byte) (int, error) {
	if bytes.Contains(p, []byte(" ready\n")) {
		w.once.Do(func() {
			close(w.ready)
		})
	}

	return w.dst.Write(p)
}
