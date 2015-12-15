package solution

import (
	"testing"
)

var calculator = Calculator{};

func Test_wrapping_dimensions_for_gift_of_2_on_3_on_4(t *testing.T) {
	paper := calculator.PaperForGift(2, 3, 4)
	if paper != 58 {
		t.Errorf("Unexpected amount of paper: %s", paper);
	}
}

func Test_wrapping_dimensions_for_gift_of_1_on_1_on_10(t *testing.T) {
	paper := calculator.PaperForGift(1, 1, 10)
	if paper != 43 {
		t.Errorf("Unexpected amount of paper: %s", paper);
	}
}

func Test_wrapping_dimensions_for_gift_with_parsing_of_size(t *testing.T) {
	length, width, height := calculator.GiftDimensions("2x3x4")
	if length != 2 {
		t.Errorf("Expected length to be 2, but got %d", length);
	} else if width != 3 {
		t.Errorf("Expected width to be 3, but got %d", width);
	} else if height != 4 {
		t.Errorf("Expected height to be 4, but got %d", height);
	}
}

func Test_ribbon_length_for_gift_of_2_on_3_on_4(t *testing.T) {
	ribbon := calculator.RibbonForGift(2, 3, 4)
	if ribbon != 34 {
		t.Errorf("Unexpected amount of ribbon: %s", ribbon);
	}
}
