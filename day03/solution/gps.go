package solution

type location struct {
	x, y int
}

type Gps struct {
	current location
	visited map[location]int
}

func NewGps() *Gps {
	gps := &Gps{
		location{0, 0},
		make(map[location]int),
	}
	gps.visited[gps.current]++;
	return gps;
}

func (this *Gps) Move(move rune) {
	this.current = this.locate(move, this.current);
	this.visited[this.current]++;
}

func (this *Gps) Route(route string) {
	for _, move := range route {
		this.Move(move);
	}
}

func (this *Gps) VisitedPlaces() int {
	return len(this.visited);
}

func (this *Gps) locate(move rune, current location) location {
	if move == '^' {
		current.y++;
	} else if move == 'v' {
		current.y--;
	} else if move == '<' {
		current.x--;
	} else if move == '>' {
		current.x++;
	}
	return current;
}
