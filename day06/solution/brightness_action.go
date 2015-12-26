package solution

type TurnOnBrightness struct {}

func (this TurnOnBrightness) Do(light Light) Light {
	light.Brightness++
	return light
}

func (this TurnOnBrightness) String() string {
	return "turn on "
}

type TurnOffBrightness struct {}

func (this TurnOffBrightness) Do(light Light) Light {
	if (light.IsOn()) {
		light.Brightness--
	}
	return light
}

func (this TurnOffBrightness) String() string {
	return "turn off "
}

type ToggleBrightness struct {}

func (this ToggleBrightness) Do(light Light) Light {
	light.Brightness += 2
	return light
}

func (this ToggleBrightness) String() string {
	return "toggle "
}
