package main

import (
	"sort"
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
	m := map[int]map[int]int{}
	// Build map
	for y := range inp {
		m[y] = map[int]int{}
		for x, v := range c.GetInts(strings.Split(inp[y], "")) {
			m[y][x] = v
		}
	}

	sum := 0
	for y := range m {
		for x := range m[y] {
			// Extract values of neighbours
			n := []int{}
			if _, ok := m[y-1]; ok {
				n = append(n, m[y-1][x])
			}
			if _, ok := m[y+1]; ok {
				n = append(n, m[y+1][x])
			}
			if _, ok := m[y][x-1]; ok {
				n = append(n, m[y][x-1])
			}
			if _, ok := m[y][x+1]; ok {
				n = append(n, m[y][x+1])
			}
			// Check if current value is lower than neighbours
			lowest := true
			for _, v := range n {
				if v <= m[y][x] {
					lowest = false
					break
				}
			}
			if lowest {
				sum += m[y][x] + 1
			}
		}
	}

	return sum
}

type point struct {
	val   int
	basin int
}

func part2() int {
	inp := c.ReadInputFile()
	m := map[int]map[int]*point{}
	// Build map
	for y := range inp {
		m[y] = map[int]*point{}
		for x, v := range c.GetInts(strings.Split(inp[y], "")) {
			m[y][x] = &point{val: v}
		}
	}
	// Find basins
	basid := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x].val == 9 {
				continue
			}
			// Check if one neighbour is part of a basin
			nbasid := -1
			if _, ok := m[y-1]; ok && m[y-1][x].basin != 0 {
				nbasid = m[y-1][x].basin
			}
			if _, ok := m[y+1]; ok && m[y+1][x].basin != 0 {
				nbasid = m[y+1][x].basin
			}
			if _, ok := m[y][x-1]; ok && m[y][x-1].basin != 0 {
				nbasid = m[y][x-1].basin
			}
			if _, ok := m[y][x+1]; ok && m[y][x+1].basin != 0 {
				nbasid = m[y][x+1].basin
			}
			// One neighbour has a basin
			if nbasid != -1 {
				m[y][x].basin = nbasid
			} else {
				// Otherwise create a new one
				basid++
				m[y][x].basin = basid
			}
		}
	}
	// Merge basins by checking if the one below belongs to another basin, then invite that one to current
	for y := 0; y < len(m)-1; y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x].basin != 0 && m[y+1][x].basin != 0 && m[y][x].basin != m[y+1][x].basin {
				n := m[y+1][x].basin // Replace this basin
				t := m[y][x].basin   // With this basin
				for my := range m {
					for mx := range m[my] {
						if m[my][mx].basin == n {
							m[my][mx].basin = t
						}
					}
				}
			}

		}

	}
	// Count the size of each basin
	basins := []int{}
	for i := 0; i <= basid; i++ {
		basins = append(basins, 0)
	}
	for y := range m {
		for x := range m[y] {
			if m[y][x].basin > 0 {
				basins[m[y][x].basin]++
			}
		}
	}
	// Sort the basins to find the largest three
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]

}
