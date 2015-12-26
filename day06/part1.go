package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day06/solution"
	"fmt"
)

func main() {
	parser := solution.InstructionParser{};
	instructions := make([]solution.Instruction, 0);
	for instruction := io.ReadLine(); instruction != ""; instruction = io.ReadLine() {
		fmt.Println(instruction);
		instructions = append(instructions, parser.Parse(instruction));
	}
	grid := solution.NewGrid(1000);
	grid.Apply(instructions);
	fmt.Println(grid.OnLights());
}
