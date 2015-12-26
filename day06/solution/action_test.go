package solution

import (
	"testing"
)

func Test_turn_on_action_turns_on_or_keeps_on(t *testing.T) {
	action := TurnOn{};
	if action.Do(false) != true {
		t.Error("Expected 'turn on' to turn on the light.")
	} else if action.Do(true) != true {
		t.Error("Expected 'turn on' to keep the light turned on.")
	}
}

func Test_turn_off_action_turns_off_or_keeps_off(t *testing.T) {
	action := TurnOff{};
	if action.Do(true) != false {
		t.Error("Expected 'turn off' to turn off the light.")
	} else if action.Do(false) != false {
		t.Error("Expected 'turn off' to keep the light off.")
	}
}

func Test_toggle_action_turns_on_off_light_and_turns_off_on_light(t *testing.T) {
	action := Toggle{};
	if action.Do(true) != false {
		t.Error("Expected 'toggle' to turn off the light, which is on.")
	} else if action.Do(false) != true {
		t.Error("Expected 'toggle' to turn on the light, which is off.")
	}
}

