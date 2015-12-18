package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day05/solution"
	"fmt"
)

func main() {
	santaHelper := &solution.Exclusive{
		[]solution.Evaluator{
			solution.Pairs{},
			solution.SplitRepeat{},
		},
	}
	numberOfNiceOnes := 0;
	for name := io.ReadLine(); name != ""; name = io.ReadLine() {
		if santaHelper.IsNice(name) {
			numberOfNiceOnes++;
		}
	}
	fmt.Println(numberOfNiceOnes);
}
