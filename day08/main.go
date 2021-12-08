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
		g := strings.Split(v, " | ")
		for _, s := range strings.Split(g[1], " ") {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				sum++
			}
		}
	}

	return sum
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range inp {

		g := strings.Split(v, " | ")
		/*
			0=6
			1=2
			2=5
			3=5
			4=4
			5=5
			6=6
			7=3
			8=7
			9=6
		*/
		nums := strings.Split(g[0], " ")
		for i := range nums {
			a := strings.Split(nums[i], "")
			sort.Strings(a)
			nums[i] = strings.Join(a, "")
		}
		w := map[string]int{"abcdefg": 8}
		n := map[int]string{8: "abcdefg"}

		//		var t, m, b, tl, tr, bl, br string
		for len(n) != 10 {
			for _, v := range nums {
				if len(v) == 2 {
					w[v] = 1
					n[1] = v
				} else if len(v) == 3 {
					w[v] = 7
					n[7] = v
				} else if len(v) == 4 {
					w[v] = 4
					n[4] = v
				} else if len(v) == 5 {
					if len(remove(n[7], v)) == 2 {
						w[v] = 3
						n[3] = v
					} else {
						if len(remove(n[4], v)) == 2 {
							w[v] = 5
							n[5] = v
						} else {
							w[v] = 2 // Wrong
							n[2] = v // Wrong
						}
					}
				} else if len(v) == 6 {
					if len(remove(n[4], remove(n[7], v))) == 1 {
						w[v] = 9
						n[9] = v
					} else if len(remove(n[1], v)) == 5 {
						w[v] = 6
						n[6] = v
					} else {
						w[v] = 0
						n[0] = v
					}

				}
			}
		}
		return -1

		for _, s := range strings.Split(g[1], " ") {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				sum++
			}
		}
	}

	return sum
}
func getMissing(longStr, shortStr string) string {
	for i := range longStr {
		if strings.Contains(shortStr, string(longStr[i])) == false {
			return string(longStr[i])
		}
	}
	return ""
}
func remove(str, from string) string {
	for i := range str {
		from = strings.ReplaceAll(string(str[i]), from, "")
	}
	return from
}
