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
	inp := strings.Split(c.ReadInputFile()[0], ",")
	val := c.GetInts(inp)

	sum := -1
	for i := c.MinInt(val); i < c.MaxInt(val); i++ {
		s := 0
		for _, v := range val {
			s += c.AbsInt(i, v)
		}
		if sum == -1 || sum > s {
			sum = s
		}
	}

	return sum
}
func part2() int {
	inp := strings.Split(c.ReadInputFile()[0], ",")
	val := c.GetInts(inp)

	sum := -1
	for i := c.MinInt(val); i < c.MaxInt(val); i++ {
		s := 0
		for _, v := range val {
			s += c.ConsecutiveSum(c.AbsInt(i, v))
		}
		if sum == -1 || sum > s {
			sum = s
		}
	}

	return sum
}
