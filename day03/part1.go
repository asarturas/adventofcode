package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day03/solution"
	"fmt"
)

func main() {
	navigator := solution.NewGps();
	navigator.Route(io.ReadLine());
	fmt.Println(navigator.VisitedPlaces());
}
