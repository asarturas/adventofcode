package solution
import (
	"strings"
	"math"
)

type Evaluator interface {
	IsNice(word string) bool;
}

type Vowels struct {}

func (this Vowels) IsNice(word string) bool {
	nr := 0;
	for _, char := range word {
		if strings.ContainsRune("aeiou", char) {
			nr++;
		}
		if nr == 3 {
			return true;
		}
	}
	return false;
}

type Duplicates struct {}

func (this Duplicates) IsNice(word string) bool {
	previousChar := rune(math.MaxInt8);
	for _, char := range word {
		if previousChar == char {
			return true;
		}
		previousChar = char;
	}
	return false;
}

type Restricted struct {}

func (this Restricted) IsNice(word string) bool {
	allRestricted := []string{"ab", "cd", "pq", "xy"};
	for _, restricted := range allRestricted {
		if strings.Contains(word, restricted) {
			return false;
		}
	}
	return true;
}

type Exclusive struct {
	Evaluators []Evaluator
}

func (this *Exclusive) IsNice(word string) bool {
	for _, evaluator := range this.Evaluators {
		if !evaluator.IsNice(word) {
			return false;
		}
	}
	return true;
}

type Pairs struct{};

func (this Pairs) IsNice(word string) bool {
	for i := 0; i < len(word) - 1; i++ {
		start := i;
		end := i + 2;
		pair := word[start:end];
		if strings.Contains(word[end:], pair) {
			return true;
		}
	}
	return false;
}

type SplitRepeat struct{};

func (this SplitRepeat) IsNice(word string) bool {
	for i := 0; i < len(word) - 2; i++ {
		if word[i] == word[i + 2] {
			return true;
		}
	}
	return false;
}
