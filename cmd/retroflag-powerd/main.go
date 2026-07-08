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
	"strconv"
	"sync"
	"syscall"

	"github.com/toonamowasstolen/retroflag-power/internal/app"
	"github.com/toonamowasstolen/retroflag-power/internal/config"
	"github.com/toonamowasstolen/retroflag-power/internal/diagnostics"
	"github.com/toonamowasstolen/retroflag-power/internal/events"
	"github.com/toonamowasstolen/retroflag-power/internal/gpio"
	"github.com/toonamowasstolen/retroflag-power/internal/input"
	"github.com/toonamowasstolen/retroflag-power/internal/power"
	"github.com/toonamowasstolen/retroflag-power/internal/version"
)

var probeGPIOSignal = gpio.ProbeSignal

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	os.Exit(run(ctx, os.Args[1:], os.Stdout, os.Stderr))
}

func run(ctx context.Context, args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) > 0 && args[0] == "diagnostics" {
		return runDiagnosticsCommand(args[1:], stdout, stderr)
	}

	cfg := config.Default()
	flags := flag.NewFlagSet("retroflag-powerd", flag.ContinueOnError)
	flags.SetOutput(stderr)

	showVersion := flags.Bool("version", false, "print version and exit")
	dryRunPowerButton := flags.Bool("dry-run-power-button", false, "process the dry-run power button intent and exit")
	fakePowerButtonObserver := flags.Bool("fake-power-button-observer", false, "emit a fake power button observer event and exit")
	fakePowerSignal := flags.String("fake-power-signal", "", "interpret a fake raw power signal (low, high, unverified) and exit")
	probeGPIOSignal := flags.String("probe-gpio-signal", "", "read a candidate GPIO pin signal and exit")
	powerButtonAction := flags.String("power-button-action", cfg.PowerButtonAction, "dry-run power button action policy")
	powerSwitchActiveSignal := flags.String("power-switch-active-signal", string(cfg.LatchingPowerSwitch.ActiveSignal), "latching power switch active signal (low, high)")
	powerSwitchActiveState := flags.String("power-switch-active-state", string(cfg.LatchingPowerSwitch.ActiveSwitchState), "latching power switch active state (off, on)")

	if err := flags.Parse(args); err != nil {
		return 2
	}

	fakePowerSignalProvided := false
	probeGPIOSignalProvided := false
	flags.Visit(func(flag *flag.Flag) {
		switch flag.Name {
		case "fake-power-signal":
			fakePowerSignalProvided = true
		case "probe-gpio-signal":
			probeGPIOSignalProvided = true
		}
	})

	if *showVersion {
		fmt.Fprintln(stdout, version.String())
		return 0
	}

	cfg.PowerButtonAction = *powerButtonAction
	cfg.LatchingPowerSwitch.ActiveSignal = input.ActiveSignal(*powerSwitchActiveSignal)
	cfg.LatchingPowerSwitch.ActiveSwitchState = input.ActiveSwitchState(*powerSwitchActiveState)
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
	if fakePowerSignalProvided {
		signalState, err := parseFakePowerSignal(*fakePowerSignal)
		if err != nil {
			fmt.Fprintf(stderr, "fake power signal failed: %v\n", err)
			return 1
		}
		if err := runFakePowerSignal(ctx, cfg, signalState, stdout, stderr); err != nil {
			fmt.Fprintf(stderr, "fake power signal failed: %v\n", err)
			return 1
		}
		return 0
	}
	if probeGPIOSignalProvided {
		pin, err := parseGPIOPin(*probeGPIOSignal)
		if err != nil {
			fmt.Fprintf(stderr, "gpio signal probe failed: %v\n", err)
			return 1
		}
		runProbeGPIOSignal(ctx, pin, stdout)
		return 0
	}

	app.New(log.New(stderr, "", log.LstdFlags), cfg).Run(ctx)
	return 0
}

func runDiagnosticsCommand(args []string, stdout io.Writer, stderr io.Writer) int {
	if diagnosticsHelpRequested(args) {
		diagnostics.RenderUsage(stdout)
		return 0
	}

	flags := flag.NewFlagSet("retroflag-powerd diagnostics", flag.ContinueOnError)
	flags.SetOutput(stderr)

	format := flags.String("format", "text", "diagnostics output format (text, json)")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if flags.NArg() != 0 {
		fmt.Fprintf(stderr, "diagnostics failed: unexpected argument %q\n", flags.Arg(0))
		return 2
	}

	switch *format {
	case "text":
		diagnostics.RenderText(stdout)
	case "json":
		diagnostics.RenderJSON(stdout)
	default:
		fmt.Fprintf(stderr, "diagnostics failed: unsupported format %q (supported: text, json)\n", *format)
		return 2
	}

	return 0
}

func diagnosticsHelpRequested(args []string) bool {
	return len(args) == 1 && (args[0] == "--help" || args[0] == "-h")
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

func runFakePowerSignal(ctx context.Context, cfg config.Config, signalState input.SignalState, stdout io.Writer, stderr io.Writer) error {
	return runAppAndProcess(ctx, cfg, stderr, func(daemon *app.App) error {
		rawEvent := input.SignalEvent(cfg.PowerInputName, signalState)
		switchEvent, err := input.InterpretLatchingPowerSwitchEvent(rawEvent, cfg.LatchingPowerSwitch)
		if err != nil {
			return err
		}

		processed := false
		executionSucceeded := false
		dryRun := false
		noopOnly := false
		actionsHandled := 0

		if switchEvent.SwitchState == input.SwitchOff {
			result, processErr := daemon.ProcessInputEvent(switchEvent)
			summary := result.Summary()
			processed = processErr == nil
			executionSucceeded = summary.Succeeded
			dryRun = summary.DryRun
			noopOnly = summary.NoopOnly
			actionsHandled = summary.ActionsHandled
			err = processErr
		}

		fmt.Fprintf(
			stdout,
			"fake_power_signal raw=%s input=%s active_signal=%s active_switch_state=%s interpreted=%s processed=%t execution_success=%t dry_run=%t noop_only=%t actions_handled=%d real_shutdown=false hardware_action=false\n",
			rawEvent.SignalState,
			rawEvent.Name,
			cfg.LatchingPowerSwitch.ActiveSignal,
			cfg.LatchingPowerSwitch.ActiveSwitchState,
			switchEvent.SwitchState,
			processed,
			executionSucceeded,
			dryRun,
			noopOnly,
			actionsHandled,
		)
		printEventBreadcrumbs(stdout, daemon.Events())

		return err
	})
}

func parseFakePowerSignal(value string) (input.SignalState, error) {
	state := input.SignalState(value)
	if state.Valid() {
		return state, nil
	}

	return "", input.UnsupportedSignalStateError{State: state}
}

func runProbeGPIOSignal(ctx context.Context, pin int, stdout io.Writer) {
	signalState := probeGPIOSignal(ctx, pin)
	fmt.Fprintf(
		stdout,
		"gpio_signal_probe pin=%d raw=%s interpreted=false processed=false real_shutdown=false hardware_action=false\n",
		pin,
		signalState.Label(),
	)
}

func parseGPIOPin(value string) (int, error) {
	pin, err := strconv.Atoi(value)
	if err != nil || pin < 0 {
		return 0, fmt.Errorf("unsupported GPIO pin %q (expected non-negative integer)", value)
	}

	return pin, nil
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
