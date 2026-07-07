package power

type Intent string

const IntentPowerButtonPressed Intent = "power_button_pressed"

func (i Intent) String() string {
	return string(i)
}
