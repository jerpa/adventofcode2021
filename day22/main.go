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
	//c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

type rnge struct {
	from, to int
}

func part1() int {
	inp := c.ReadInputFile()

	m := map[int]map[int]map[int]bool{}
	sum := 0
	for _, v := range inp {
		s := strings.Split(v, " ")
		onoff := s[0] == "on"
		ranges := []rnge{}
		skip := false
		for _, d := range strings.Split(s[1], ",") {

			n := c.GetInts(strings.Split(d[2:], ".."))
			r := rnge{from: c.MinInt(n), to: c.MaxInt(n)}
			if r.to < (-50) || r.from > 50 {
				skip = true
				break
			}
			if r.from < (-50) {
				r.from = -50
			}
			if r.to > 50 {
				r.to = 50
			}
			ranges = append(ranges, r)
		}
		if skip {
			continue
		}
		for z := ranges[2].from; z <= ranges[2].to; z++ {
			if z < (-50) || z > 50 {
				continue
			}
			for y := ranges[1].from; y <= ranges[1].to; y++ {
				if y < (-50) || y > 50 {
					continue
				}
				if _, ok := m[z]; !ok {
					m[z] = map[int]map[int]bool{}
				}
				for x := ranges[0].from; x <= ranges[0].to; x++ {
					if x < (-50) || y > 50 {
						continue
					}
					if _, ok := m[z][y]; !ok {
						m[z][y] = map[int]bool{}
					}
					m[z][y][x] = onoff
				}
			}
		}
	}
	for z := range m {
		for y := range m[z] {
			for x := range m[z][y] {
				if m[z][y][x] {
					sum++
				}
			}
		}
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
