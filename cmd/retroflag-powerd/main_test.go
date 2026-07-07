package main

import (
	"bytes"
	"context"
	"strings"
	"testing"
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
