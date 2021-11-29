package main

import (
	"time"

	c "github.com/jerpa/adventofcode2021/common"
)

func main() {
	part1()
	part2()

}
func part1() {
	start := time.Now()
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	c.Print("Part1: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
func part2() {
	start := time.Now()
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	c.Print("Part2: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
