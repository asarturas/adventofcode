package solution

import "errors"

type Counter struct {}

func (this Counter) Floor(directions string) int {
	floor := 0;
	for _, step := range directions {
		floor += this.oneStep(step)
	}
	return floor;
}

func (this Counter) FirstBasementPosition(directions string) (error, int) {
	floor := 0;
	for position, step := range directions {
		floor += this.oneStep(step)
		if floor == -1 {
			return nil, position + 1
		}
	}
	return errors.New("Haven't been in basement"), 0
}

func (this Counter) oneStep(step rune) int {
	if step == '(' {
		return 1;
	} else if step == ')' {
		return -1;
	}
	return 0;
}
