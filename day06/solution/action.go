package solution

type Action interface {
	Do(light Light) Light
}

type TurnOn struct {}

func (this TurnOn) Do(light Light) Light {
	light.Brightness = 1
	return light
}

type TurnOff struct {}

func (this TurnOff) Do(light Light) Light {
	light.Brightness = 0
	return light
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

