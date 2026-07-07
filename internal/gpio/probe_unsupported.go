//go:build !linux

package gpio

import (
	"context"

	"github.com/toonamowasstolen/retroflag-power/internal/input"
)

func probeSignal(context.Context, int) input.SignalState {
	return input.SignalUnverified
}
