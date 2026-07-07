package config

import (
	"fmt"

	"github.com/toonamowasstolen/retroflag-power/internal/version"
)

const PowerButtonActionNoop = "noop"

type Config struct {
	AppName           string
	Version           string
	DryRun            bool
	PowerButtonAction string
}

type UnsupportedPowerButtonActionError struct {
	Action string
}

func (e UnsupportedPowerButtonActionError) Error() string {
	return fmt.Sprintf("unsupported power_button_action %q (supported: %s)", e.Action, PowerButtonActionNoop)
}

func Default() Config {
	return Config{
		AppName:           version.Name,
		Version:           version.Version,
		DryRun:            true,
		PowerButtonAction: PowerButtonActionNoop,
	}
}

func (c Config) ValidatePowerButtonAction() error {
	if c.PowerButtonAction == PowerButtonActionNoop {
		return nil
	}

	return UnsupportedPowerButtonActionError{Action: c.PowerButtonAction}
}
