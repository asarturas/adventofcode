package solution
import (
	"strings"
	"strconv"
)

type Calculator struct {}

func (this Calculator) GiftDimensions(dimensions string) (int, int, int) {
	threeDimensions := strings.SplitN(dimensions, "x", 3)
	length, _ := strconv.Atoi(threeDimensions[0])
	width, _ := strconv.Atoi(threeDimensions[1])
	height, _ := strconv.Atoi(threeDimensions[2])
	return length, width, height
}

func (this Calculator) PaperForGift(length, width, height int) int {
	return this.wrapping(length, width, height) + this.slack(length, width, height)
}

func (this Calculator) wrapping(length, width, height int) int {
	return 2 * length * width + 2 * width * height + 2 * height * length;
}

func (this Calculator) slack(sides ...int) int {
	if len(sides) < 2 {
		return 0;
	}
	if len(sides) >= 3 {
		sides = this.insertionSort(sides)
	}
	return sides[0] * sides[1];
}

func (this Calculator) RibbonForGift(length, width, height int) int {
	return this.ribbon(length, width, height) + this.bow(length, width, height)
}

func (this Calculator) ribbon(sides ...int) int {
	if len(sides) < 2 {
		return 0;
	}
	if len(sides) >= 3 {
		sides = this.insertionSort(sides)
	}
	return sides[0] * 2 + sides[1] * 2;
}

func (this Calculator) bow(length, width, height int) int {
	return length * width * height;
}

func (this Calculator) insertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i; j > 0 && numbers[j - 1] > numbers[j]; j-- {
			tmp := numbers[j];
			numbers[j] = numbers[j - 1];
			numbers[j - 1] = tmp;
		}
	}
	return numbers;
}
