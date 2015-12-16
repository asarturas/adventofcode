package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day03/solution"
	"fmt"
)

func main() {
	santa := solution.NewGps();
	robo := solution.NewGps();
	for pos, move := range io.ReadLine() {
		if pos % 2 == 0 {
			santa.Move(move);
		} else {
			robo.Move(move);
		}
	}
	santa.Add(robo);
	fmt.Println(santa.VisitedPlaces());
}
