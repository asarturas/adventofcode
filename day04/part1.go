package main

import(
	"github.com/asarturas/adventofcode/io"
	"github.com/asarturas/adventofcode/day04/solution"
	"fmt"
)

func main() {
	miner := &solution.Miner{solution.Md5Hasher{}};
	_, number := miner.Mine(io.ReadLine());
	fmt.Println(number);
}
