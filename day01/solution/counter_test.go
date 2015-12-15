package solution

import (
	"testing"
)

var counter = Counter{}

func Test_going_up_one_floor(t *testing.T) {
	if counter.Floor("(") != 1 {
		t.Errorf("Expected to go up one floor, but got to %d instead.", counter.Floor("("));
	}
}

func Test_going_up_five_floors(t *testing.T) {
	if counter.Floor("(((((") != 5 {
		t.Errorf("Expected to go up five floors, but got to %d instead.", counter.Floor("((((("));
	}
}

func Test_going_down_three_floors(t *testing.T) {
	if counter.Floor(")))") != -3 {
		t.Errorf("Expected to go down three floors, but got to %d instead.", counter.Floor(")))"));
	}
}

func Test_samples_from_exercise(t *testing.T) {
	if counter.Floor("(())") != 0 {
		t.Errorf("Expected to get 0, but got %d instead.", counter.Floor("(())"));
	} else if counter.Floor("()()") != 0 {
		t.Errorf("Expected to get 0, but got %d instead.", counter.Floor("()()"));
	} else if counter.Floor("(()(()(") != 3 {
		t.Errorf("Expected to get 3, but got %d instead.", counter.Floor("(()(()("));
	} else if counter.Floor("))(((((") != 3 {
		t.Errorf("Expected to get 3, but got %d instead.", counter.Floor("))((((("));
	} else if counter.Floor("())") != -1 {
		t.Errorf("Expected to get -1, but got %d instead.", counter.Floor("())"));
	} else if counter.Floor("))(") != -1 {
		t.Errorf("Expected to get -1, but got %d instead.", counter.Floor("))("));
	} else if counter.Floor(")())())") != -3 {
		t.Errorf("Expected to get -1, but got %d instead.", counter.Floor(")())())"));
	}
}

func Test_it_enters_basement_on_first_step(t *testing.T) {
	basementPosition, err := counter.FirstBasementPosition(")")
	if err != nil {
		t.Error("Unexpected error %s", err)
	} else if basementPosition != 1 {
		t.Error("Expected to enter basement in first step, but got there at %d", basementPosition);
	}
}

func Test_it_enters_basement_on_fifth_step(t *testing.T) {
	basementPosition, err := counter.FirstBasementPosition("()())")
	if err != nil {
		t.Error("Unexpected error %s", err)
	} else if basementPosition != 5 {
		t.Error("Expected to enter basement in fifth step, but got there at %d", basementPosition);
	}
}
