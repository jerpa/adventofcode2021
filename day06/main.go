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
	// part1 is way too slow and uses so much memory that I had to restart it... So this solution uses groups instead
	inp := c.ReadInputFile()
	val := c.GetInts(strings.Split(inp[0], ","))
	res := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0}

	for _, v := range val {
		res[v]++
	}

	for day := 0; day < 256; day++ {
		for p := 0; p <= 8; p++ {
			if p == 0 {
				res[7] += res[0]
				res[9] = res[0]
			}
			res[p] = res[p+1]

		}
	}
	sum := 0
	for i := 0; i < 9; i++ {
		sum += res[i]
	}

	return sum
}
