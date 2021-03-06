package solution

import (
	"github.com/asarturas/adventofcode/Godeps/_workspace/src/github.com/golang/mock/gomock"
	"testing"
)

func Test_it_is_nice_when_it_has_at_lest_three_vowels(t *testing.T) {
	evaluator := Vowels{}
	if !evaluator.IsNice("asei") {
		t.Error("Expected 'asei' to be nice.")
	}
}

func Test_it_is_not_nice_when_it_has_less_than_three_vowels(t *testing.T) {
	evaluator := Vowels{}
	if evaluator.IsNice("asasss") {
		t.Error("Expected 'asasss' not to be nice.")
	}
}

func Test_it_is_nice_when_it_has_same_letter_twice_in_a_row(t *testing.T) {
	evaluator := Duplicates{}
	if !evaluator.IsNice("assa") {
		t.Error("Expected 'ssa' to be nice.")
	}
}

func Test_it_is_not_nice_when_it_does_not_have_diplicates_going_one_after_another(t *testing.T) {
	evaluator := Duplicates{}
	if evaluator.IsNice("asas") {
		t.Error("Expected 'asas' not to be nice.")
	}
}

func Test_it_is_not_nice_when_it_has_restricted_strings(t *testing.T) {
	evaluator := Restricted{}
	if evaluator.IsNice("abc") {
		t.Error("Expected 'abc' not to be nice.")
	}
}

func Test_it_is_nice_when_it_does_not_have_restricted_strings(t *testing.T) {
	evaluator := Restricted{}
	if !evaluator.IsNice("pcacgxz") {
		t.Error("Expected 'pcacgxz' to be nice.")
	}
}

func Test_exclusive_stops_on_first_success(t *testing.T) {
	controller := gomock.NewController(t)
	oneEvaluator := NewMockEvaluator(controller)
	oneEvaluator.EXPECT().IsNice("aaabb").Return(true)
	anotherEvaluator := NewMockEvaluator(controller)
	anotherEvaluator.EXPECT().IsNice("aaabb").Return(false)
	thirdEvaluator := NewMockEvaluator(controller)
	thirdEvaluator.EXPECT().IsNice("aaabb").Times(0)
	evaluators := []Evaluator{oneEvaluator, anotherEvaluator, thirdEvaluator}
	evaluator := &Exclusive{evaluators}
	if evaluator.IsNice("aaabb") {
		t.Error("Expected 'aaaabb' not to be nice.")
	}
}

func Test_exclusive_returns_true_if_all_evaluators_returned_true(t *testing.T) {
	controller := gomock.NewController(t)
	oneEvaluator := NewMockEvaluator(controller)
	oneEvaluator.EXPECT().IsNice("aaabb").Return(true)
	anotherEvaluator := NewMockEvaluator(controller)
	anotherEvaluator.EXPECT().IsNice("aaabb").Return(true)
	evaluators := []Evaluator{oneEvaluator, anotherEvaluator}
	evaluator := &Exclusive{evaluators}
	if !evaluator.IsNice("aaabb") {
		t.Error("Expected 'aaaabb' to be nice.")
	}
}

func Test_it_is_nice_when_it_has_non_overlapping_pair_going_one_after_another(t *testing.T) {
	evaluator := Pairs{}
	if !evaluator.IsNice("agag") {
		t.Error("Expected 'agag' to be nice.")
	}
}

func Test_it_is_not_nice_when_it_does_not_have_non_overlapping_repeating_pair(t *testing.T) {
	evaluator := Pairs{}
	if evaluator.IsNice("aggga") {
		t.Error("Expected 'aggga' not to be nice.")
	}
}

func Test_it_is_nice_when_it_has_repeated_letter_with_other_letter_in_between(t *testing.T) {
	evaluator := SplitRepeat{}
	if !evaluator.IsNice("aca") {
		t.Error("Expected 'aca' to be nice.")
	}
}

func Test_it_is_not_nice_when_it_does_not_have_repeated_letter_with_other_letter_in_between(t *testing.T) {
	evaluator := SplitRepeat{}
	if evaluator.IsNice("acbaad") {
		t.Error("Expected 'acbaad' not to be nice.")
	}
}
