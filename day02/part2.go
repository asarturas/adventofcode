package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day02/solution"
	"fmt"
)

func main() {
	calculator := solution.Calculator{};
	allRibbon := 0;
	for input := io.ReadLine(); input != ""; input = io.ReadLine() {
		allRibbon += calculator.RibbonForGift(calculator.GiftDimensions(input));
	}
	fmt.Println(allRibbon);
}
