package config

import (
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/input"
)

func TestDefault(t *testing.T) {
	want := Config{
		AppName:           "retroflag-powerd",
		Version:           "0.1.0-dev",
		DryRun:            true,
		PowerButtonAction: PowerButtonActionNoop,
		PowerInputName:    "power_switch_line",
		LatchingPowerSwitch: input.LatchingPowerSwitchOptions{
			ActiveSignal:      input.ActiveSignalLow,
			ActiveSwitchState: input.ActiveSwitchStateOff,
		},
	}

	if got := Default(); got != want {
		t.Fatalf("Default() = %#v, want %#v", got, want)
	}
}

func TestValidatePowerButtonActionAcceptsNoop(t *testing.T) {
	cfg := Default()
	cfg.PowerButtonAction = PowerButtonActionNoop

	if err := cfg.ValidatePowerButtonAction(); err != nil {
		t.Fatalf("ValidatePowerButtonAction() error = %v, want nil", err)
	}
}

func TestValidatePowerButtonActionRejectsUnsupportedValueClearly(t *testing.T) {
	cfg := Default()
	cfg.PowerButtonAction = "shutdown"

	err := cfg.ValidatePowerButtonAction()
	if err == nil {
		t.Fatal("ValidatePowerButtonAction() error = nil, want unsupported policy error")
	}

	const want = `unsupported power_button_action "shutdown" (supported: noop)`
	if err.Error() != want {
		t.Fatalf("ValidatePowerButtonAction() error = %q, want %q", err.Error(), want)
	}
}
