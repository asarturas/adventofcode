package solution

import (
	"testing"
)

func Test_turn_on_action_turns_on_or_keeps_on(t *testing.T) {
	action := TurnOn{}
	expected := Light{1}
	if action.Do(Light{0}) != expected {
		t.Error("Expected 'turn on' to turn on the light.")
	} else if action.Do(Light{1}) != expected {
		t.Error("Expected 'turn on' to keep the light turned on.")
	}
}

func Test_turn_off_action_turns_off_or_keeps_off(t *testing.T) {
	action := TurnOff{}
	expected := Light{0}
	if action.Do(Light{1}) != expected {
		t.Error("Expected 'turn off' to turn off the light.")
	} else if action.Do(Light{1}) != expected {
		t.Error("Expected 'turn off' to keep the light off.")
	}
}

func Test_toggle_action_turns_on_off_light_and_turns_off_on_light(t *testing.T) {
	action := Toggle{}
	expectedOff := Light{0}
	expectedOn := Light{1}
	if action.Do(Light{1}) != expectedOff {
		t.Error("Expected 'toggle' to turn off the light, which is on.")
	} else if action.Do(Light{0}) != expectedOn {
		t.Error("Expected 'toggle' to turn on the light, which is off.")
	}
}

