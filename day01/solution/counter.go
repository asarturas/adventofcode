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

func (this Counter) FirstBasementPosition(directions string) (int, error) {
	floor := 0;
	for position, step := range directions {
		floor += this.oneStep(step)
		if floor == -1 {
			return position + 1, nil
		}
	}
	return 0, errors.New("Haven't been in basement")
}

func (this Counter) oneStep(step rune) int {
	if step == '(' {
		return 1;
	} else if step == ')' {
		return -1;
	}
	return 0;
}
