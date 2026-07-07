package gpio

import (
	"context"

	"github.com/toonamowasstolen/retroflag-power/internal/input"
)

func ProbeSignal(ctx context.Context, pin int) input.SignalState {
	if err := ctx.Err(); err != nil {
		return input.SignalUnverified
	}

	if pin < 0 {
		return input.SignalUnverified
	}

	state := probeSignal(ctx, pin)
	if !state.Valid() {
		return input.SignalUnverified
	}

	return state
}
