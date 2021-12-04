package main

import (
	"time"

	c "github.com/jerpa/adventofcode2021/common"
)

func main() {
	start := time.Now()
	c.Print("Part1: ", part1())
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	//c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

func part1() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	return sum
}
