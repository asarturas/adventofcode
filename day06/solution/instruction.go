package solution
import (
	"strings"
	"strconv"
)

type Coord struct {
	Row, Col int
}

type Action interface {
	Do(current bool) bool
}

type TurnOn struct {}

func (this TurnOn) Do(current bool) bool {
	return true;
}

type TurnOff struct {}

func (this TurnOff) Do(current bool) bool {
	return false;
}

type Toggle struct {}

func (this Toggle) Do(current bool) bool {
	return !current;
}

type Instruction struct {
	Act Action
	Start, End Coord
}

type InstructionParser struct {}

func (this InstructionParser) Parse(writing string) Instruction {
	ins := Instruction{}
	ins.Act, writing = this.parseAction(writing)
	ins.Start, ins.End = this.parseCoords(writing)
	return ins;
}

func (this InstructionParser) parseAction(writing string) (Action, string) {
	if strings.HasPrefix(writing, "turn on") {
		return TurnOn{}, strings.Replace(writing, "turn on ", "", 1);
	} else if strings.HasPrefix(writing, "turn off") {
		return TurnOff{}, strings.Replace(writing, "turn off ", "", 1);
	} else {
		return Toggle{}, strings.Replace(writing, "toggle ", "", 1);
	}
}

func (this InstructionParser) parseCoords(writing string) (Coord, Coord) {
	parts := strings.Split(writing, " ");
	start := strings.Split(parts[0], ",");
	startRow, _ := strconv.Atoi(start[0]);
	startCol, _ := strconv.Atoi(start[1]);
	end := strings.Split(parts[2], ",");
	endRow, _ := strconv.Atoi(end[0]);
	endCol, _ := strconv.Atoi(end[1]);
	return Coord{startRow, startCol}, Coord{endRow, endCol}
}

type Grid struct {
	Size int
	lights []bool
}

func NewGrid(size int) *Grid {
	grid := Grid{};
	grid.Size = size;
	grid.lights = make([]bool, size * size);
	return &grid;
}

func (this *Grid) HasLightOn(coord Coord) bool {
	return this.lights[this.getLocation(coord)]
}

func (this *Grid) getLocation(coord Coord) int {
	return coord.Row * this.Size + coord.Col;
}

func (this *Grid) OnLights() int {
	number := 0;
	for i := range this.lights {
		if this.lights[i] {
			number++;
		}
	}
	return number;
}

func (this *Grid) Apply(instructions []Instruction) {
	for _, instruction := range instructions {
		this.applyOne(instruction);
	}
}

func (this *Grid) applyOne(instruction Instruction) {
	for row := instruction.Start.Row; row <= instruction.End.Row; row++ {
		for col := instruction.Start.Col; col <= instruction.End.Col; col++ {
			this.applyTo(Coord{row, col}, instruction.Act);
		}
	}
}

func (this *Grid) applyTo(coord Coord, act Action) {
	location := this.getLocation(coord);
	current := this.lights[location];
	this.lights[location] = act.Do(current);
}
