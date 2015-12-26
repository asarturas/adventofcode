package solution

import (
	"testing"
)

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

func Test_it_returns_a_sum_of_brightnesses_from_a_grid(t *testing.T) {
	grid := NewGrid(2)
	instructions := []Instruction{
		Instruction{TurnOnBrightness{}, Coord{0,0}, Coord{1,1}},
		Instruction{ToggleBrightness{}, Coord{1,0}, Coord{1,1}},
		Instruction{TurnOffBrightness{}, Coord{1,1}, Coord{1,1}},
	}
	grid.Apply(instructions)
	if grid.TotalBrightness() != 7 {
		t.Error("Expected a total brightness of 7, but instead got", grid.TotalBrightness())
	}
}
