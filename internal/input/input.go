package input

import "context"

type EventType string

const EventTypePowerButtonPressed EventType = "power_button_pressed"

type Event struct {
	Type EventType
}

type Observer interface {
	NextEvent(context.Context) (Event, error)
}
