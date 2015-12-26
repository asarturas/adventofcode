package solution

type Action interface {
	Do(light Light) Light
	String() string
}

type TurnOn struct {}

func (this TurnOn) Do(light Light) Light {
	light.Brightness = 1
	return light
}

func (this TurnOn) String() string {
	return "turn on "
}

type TurnOff struct {}

func (this TurnOff) Do(light Light) Light {
	light.Brightness = 0
	return light
}

func (this TurnOff) String() string {
	return "turn off "
}

type Toggle struct {}

func (this Toggle) Do(light Light) Light {
	if light.IsOn() {
		light.Brightness = 0
	} else {
		light.Brightness = 1
	}
	return light
}

func (this Toggle) String() string {
	return "toggle "
}
