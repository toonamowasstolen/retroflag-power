package input

import (
	"context"
	"testing"
)

func TestSignalEventWithSignalLowIsDeterministic(t *testing.T) {
	got := SignalEvent("power_switch_line", SignalLow)
	want := Event{
		Type:        EventTypeSignal,
		Name:        "power_switch_line",
		SignalState: SignalLow,
	}

	if got != want {
		t.Fatalf("SignalEvent() with SignalLow = %#v, want %#v", got, want)
	}
	if !got.SignalState.Valid() {
		t.Fatalf("SignalEvent() with SignalLow has invalid state %q", got.SignalState)
	}
}

func TestSignalEventWithSignalHighIsDeterministic(t *testing.T) {
	got := SignalEvent("power_switch_line", SignalHigh)
	want := Event{
		Type:        EventTypeSignal,
		Name:        "power_switch_line",
		SignalState: SignalHigh,
	}

	if got != want {
		t.Fatalf("SignalEvent() with SignalHigh = %#v, want %#v", got, want)
	}
	if !got.SignalState.Valid() {
		t.Fatalf("SignalEvent() with SignalHigh has invalid state %q", got.SignalState)
	}
}

func TestSignalEventWithSignalUnverifiedIsDeterministic(t *testing.T) {
	got := SignalEvent("power_switch_line", SignalUnverified)
	want := Event{
		Type:        EventTypeSignal,
		Name:        "power_switch_line",
		SignalState: SignalUnverified,
	}

	if got != want {
		t.Fatalf("SignalEvent() with SignalUnverified = %#v, want %#v", got, want)
	}
	if !got.SignalState.Valid() {
		t.Fatalf("SignalEvent() with SignalUnverified has invalid state %q", got.SignalState)
	}
}

func TestSignalStateValidRejectsUnknownStates(t *testing.T) {
	if SignalState("floating").Valid() {
		t.Fatal(`SignalState("floating").Valid() = true, want false`)
	}
}

func TestFakePowerButtonObserverEmitsPowerButtonPressed(t *testing.T) {
	observer := NewFakePowerButtonObserver()

	got, err := observer.NextEvent(context.Background())
	if err != nil {
		t.Fatalf("NextEvent() error = %v, want nil", err)
	}

	want := Event{Type: EventTypePowerButtonPressed}
	if got != want {
		t.Fatalf("NextEvent() = %#v, want %#v", got, want)
	}
}

func TestPowerButtonPressedEventNamesObserverEvent(t *testing.T) {
	got := PowerButtonPressedEvent()
	want := Event{Type: EventTypePowerButtonPressed}
	if got != want {
		t.Fatalf("PowerButtonPressedEvent() = %#v, want %#v", got, want)
	}
}

func TestFakeObserverCopiesEvents(t *testing.T) {
	events := []Event{{Type: EventTypePowerButtonPressed}}
	observer := NewFakeObserver(events...)
	events[0].Type = "changed"

	got, err := observer.NextEvent(context.Background())
	if err != nil {
		t.Fatalf("NextEvent() error = %v, want nil", err)
	}

	want := Event{Type: EventTypePowerButtonPressed}
	if got != want {
		t.Fatalf("NextEvent() = %#v after source mutation, want %#v", got, want)
	}
}

func TestFakeObserverHonorsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	observer := NewFakePowerButtonObserver()

	if _, err := observer.NextEvent(ctx); err == nil {
		t.Fatal("NextEvent() error = nil, want canceled context error")
	}
}
