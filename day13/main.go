package main

import (
	"fmt"
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

	sum := 0
	m, ins := readData()
	for _, v := range ins {
		a := strings.Split(v, "=")
		l := c.GetInt(a[1])
		m, _, _ = fold(m, a[0], l)
		break
	}
	for x := range m {
		sum += len(m[x])
	}

	return sum
}
func part2() int {
	m, ins := readData()
	w := 0
	h := 0
	for _, v := range ins {
		a := strings.Split(v, "=")
		l := c.GetInt(a[1])
		m, w, h = fold(m, a[0], l)

	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if _, ok := m[x][y]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return 0
}
func readData() (map[int]map[int]bool, []string) {
	inp := c.ReadInputFile()

	m := map[int]map[int]bool{}
	ins := []string{}
	cMode := true
	for _, v := range inp {
		if v == "" {
			cMode = false
			continue
		}
		if cMode {
			d := c.GetInts(strings.Split(v, ","))
			if _, ok := m[d[0]]; !ok {
				m[d[0]] = map[int]bool{}
			}
			m[d[0]][d[1]] = true
		} else {
			ins = append(ins, strings.Replace(v, "fold along ", "", 1))
		}
	}
	return m, ins
}
func fold(paper map[int]map[int]bool, along string, line int) (map[int]map[int]bool, int, int) {
	w := 0
	h := 0
	if along == "x" {
		w = line
		for x := range paper {
			for y := range paper[x] {
				if x == line {
					delete(paper[x], y)
				} else if x > line {
					if _, ok := paper[line-(x-line)]; !ok {
						paper[line-(x-line)] = map[int]bool{}
					}
					paper[line-(x-line)][y] = true
					delete(paper[x], y)
				}
				if y > h {
					h = y
				}
			}
		}
	} else {
		h = line
		for x := range paper {
			for y := range paper[x] {
				if y == line {
					delete(paper[x], y)
				} else if y > line {
					paper[x][line-(y-line)] = true
					delete(paper[x], y)
				}
			}
			if x > w {
				w = x
			}
		}
	}
	return paper, w, h
}
