package gpio

import (
	"context"
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/input"
)

func TestProbeSignalReturnsUnverifiedForNegativePin(t *testing.T) {
	got := ProbeSignal(context.Background(), -1)

	if got != input.SignalUnverified {
		t.Fatalf("ProbeSignal() = %q, want %q", got, input.SignalUnverified)
	}
}

func TestProbeSignalReturnsUnverifiedForCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	got := ProbeSignal(ctx, 4)

	if got != input.SignalUnverified {
		t.Fatalf("ProbeSignal() with canceled context = %q, want %q", got, input.SignalUnverified)
	}
}
