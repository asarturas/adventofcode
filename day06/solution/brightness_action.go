package solution

type TurnOnBrightness struct {}

func (this TurnOnBrightness) Do(light Light) Light {
	light.Brightness++
	return light
}

type TurnOffBrightness struct {}

func (this TurnOffBrightness) Do(light Light) Light {
	if (light.IsOn()) {
		light.Brightness--
	}
	return light
}

type ToggleBrightness struct {}

func (this ToggleBrightness) Do(light Light) Light {
	light.Brightness += 2
	return light
}
