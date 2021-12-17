package main

import (
	"regexp"
	"time"

	c "github.com/jerpa/adventofcode2021/common"
)

type target struct {
	x1, x2, y1, y2 int
}

func main() {
	start := time.Now()
	c.Print("Part1: ", part1())
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

func part1() int {
	t := loadData()
	b := 0
	for x := 1; x < t.x1; x++ {
		for y := 0; y < 200; y++ { // Only throw up since we want the highest point
			v, m := throw(x, y, t)
			if !m && v > b {
				b = v
			}
		}
	}
	return b
}
func part2() int {
	t := loadData()
	sum := 0
	for x := 1; x <= t.x2; x++ {
		for y := t.y1; y < 200; y++ {
			_, m := throw(x, y, t)
			if !m {
				sum++

			}
		}
	}
	return sum
}
func loadData() target {
	inp := c.ReadInputFile()
	re := regexp.MustCompile(`x=([\-\d]*)..([-\d]*),\sy=([-\d]*)..([-\d]*)`)
	s := re.FindStringSubmatch(inp[0])
	m := c.GetInts(s[1:5])
	return target{x1: c.MinInt(m[0:2]), x2: c.MaxInt(m[0:2]), y1: c.MinInt(m[2:4]), y2: c.MaxInt(m[2:4])}
}
func throw(xv, yv int, t target) (maxY int, miss bool) {
	x, y := 0, 0

	for {

		m, o := inTarget(x, y, t)

		if o {
			miss = true
			return
		}
		if m {
			return
		}
		x += xv
		y += yv
		if y > maxY {
			maxY = y
		}
		if xv > 0 {
			xv--
		}
		yv--
	}
}
func inTarget(x, y int, t target) (bool, bool) {
	if x > t.x2 || y < t.y1 {
		return false, true
	}
	return x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2, false
}
