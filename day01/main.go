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
	prev := -1
	for _, v := range c.GetInts(inp) {
		if prev != -1 && v > prev {
			sum++
		}
		prev = v
	}

	c.Print("Part1: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
func part2() {
	start := time.Now()
	inp := c.ReadInputFile()

	sum := 0
	prev := -1
	data := c.GetInts(inp)
	for i := 2; i < len(data); i++ {
		val := data[i] + data[i-1] + data[i-2]
		if prev != -1 && val > prev {
			sum++
		}
		prev = val
	}

	c.Print("Part2: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
