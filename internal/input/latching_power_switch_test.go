package input

import "testing"

func TestInterpretLatchingPowerSwitchSignalMapsLowToOffWhenConfigured(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalLow,
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	got, err := InterpretLatchingPowerSwitchSignal(SignalLow, options)
	if err != nil {
		t.Fatalf("InterpretLatchingPowerSwitchSignal() error = %v, want nil", err)
	}

	if got != SwitchOff {
		t.Fatalf("InterpretLatchingPowerSwitchSignal(SignalLow) = %q, want %q", got, SwitchOff)
	}
}

func TestInterpretLatchingPowerSwitchSignalMapsHighToOffWhenConfigured(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalHigh,
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	got, err := InterpretLatchingPowerSwitchSignal(SignalHigh, options)
	if err != nil {
		t.Fatalf("InterpretLatchingPowerSwitchSignal() error = %v, want nil", err)
	}

	if got != SwitchOff {
		t.Fatalf("InterpretLatchingPowerSwitchSignal(SignalHigh) = %q, want %q", got, SwitchOff)
	}
}

func TestInterpretLatchingPowerSwitchSignalMapsInactiveSignalToOppositeState(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalLow,
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	got, err := InterpretLatchingPowerSwitchSignal(SignalHigh, options)
	if err != nil {
		t.Fatalf("InterpretLatchingPowerSwitchSignal() error = %v, want nil", err)
	}

	if got != SwitchOn {
		t.Fatalf("InterpretLatchingPowerSwitchSignal(SignalHigh) = %q, want %q", got, SwitchOn)
	}
}

func TestInterpretLatchingPowerSwitchSignalMapsUnverifiedToUnknown(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalLow,
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	got, err := InterpretLatchingPowerSwitchSignal(SignalUnverified, options)
	if err != nil {
		t.Fatalf("InterpretLatchingPowerSwitchSignal() error = %v, want nil", err)
	}

	if got != SwitchUnknown {
		t.Fatalf("InterpretLatchingPowerSwitchSignal(SignalUnverified) = %q, want %q", got, SwitchUnknown)
	}
}

func TestInterpretLatchingPowerSwitchSignalRejectsUnsupportedActiveSignalClearly(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      "floating",
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	_, err := InterpretLatchingPowerSwitchSignal(SignalLow, options)
	if err == nil {
		t.Fatal("InterpretLatchingPowerSwitchSignal() error = nil, want unsupported config error")
	}

	const want = `unsupported active_signal "floating" (supported: low, high)`
	if err.Error() != want {
		t.Fatalf("InterpretLatchingPowerSwitchSignal() error = %q, want %q", err.Error(), want)
	}
}

func TestInterpretLatchingPowerSwitchSignalRejectsUnsupportedActiveSwitchStateClearly(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalLow,
		ActiveSwitchState: "standby",
	}

	_, err := InterpretLatchingPowerSwitchSignal(SignalLow, options)
	if err == nil {
		t.Fatal("InterpretLatchingPowerSwitchSignal() error = nil, want unsupported config error")
	}

	const want = `unsupported active_switch_state "standby" (supported: off, on)`
	if err.Error() != want {
		t.Fatalf("InterpretLatchingPowerSwitchSignal() error = %q, want %q", err.Error(), want)
	}
}

func TestInterpretLatchingPowerSwitchSignalRejectsUnsupportedSignalStateClearly(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalLow,
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	_, err := InterpretLatchingPowerSwitchSignal("floating", options)
	if err == nil {
		t.Fatal("InterpretLatchingPowerSwitchSignal() error = nil, want unsupported signal state error")
	}

	const want = `unsupported signal state "floating" (supported: low, high, unverified)`
	if err.Error() != want {
		t.Fatalf("InterpretLatchingPowerSwitchSignal() error = %q, want %q", err.Error(), want)
	}
}

func TestInterpretLatchingPowerSwitchEventConsumesSignalEvent(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalLow,
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	got, err := InterpretLatchingPowerSwitchEvent(SignalEvent("power_switch_line", SignalLow), options)
	if err != nil {
		t.Fatalf("InterpretLatchingPowerSwitchEvent() error = %v, want nil", err)
	}

	want := PowerSwitchEvent(SwitchOff)
	if got != want {
		t.Fatalf("InterpretLatchingPowerSwitchEvent() = %#v, want %#v", got, want)
	}
}

func TestInterpretLatchingPowerSwitchEventRejectsUnsupportedEventTypeClearly(t *testing.T) {
	options := LatchingPowerSwitchOptions{
		ActiveSignal:      ActiveSignalLow,
		ActiveSwitchState: ActiveSwitchStateOff,
	}

	_, err := InterpretLatchingPowerSwitchEvent(PowerButtonPressedEvent(), options)
	if err == nil {
		t.Fatal("InterpretLatchingPowerSwitchEvent() error = nil, want unsupported event error")
	}

	const want = `unsupported latching power switch input event "power_button_pressed" (supported: signal)`
	if err.Error() != want {
		t.Fatalf("InterpretLatchingPowerSwitchEvent() error = %q, want %q", err.Error(), want)
	}
}
