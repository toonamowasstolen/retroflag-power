package main

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/input"
)

func TestRunVersionPrintsVersion(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--version"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--version) exit = %d, want 0", got)
	}
	const want = "retroflag-powerd 0.1.0-dev\n"
	if stdout.String() != want {
		t.Fatalf("run(--version) stdout = %q, want %q", stdout.String(), want)
	}
	if stderr.String() != "" {
		t.Fatalf("run(--version) stderr = %q, want empty", stderr.String())
	}
}

func TestRunDryRunPowerButtonProcessesIntentAndExits(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--dry-run-power-button"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--dry-run-power-button) exit = %d, want 0; stderr = %q", got, stderr.String())
	}
	const want = "dry_run_power_button intent=power_button_pressed processed=true execution_success=true dry_run=true noop_only=true actions_handled=1 real_shutdown=false hardware_action=false\n"
	if stdout.String() != want {
		t.Fatalf("run(--dry-run-power-button) stdout = %q, want %q", stdout.String(), want)
	}

	for _, wantLog := range []string{
		"retroflag-powerd 0.1.0-dev starting dry_run=true",
		"retroflag-powerd ready",
		"shutdown signal received",
		"retroflag-powerd stopped",
	} {
		if !strings.Contains(stderr.String(), wantLog) {
			t.Fatalf("run(--dry-run-power-button) stderr = %q, want it to contain %q", stderr.String(), wantLog)
		}
	}
}

func TestRunDryRunPowerButtonUsesExplicitNoopPolicy(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--dry-run-power-button", "--power-button-action", "noop"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--dry-run-power-button --power-button-action noop) exit = %d, want 0; stderr = %q", got, stderr.String())
	}
	const want = "dry_run_power_button intent=power_button_pressed processed=true execution_success=true dry_run=true noop_only=true actions_handled=1 real_shutdown=false hardware_action=false\n"
	if stdout.String() != want {
		t.Fatalf("run(--dry-run-power-button --power-button-action noop) stdout = %q, want %q", stdout.String(), want)
	}
}

func TestRunDryRunPowerButtonRejectsUnsupportedPolicy(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--dry-run-power-button", "--power-button-action", "shutdown"}, &stdout, &stderr)

	if got != 1 {
		t.Fatalf("run(--dry-run-power-button --power-button-action shutdown) exit = %d, want 1", got)
	}
	if stdout.String() != "dry_run_power_button intent=power_button_pressed processed=false execution_success=false dry_run=false noop_only=false actions_handled=0 real_shutdown=false hardware_action=false\n" {
		t.Fatalf("run(--dry-run-power-button --power-button-action shutdown) stdout = %q, want deterministic failed dry-run line", stdout.String())
	}
	const wantError = `dry-run power button failed: unsupported power_button_action "shutdown" (supported: noop)`
	if !strings.Contains(stderr.String(), wantError) {
		t.Fatalf("run(--dry-run-power-button --power-button-action shutdown) stderr = %q, want it to contain %q", stderr.String(), wantError)
	}
}

func TestRunFakePowerButtonObserverProcessesEventAndExits(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--fake-power-button-observer"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--fake-power-button-observer) exit = %d, want 0; stderr = %q", got, stderr.String())
	}

	wantLines := []string{
		"fake_power_button_observer event=power_button_pressed processed=true execution_success=true dry_run=true noop_only=true actions_handled=1 real_shutdown=false hardware_action=false",
		`event_breadcrumb index=0 type=daemon.starting message="retroflag-powerd 0.1.0-dev starting dry_run=true"`,
		`event_breadcrumb index=1 type=daemon.ready message="retroflag-powerd ready"`,
		`event_breadcrumb index=2 type=power.intent_received message="power intent received intent=power_button_pressed"`,
		`event_breadcrumb index=3 type=power.dry_run_plan_prepared message="dry-run plan prepared intent=power_button_pressed action=noop"`,
		`event_breadcrumb index=4 type=power.noop_execution_completed message="noop execution completed intent=power_button_pressed actions_handled=1"`,
	}
	gotLines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(gotLines) != len(wantLines) {
		t.Fatalf("run(--fake-power-button-observer) stdout lines = %#v, want %#v", gotLines, wantLines)
	}
	for i, want := range wantLines {
		if gotLines[i] != want {
			t.Fatalf("run(--fake-power-button-observer) stdout line %d = %q, want %q", i, gotLines[i], want)
		}
	}

	for _, wantLog := range []string{
		"retroflag-powerd 0.1.0-dev starting dry_run=true",
		"retroflag-powerd ready",
		"shutdown signal received",
		"retroflag-powerd stopped",
	} {
		if !strings.Contains(stderr.String(), wantLog) {
			t.Fatalf("run(--fake-power-button-observer) stderr = %q, want it to contain %q", stderr.String(), wantLog)
		}
	}
}

func TestRunFakePowerButtonObserverRejectsUnsupportedPolicy(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--fake-power-button-observer", "--power-button-action", "shutdown"}, &stdout, &stderr)

	if got != 1 {
		t.Fatalf("run(--fake-power-button-observer --power-button-action shutdown) exit = %d, want 1", got)
	}
	wantLines := []string{
		"fake_power_button_observer event=power_button_pressed processed=false execution_success=false dry_run=false noop_only=false actions_handled=0 real_shutdown=false hardware_action=false",
		`event_breadcrumb index=0 type=daemon.starting message="retroflag-powerd 0.1.0-dev starting dry_run=true"`,
		`event_breadcrumb index=1 type=daemon.ready message="retroflag-powerd ready"`,
		`event_breadcrumb index=2 type=power.intent_received message="power intent received intent=power_button_pressed"`,
	}
	gotLines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(gotLines) != len(wantLines) {
		t.Fatalf("run(--fake-power-button-observer --power-button-action shutdown) stdout lines = %#v, want %#v", gotLines, wantLines)
	}
	for i, want := range wantLines {
		if gotLines[i] != want {
			t.Fatalf("run(--fake-power-button-observer --power-button-action shutdown) stdout line %d = %q, want %q", i, gotLines[i], want)
		}
	}
	const wantError = `fake power button observer failed: unsupported power_button_action "shutdown" (supported: noop)`
	if !strings.Contains(stderr.String(), wantError) {
		t.Fatalf("run(--fake-power-button-observer --power-button-action shutdown) stderr = %q, want it to contain %q", stderr.String(), wantError)
	}
}

func TestRunFakePowerSignalLowProcessesSwitchOffAndExits(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--fake-power-signal", "low"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--fake-power-signal low) exit = %d, want 0; stderr = %q", got, stderr.String())
	}

	wantLines := []string{
		"fake_power_signal raw=low input=power_switch_line active_signal=low active_switch_state=off interpreted=off processed=true execution_success=true dry_run=true noop_only=true actions_handled=1 real_shutdown=false hardware_action=false",
		`event_breadcrumb index=0 type=daemon.starting message="retroflag-powerd 0.1.0-dev starting dry_run=true"`,
		`event_breadcrumb index=1 type=daemon.ready message="retroflag-powerd ready"`,
		`event_breadcrumb index=2 type=power.intent_received message="power intent received intent=power_button_pressed"`,
		`event_breadcrumb index=3 type=power.dry_run_plan_prepared message="dry-run plan prepared intent=power_button_pressed action=noop"`,
		`event_breadcrumb index=4 type=power.noop_execution_completed message="noop execution completed intent=power_button_pressed actions_handled=1"`,
	}
	assertLines(t, stdout.String(), wantLines)
	assertLifecycleLogs(t, stderr.String(), "run(--fake-power-signal low)")
}

func TestRunFakePowerSignalHighReportsSwitchOnWithoutShutdownRequest(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--fake-power-signal", "high"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--fake-power-signal high) exit = %d, want 0; stderr = %q", got, stderr.String())
	}

	wantLines := []string{
		"fake_power_signal raw=high input=power_switch_line active_signal=low active_switch_state=off interpreted=on processed=false execution_success=false dry_run=false noop_only=false actions_handled=0 real_shutdown=false hardware_action=false",
		`event_breadcrumb index=0 type=daemon.starting message="retroflag-powerd 0.1.0-dev starting dry_run=true"`,
		`event_breadcrumb index=1 type=daemon.ready message="retroflag-powerd ready"`,
	}
	assertLines(t, stdout.String(), wantLines)
	assertLifecycleLogs(t, stderr.String(), "run(--fake-power-signal high)")
	if strings.Contains(stdout.String(), "power intent received") {
		t.Fatalf("run(--fake-power-signal high) stdout = %q, want no power intent breadcrumb", stdout.String())
	}
}

func TestRunFakePowerSignalUnverifiedReportsSwitchUnknownWithoutShutdownRequest(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--fake-power-signal", "unverified"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--fake-power-signal unverified) exit = %d, want 0; stderr = %q", got, stderr.String())
	}

	wantLines := []string{
		"fake_power_signal raw=unverified input=power_switch_line active_signal=low active_switch_state=off interpreted=unknown processed=false execution_success=false dry_run=false noop_only=false actions_handled=0 real_shutdown=false hardware_action=false",
		`event_breadcrumb index=0 type=daemon.starting message="retroflag-powerd 0.1.0-dev starting dry_run=true"`,
		`event_breadcrumb index=1 type=daemon.ready message="retroflag-powerd ready"`,
	}
	assertLines(t, stdout.String(), wantLines)
	assertLifecycleLogs(t, stderr.String(), "run(--fake-power-signal unverified)")
	if strings.Contains(stdout.String(), "power intent received") {
		t.Fatalf("run(--fake-power-signal unverified) stdout = %q, want no power intent breadcrumb", stdout.String())
	}
}

func TestRunFakePowerSignalHonorsConfiguredInterpretation(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{
		"--fake-power-signal", "high",
		"--power-switch-active-signal", "high",
		"--power-switch-active-state", "off",
	}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--fake-power-signal high with active high config) exit = %d, want 0; stderr = %q", got, stderr.String())
	}

	wantLines := []string{
		"fake_power_signal raw=high input=power_switch_line active_signal=high active_switch_state=off interpreted=off processed=true execution_success=true dry_run=true noop_only=true actions_handled=1 real_shutdown=false hardware_action=false",
		`event_breadcrumb index=0 type=daemon.starting message="retroflag-powerd 0.1.0-dev starting dry_run=true"`,
		`event_breadcrumb index=1 type=daemon.ready message="retroflag-powerd ready"`,
		`event_breadcrumb index=2 type=power.intent_received message="power intent received intent=power_button_pressed"`,
		`event_breadcrumb index=3 type=power.dry_run_plan_prepared message="dry-run plan prepared intent=power_button_pressed action=noop"`,
		`event_breadcrumb index=4 type=power.noop_execution_completed message="noop execution completed intent=power_button_pressed actions_handled=1"`,
	}
	assertLines(t, stdout.String(), wantLines)
}

func TestRunFakePowerSignalRejectsInvalidInputClearly(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--fake-power-signal", "floating"}, &stdout, &stderr)

	if got != 1 {
		t.Fatalf("run(--fake-power-signal floating) exit = %d, want 1", got)
	}
	if stdout.String() != "" {
		t.Fatalf("run(--fake-power-signal floating) stdout = %q, want empty", stdout.String())
	}
	const want = `fake power signal failed: unsupported signal state "floating" (supported: low, high, unverified)`
	if !strings.Contains(stderr.String(), want) {
		t.Fatalf("run(--fake-power-signal floating) stderr = %q, want it to contain %q", stderr.String(), want)
	}
}

func TestRunProbeGPIOSignalReportsRawSignalOnly(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	originalProbe := probeGPIOSignal
	probeGPIOSignal = func(context.Context, int) input.SignalState {
		return input.SignalHigh
	}
	defer func() {
		probeGPIOSignal = originalProbe
	}()

	got := run(context.Background(), []string{"--probe-gpio-signal", "4"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--probe-gpio-signal 4) exit = %d, want 0; stderr = %q", got, stderr.String())
	}
	const want = "gpio_signal_probe pin=4 raw=SignalHigh interpreted=false processed=false real_shutdown=false hardware_action=false\n"
	if stdout.String() != want {
		t.Fatalf("run(--probe-gpio-signal 4) stdout = %q, want %q", stdout.String(), want)
	}
	if stderr.String() != "" {
		t.Fatalf("run(--probe-gpio-signal 4) stderr = %q, want empty", stderr.String())
	}
	if strings.Contains(stdout.String(), "Switch") {
		t.Fatalf("run(--probe-gpio-signal 4) stdout = %q, want no interpreted switch state", stdout.String())
	}
}

func TestRunProbeGPIOSignalReportsUnverifiedFallback(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	originalProbe := probeGPIOSignal
	probeGPIOSignal = func(context.Context, int) input.SignalState {
		return input.SignalUnverified
	}
	defer func() {
		probeGPIOSignal = originalProbe
	}()

	got := run(context.Background(), []string{"--probe-gpio-signal", "27"}, &stdout, &stderr)

	if got != 0 {
		t.Fatalf("run(--probe-gpio-signal 27) exit = %d, want 0; stderr = %q", got, stderr.String())
	}
	const want = "gpio_signal_probe pin=27 raw=SignalUnverified interpreted=false processed=false real_shutdown=false hardware_action=false\n"
	if stdout.String() != want {
		t.Fatalf("run(--probe-gpio-signal 27) stdout = %q, want %q", stdout.String(), want)
	}
}

func TestRunProbeGPIOSignalRejectsInvalidPin(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	got := run(context.Background(), []string{"--probe-gpio-signal", "GPIO4"}, &stdout, &stderr)

	if got != 1 {
		t.Fatalf("run(--probe-gpio-signal GPIO4) exit = %d, want 1", got)
	}
	if stdout.String() != "" {
		t.Fatalf("run(--probe-gpio-signal GPIO4) stdout = %q, want empty", stdout.String())
	}
	const want = `gpio signal probe failed: unsupported GPIO pin "GPIO4" (expected non-negative integer)`
	if !strings.Contains(stderr.String(), want) {
		t.Fatalf("run(--probe-gpio-signal GPIO4) stderr = %q, want it to contain %q", stderr.String(), want)
	}
}

func assertLines(t *testing.T, output string, wantLines []string) {
	t.Helper()

	gotLines := strings.Split(strings.TrimSpace(output), "\n")
	if len(gotLines) != len(wantLines) {
		t.Fatalf("stdout lines = %#v, want %#v", gotLines, wantLines)
	}
	for i, want := range wantLines {
		if gotLines[i] != want {
			t.Fatalf("stdout line %d = %q, want %q", i, gotLines[i], want)
		}
	}
}

func assertLifecycleLogs(t *testing.T, logs string, command string) {
	t.Helper()

	for _, wantLog := range []string{
		"retroflag-powerd 0.1.0-dev starting dry_run=true",
		"retroflag-powerd ready",
		"shutdown signal received",
		"retroflag-powerd stopped",
	} {
		if !strings.Contains(logs, wantLog) {
			t.Fatalf("%s stderr = %q, want it to contain %q", command, logs, wantLog)
		}
	}
}
