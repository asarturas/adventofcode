package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day01/solution"
	"fmt"
)

func main() {
	counter := solution.Counter{};
	input := io.ReadLine();
	err, basementPosition := counter.FirstBasementPosition(input);
	if err == nil {
		fmt.Println(basementPosition);
	}
}
