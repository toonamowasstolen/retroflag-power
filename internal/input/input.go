package input

import "context"

type EventType string

const (
	EventTypeSignal             EventType = "signal"
	EventTypePowerSwitch        EventType = "power_switch"
	EventTypePowerButtonPressed EventType = "power_button_pressed"
)

type SignalState string

const (
	SignalLow        SignalState = "low"
	SignalHigh       SignalState = "high"
	SignalUnverified SignalState = "unverified"
)

func (s SignalState) Valid() bool {
	switch s {
	case SignalLow, SignalHigh, SignalUnverified:
		return true
	default:
		return false
	}
}

func (s SignalState) Label() string {
	switch s {
	case SignalLow:
		return "SignalLow"
	case SignalHigh:
		return "SignalHigh"
	case SignalUnverified:
		return "SignalUnverified"
	default:
		return "SignalUnverified"
	}
}

type SwitchState string

const (
	SwitchOff     SwitchState = "off"
	SwitchOn      SwitchState = "on"
	SwitchUnknown SwitchState = "unknown"
)

type ButtonState string

const (
	ButtonReleased ButtonState = "released"
	ButtonPressed  ButtonState = "pressed"
	ButtonUnknown  ButtonState = "unknown"
)

type Event struct {
	Type        EventType
	Name        string
	SignalState SignalState
	SwitchState SwitchState
}

func SignalEvent(name string, state SignalState) Event {
	return Event{
		Type:        EventTypeSignal,
		Name:        name,
		SignalState: state,
	}
}

func PowerButtonPressedEvent() Event {
	return Event{Type: EventTypePowerButtonPressed}
}

func PowerSwitchEvent(state SwitchState) Event {
	return Event{
		Type:        EventTypePowerSwitch,
		SwitchState: state,
	}
}

type Observer interface {
	NextEvent(context.Context) (Event, error)
}
