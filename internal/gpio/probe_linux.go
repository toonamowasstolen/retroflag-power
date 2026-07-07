//go:build linux

package gpio

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/toonamowasstolen/retroflag-power/internal/input"
)

const debugGPIOPath = "/sys/kernel/debug/gpio"

func probeSignal(ctx context.Context, pin int) input.SignalState {
	if state, ok := readSysfsGPIOValue(pin); ok {
		return state
	}
	if state, ok := readDebugGPIOValue(pin); ok {
		return state
	}
	if state, ok := readGPIODValue(ctx, pin); ok {
		return state
	}

	return input.SignalUnverified
}

func readSysfsGPIOValue(pin int) (input.SignalState, bool) {
	valuePath := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
	contents, err := os.ReadFile(valuePath)
	if err != nil {
		return input.SignalUnverified, false
	}

	return signalFromGPIOValue(string(contents))
}

func readDebugGPIOValue(pin int) (input.SignalState, bool) {
	contents, err := os.ReadFile(debugGPIOPath)
	if err != nil {
		return input.SignalUnverified, false
	}

	return signalFromDebugGPIO(string(contents), pin)
}

func readGPIODValue(ctx context.Context, pin int) (input.SignalState, bool) {
	path, err := exec.LookPath("gpioget")
	if err != nil {
		return input.SignalUnverified, false
	}

	probeCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	output, err := exec.CommandContext(probeCtx, path, "gpiochip0", strconv.Itoa(pin)).Output()
	if err != nil {
		return input.SignalUnverified, false
	}

	return signalFromGPIOValue(string(output))
}

func signalFromGPIOValue(value string) (input.SignalState, bool) {
	switch strings.TrimSpace(value) {
	case "0":
		return input.SignalLow, true
	case "1":
		return input.SignalHigh, true
	default:
		return input.SignalUnverified, false
	}
}

func signalFromDebugGPIO(contents string, pin int) (input.SignalState, bool) {
	want := fmt.Sprintf("gpio-%d", pin)
	for _, line := range strings.Split(contents, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 || fields[0] != want {
			continue
		}

		switch fields[len(fields)-1] {
		case "lo":
			return input.SignalLow, true
		case "hi":
			return input.SignalHigh, true
		default:
			return input.SignalUnverified, false
		}
	}

	return input.SignalUnverified, false
}
