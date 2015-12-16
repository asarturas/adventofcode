package solution

import(
	"testing"
)

func Test_it_returns_updated_position_after_a_move(t *testing.T) {
	navigator := NewGps();
	var newLocation location;
	newLocation = navigator.locate('^', location{0, 0})
	if newLocation.x != 0 || newLocation.y != 1 {
		t.Errorf("Expected to be at (0,1), but are at (%d,%d) instead.", newLocation.x, newLocation.y);
	}
	newLocation = navigator.locate('v', location{0, 0})
	if newLocation.x != 0 || newLocation.y != -1 {
		t.Error("Expected to be at (0,-1), but are at (%d,%d) instead.", newLocation.x, newLocation.y);
	}
	newLocation = navigator.locate('<', location{0, 0})
	if newLocation.x != -1 || newLocation.y != 0 {
		t.Error("Expected to be at (-1,0), but are at (%d,%d) instead.", newLocation.x, newLocation.y);
	}
	newLocation = navigator.locate('>', location{0, 0})
	if newLocation.x != 1 || newLocation.y != 0 {
		t.Error("Expected to be at (1,0), but are at (%d,%d) instead.", newLocation.x, newLocation.y);
	}
}

func Test_it_stores_visited_locations_and_their_counts(t *testing.T) {
	navigator := NewGps();
	navigator.Move('>');
	navigator.Move('>');
	navigator.Move('<');
	if navigator.visited[location{0, 0}] != 1 {
		t.Error("Expected to have visited the starting position");
	} else if navigator.visited[location{1, 0}] != 2 {
		t.Error("Expected to have visited (1,0) twice, got", navigator.visited[location{1, 0}]);
	} else if navigator.visited[location{2, 0}] != 1 {
		t.Error("Expected to have visited (2,0) once, got", navigator.visited[location{2, 0}]);
	}
}

func Test_it_returns_number_of_visited_locations(t *testing.T) {
	navigator := NewGps();
	navigator.Move('>');
	navigator.Move('>');
	navigator.Move('<');
	if navigator.VisitedPlaces() != 3 {
		t.Errorf("Expected to have visited 3 places, but reported %d instead.", navigator.VisitedPlaces());
	}
}

func Test_it_parses_sequence_of_moves(t *testing.T) {
	navigator := NewGps();
	navigator.Route(">><");
	if navigator.visited[location{0, 0}] != 1 {
		t.Error("Expected to have visited the starting position");
	} else if navigator.visited[location{1, 0}] != 2 {
		t.Error("Expected to have visited (1,0) twice, got", navigator.visited[location{1, 0}]);
	} else if navigator.visited[location{2, 0}] != 1 {
		t.Error("Expected to have visited (2,0) once, got", navigator.visited[location{2, 0}]);
	} else if navigator.VisitedPlaces() != 3 {
		t.Errorf("Expected to have visited 3 places, but reported %d instead.", navigator.VisitedPlaces());
	}
}

func Test_sample_input(t *testing.T) {
	navigator := NewGps();
	navigator.Route("^>v<");
	if navigator.VisitedPlaces() != 4 {
		t.Errorf("Expected to have visited 4 places, but reported %d instead.", navigator.VisitedPlaces());
	}
}
