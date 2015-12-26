package solution

import (
	"testing"
)

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
