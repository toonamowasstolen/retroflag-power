package config

import "testing"

func TestDefault(t *testing.T) {
	want := Config{
		AppName:           "retroflag-powerd",
		Version:           "0.1.0-dev",
		DryRun:            true,
		PowerButtonAction: PowerButtonActionNoop,
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
