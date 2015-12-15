package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day01/solution"
	"fmt"
)

func main() {
	counter := solution.Counter{};
	input := io.ReadLine();
	floor := counter.Floor(input);
	fmt.Println(floor);
}
