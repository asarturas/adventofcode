package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day02/solution"
	"fmt"
)

func main() {
	calculator := solution.Calculator{};
	allPaper := 0;
	for input := io.ReadLine(); input != ""; input = io.ReadLine() {
		allPaper += calculator.PaperForGift(calculator.GiftDimensions(input));
	}
	fmt.Println(allPaper);
}
