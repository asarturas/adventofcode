package solution

import (
	"testing"
)

func Test_turn_on_action_increases_brightness_by_one(t *testing.T) {
	action := TurnOnBrightness{}
	light := Light{1}
	expected := Light{2}
	if action.Do(light) != expected {
		t.Error("Expected 'turn on' to increase brightness by one, got", action.Do(light).Brightness)
	}
}

func Test_turn_off_action_decreases_brightness_unless_it_is_already_at_zero(t *testing.T) {
	action := TurnOffBrightness{}
	light := Light{2}
	expected1 := Light{1}
	expected2 := Light{0}
	if action.Do(light) != expected1 {
		t.Error("Expected 'turn off' to lower brightness, got",  action.Do(light).Brightness)
	}
	if action.Do(expected2) != expected2 {
		t.Error("Expected 'turn off' to keep brightness at zero, got", action.Do(expected2).Brightness)
	}
}

func Test_toggle_action_increases_brightness_by_two(t *testing.T) {
	action := ToggleBrightness{}
	expected := Light{3}
	if action.Do(Light{1}) != expected {
		t.Error("Expected 'toggle' to increase brightness by two", action.Do(Light{1}))
	}
}

