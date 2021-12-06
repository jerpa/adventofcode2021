package main

import (
	"strings"
	"time"

	c "github.com/jerpa/adventofcode2021/common"
)

func main() {
	start := time.Now()
	c.Print("Part1: ", part1())
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

func part1() int {
	inp := c.ReadInputFile()
	val := c.GetInts(strings.Split(inp[0], ","))

	for day := 0; day < 80; day++ {

		for i := range val {
			if val[i] == 0 {
				val = append(val, 8)
				val[i] = 7
			}
			val[i]--
		}
	}

	return len(val)
}
func part2() int {
	inp := c.ReadInputFile()
	val := c.GetInts(strings.Split(inp[0], ","))

	for day := 0; day < 80; day++ {

		for i := range val {
			if val[i] == 0 {
				val = append(val, 8)
				val[i] = 7
			}
			val[i]--
		}
	}

	return len(val)
}
