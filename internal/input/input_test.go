package input

import (
	"context"
	"testing"
)

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
