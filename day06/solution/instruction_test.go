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

func Test_it_parses_on_instruction(t *testing.T) {
	parser := InstructionParser{};
	instruction := parser.Parse("turn on 489,959 through 759,964")
	expectedStart := Coord{489, 959};
	expectedEnd := Coord{759, 964};
	expectedAction := TurnOn{};
	if instruction.Act != expectedAction {
		t.Error("Unexpected action", instruction.Act)
	} else if instruction.Start != expectedStart {
		t.Error("Unexpected start coordinate", instruction.Start)
	} else if instruction.End != expectedEnd {
		t.Error("Unexpected end coordinate", instruction.End)
	}
}

func Test_it_parses_off_instruction(t *testing.T) {
	parser := InstructionParser{};
	instruction := parser.Parse("turn off 488,956 through 795,960")
	expectedStart := Coord{488, 956};
	expectedEnd := Coord{795, 960};
	expectedAction := TurnOff{};
	if instruction.Act != expectedAction {
		t.Error("Unexpected action", instruction.Act)
	} else if instruction.Start != expectedStart {
		t.Error("Unexpected start coordinate", instruction.Start)
	} else if instruction.End != expectedEnd {
		t.Error("Unexpected end coordinate", instruction.End)
	}
}

func Test_it_parses_toggle_instruction(t *testing.T) {
	parser := InstructionParser{};
	instruction := parser.Parse("toggle 487,955 through 794,959")
	expectedStart := Coord{487, 955};
	expectedEnd := Coord{794, 959};
	expectedAction := Toggle{};
	if instruction.Act != expectedAction {
		t.Error("Unexpected action", instruction.Act)
	} else if instruction.Start != expectedStart {
		t.Error("Unexpected start coordinate", instruction.Start)
	} else if instruction.End != expectedEnd {
		t.Error("Unexpected end coordinate", instruction.End)
	}
}

func Test_it_flattens_coordinates_internally(t *testing.T) {
	grid := NewGrid(7);
	if grid.getLocation(Coord{3, 5}) != 26 {
		t.Error("Coordinate (3, 5) should have been at location 26, but got", grid.getLocation(Coord{3, 5}));
	}
}

func Test_the_grid_initially_all_off(t *testing.T) {
	grid := NewGrid(3);
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid.HasLightOn(Coord{i, j}) {
				t.Error("Expected light at (%d, %d) to be off, but it's ON", i, j)
			}
		}
	}
	if grid.OnLights() != 0 {
		t.Error("Expected not to have any lights on")
	}
}

func Test_grid_modified_by_set_of_instructions(t *testing.T) {
	size := 5;
	grid := NewGrid(size);
	instructions := []Instruction{
		Instruction{TurnOn{}, Coord{1, 1}, Coord{3, 3}},
		Instruction{Toggle{}, Coord{2, 2}, Coord{2, 2}},
	}

	grid.Apply(instructions);

	for i := 0; i < size; i++ {
		if grid.HasLightOn(Coord{0, i}) {
			t.Error("Expected ligths on first row to be off, but (%d, %d) is on.", 0, i);
		}
	}
	if grid.HasLightOn(Coord{2, 0}) {
		t.Error("Light at (2, 0) should have been off");
	}
	if !grid.HasLightOn(Coord{2, 1}) {
		t.Error("Light at (2, 1) should have been on");
	}
	if grid.HasLightOn(Coord{2, 2}) {
		t.Error("Centre light at (2, 2) should have been off")
	}
	if !grid.HasLightOn(Coord{2, 3}) {
		t.Error("Light at (2, 3) should have been on");
	}
	if grid.HasLightOn(Coord{2, 4}) {
		t.Error("Light at (2, 4) should have been off");
	}
	for i := 0; i < size; i++ {
		if grid.HasLightOn(Coord{4, i}) {
			t.Error("Expected lights on last row to be off, but (%d, %d) is on.", 4, i);
		}
	}
	if grid.OnLights() != 8 {
		t.Error("It was supposed to have 8 lights on", grid.OnLights());
	}
}
