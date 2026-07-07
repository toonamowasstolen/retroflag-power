package input

import "fmt"

type ActiveSignal string

const (
	ActiveSignalLow  ActiveSignal = "low"
	ActiveSignalHigh ActiveSignal = "high"
)

type ActiveSwitchState string

const (
	ActiveSwitchStateOff ActiveSwitchState = "off"
	ActiveSwitchStateOn  ActiveSwitchState = "on"
)

type LatchingPowerSwitchOptions struct {
	ActiveSignal      ActiveSignal
	ActiveSwitchState ActiveSwitchState
}

type UnsupportedLatchingPowerSwitchConfigError struct {
	Field string
	Value string
}

func (e UnsupportedLatchingPowerSwitchConfigError) Error() string {
	switch e.Field {
	case "active_signal":
		return fmt.Sprintf("unsupported active_signal %q (supported: low, high)", e.Value)
	case "active_switch_state":
		return fmt.Sprintf("unsupported active_switch_state %q (supported: off, on)", e.Value)
	default:
		return fmt.Sprintf("unsupported latching power switch config %s=%q", e.Field, e.Value)
	}
}

type UnsupportedSignalEventError struct {
	EventType EventType
}

func (e UnsupportedSignalEventError) Error() string {
	return fmt.Sprintf("unsupported latching power switch input event %q (supported: signal)", e.EventType)
}

type UnsupportedSignalStateError struct {
	State SignalState
}

func (e UnsupportedSignalStateError) Error() string {
	return fmt.Sprintf("unsupported signal state %q (supported: low, high, unverified)", e.State)
}

func InterpretLatchingPowerSwitchSignal(state SignalState, options LatchingPowerSwitchOptions) (SwitchState, error) {
	if err := options.Validate(); err != nil {
		return SwitchUnknown, err
	}

	switch state {
	case SignalUnverified:
		return SwitchUnknown, nil
	case SignalLow, SignalHigh:
	default:
		return SwitchUnknown, UnsupportedSignalStateError{State: state}
	}

	if signalIsActive(state, options.ActiveSignal) {
		return SwitchState(options.ActiveSwitchState), nil
	}

	return inactiveSwitchState(options.ActiveSwitchState), nil
}

func InterpretLatchingPowerSwitchEvent(event Event, options LatchingPowerSwitchOptions) (Event, error) {
	if event.Type != EventTypeSignal {
		return Event{}, UnsupportedSignalEventError{EventType: event.Type}
	}

	state, err := InterpretLatchingPowerSwitchSignal(event.SignalState, options)
	if err != nil {
		return Event{}, err
	}

	return PowerSwitchEvent(state), nil
}

func (o LatchingPowerSwitchOptions) Validate() error {
	switch o.ActiveSignal {
	case ActiveSignalLow, ActiveSignalHigh:
	default:
		return UnsupportedLatchingPowerSwitchConfigError{
			Field: "active_signal",
			Value: string(o.ActiveSignal),
		}
	}

	switch o.ActiveSwitchState {
	case ActiveSwitchStateOff, ActiveSwitchStateOn:
	default:
		return UnsupportedLatchingPowerSwitchConfigError{
			Field: "active_switch_state",
			Value: string(o.ActiveSwitchState),
		}
	}

	return nil
}

func signalIsActive(state SignalState, active ActiveSignal) bool {
	return (state == SignalLow && active == ActiveSignalLow) ||
		(state == SignalHigh && active == ActiveSignalHigh)
}

func inactiveSwitchState(active ActiveSwitchState) SwitchState {
	if active == ActiveSwitchStateOff {
		return SwitchOn
	}

	return SwitchOff
}
