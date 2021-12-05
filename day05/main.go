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
	return checkMap(false)
}
func part2() int {
	return checkMap(true)
}

func checkMap(checkDiagonal bool) int {
	inp := c.ReadInputFile()

	sum := 0
	m := map[int]map[int]int{}

	for _, v := range inp {
		p := strings.Split(v, " -> ")
		start := strings.Split(p[0], ",")
		stop := strings.Split(p[1], ",")
		startX := c.GetInt(start[0])
		startY := c.GetInt(start[1])
		stopX := c.GetInt(stop[0])
		stopY := c.GetInt(stop[1])

		if startX == stopX {
			if startY > stopY {
				stopY, startY = startY, stopY
			}
			for y := startY; y <= stopY; y++ {
				if _, ok := m[startX]; !ok {
					m[startX] = map[int]int{}
				}
				m[startX][y]++
			}
		} else if startY == stopY {
			if startX > stopX {
				stopX, startX = startX, stopX
			}
			for x := startX; x <= stopX; x++ {
				if _, ok := m[x]; !ok {
					m[x] = map[int]int{}
				}
				m[x][startY]++
			}
		} else if checkDiagonal {
			if startX > stopX {
				stopX, startX = startX, stopX
				stopY, startY = startY, stopY
			}

			y := startY
			for x := startX; x <= stopX; x++ {
				if _, ok := m[x]; !ok {
					m[x] = map[int]int{}
				}
				m[x][y]++
				if startY > stopY {
					y--
				} else {
					y++
				}
			}
		}
	}
	for kx, x := range m {
		for ky := range x {
			if m[kx][ky] >= 2 {
				sum++
			}
		}
	}
	return sum

}
