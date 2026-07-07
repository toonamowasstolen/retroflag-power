package input

import "context"

type FakeObserver struct {
	events []Event
}

func NewFakeObserver(events ...Event) *FakeObserver {
	return &FakeObserver{
		events: append([]Event(nil), events...),
	}
}

func NewFakePowerButtonObserver() *FakeObserver {
	return NewFakeObserver(PowerButtonPressedEvent())
}

func (o *FakeObserver) NextEvent(ctx context.Context) (Event, error) {
	if err := ctx.Err(); err != nil {
		return Event{}, err
	}

	if len(o.events) == 0 {
		return Event{}, nil
	}

	event := o.events[0]
	o.events = o.events[1:]
	return event, nil
}
