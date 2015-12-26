package solution

type Grid struct {
	Size int
	lights []Light
}

func NewGrid(size int) *Grid {
	grid := Grid{};
	grid.Size = size;
	grid.lights = make([]Light, size * size);
	return &grid;
}

func (this *Grid) HasLightOn(coord Coord) bool {
	return this.lights[this.getLocation(coord)].IsOn()
}

func (this *Grid) getLocation(coord Coord) int {
	return coord.Row * this.Size + coord.Col
}

func (this *Grid) OnLights() int {
	number := 0;
	for i := range this.lights {
		if this.lights[i].IsOn() {
			number++
		}
	}
	return number
}

func (this *Grid) Apply(instructions []Instruction) {
	for _, instruction := range instructions {
		this.applyOne(instruction)
	}
}

func (this *Grid) applyOne(instruction Instruction) {
	for row := instruction.Start.Row; row <= instruction.End.Row; row++ {
		for col := instruction.Start.Col; col <= instruction.End.Col; col++ {
			this.applyTo(Coord{row, col}, instruction.Act)
		}
	}
}

func (this *Grid) applyTo(coord Coord, act Action) {
	location := this.getLocation(coord)
	current := this.lights[location]
	this.lights[location] = act.Do(current)
}

func (this *Grid) TotalBrightness() int {
	total := 0
	for _, light := range this.lights {
		total += light.Brightness
	}
	return total
}
