//go:build linux

package gpio

import (
	"testing"

	"github.com/toonamowasstolen/retroflag-power/internal/input"
)

func TestSignalFromGPIOValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  input.SignalState
		ok    bool
	}{
		{name: "low", value: "0\n", want: input.SignalLow, ok: true},
		{name: "high", value: "1\n", want: input.SignalHigh, ok: true},
		{name: "unknown", value: "active\n", want: input.SignalUnverified, ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := signalFromGPIOValue(tt.value)
			if got != tt.want || ok != tt.ok {
				t.Fatalf("signalFromGPIOValue(%q) = %q, %t; want %q, %t", tt.value, got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestSignalFromDebugGPIO(t *testing.T) {
	const contents = `
 gpio-4   (                    |sysfs               ) in  lo
 gpio-17  (                    |sysfs               ) in  hi
`

	got, ok := signalFromDebugGPIO(contents, 17)

	if got != input.SignalHigh || !ok {
		t.Fatalf("signalFromDebugGPIO() = %q, %t; want %q, true", got, ok, input.SignalHigh)
	}
}
