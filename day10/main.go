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

	sum := 0

	for _, v := range inp {
		exp := []string{}
		for _, s := range strings.Split(v, "") {
			if s == "(" {
				exp = append(exp, ")")
			} else if s == "[" {
				exp = append(exp, "]")
			} else if s == "{" {
				exp = append(exp, "}")
			} else if s == "<" {
				exp = append(exp, ">")
			} else if len(exp) > 0 && exp[len(exp)-1] == s {
				// Correct character, remove last element from expected
				exp = exp[:len(exp)-1]
			} else {
				switch s {
				case ")":
					sum += 3
				case "]":
					sum += 57
				case "}":
					sum += 1197
				case ">":
					sum += 25137
				}
				break
			}
		}

	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	points := []int{}
	for _, v := range inp {
		exp := []string{}
		corrupt := false
		for _, s := range strings.Split(v, "") {
			if s == "(" {
				exp = append(exp, ")")
			} else if s == "[" {
				exp = append(exp, "]")
			} else if s == "{" {
				exp = append(exp, "}")
			} else if s == "<" {
				exp = append(exp, ">")
			} else if len(exp) > 0 && exp[len(exp)-1] == s {
				exp = exp[:len(exp)-1]
			} else {
				corrupt = true
				break
			}
		}
		if !corrupt && len(exp) > 0 {
			p := 0
			for i := len(exp) - 1; i >= 0; i-- {
				s := exp[i]
				switch s {
				case ")":
					p = (p * 5) + 1
				case "]":
					p = (p * 5) + 2
				case "}":
					p = (p * 5) + 3
				case ">":
					p = (p * 5) + 4

				}
			}
			points = append(points, p)
		}

	}
	sort.Ints(points)

	return points[len(points)/2]
}
