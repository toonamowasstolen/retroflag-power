package status

import (
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/config"
)

func TestNew(t *testing.T) {
	want := Status{
		AppName: "retroflag-powerd",
		Version: "0.1.0-dev",
		DryRun:  true,
		State:   StateReady,
	}

	if got := New(config.Default(), StateReady); got != want {
		t.Fatalf("New() = %#v, want %#v", got, want)
	}
}

func TestStates(t *testing.T) {
	states := map[State]string{
		StateStarting: "starting",
		StateReady:    "ready",
		StateStopping: "stopping",
		StateStopped:  "stopped",
	}

	for state, want := range states {
		if got := string(state); got != want {
			t.Errorf("state = %q, want %q", got, want)
		}
	}
}
