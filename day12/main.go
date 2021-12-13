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
	conns := readData()
	r := path(conns, "start", "start")

	return len(r)
}
func part2() int {
	conns := readData()
	r := path2(conns, "start", "start", false)
	return len(r)
}
func path(conns map[string]map[string]bool, start, current string) []string {
	res := []string{}
	if start == "end" {
		return []string{current}
	}
	for k := range conns[start] {
		if strings.ToLower(k) == k && strings.Contains(current, ","+k+",") {
			continue
		}
		if strings.Contains(current, ","+start+","+k+",") {
			continue
		}
		res = append(res, path(conns, k, current+","+k)...)
	}
	return res
}

func path2(conns map[string]map[string]bool, start, current string, hasTwice bool) []string {
	res := []string{}
	if start == "end" {
		return []string{current}
	}
	for k := range conns[start] {
		tw := hasTwice
		if strings.ToLower(k) == k && CountSubString(current, ","+k+",") >= 1 {
			if tw {
				continue
			}
			tw = true
		}
		res = append(res, path2(conns, k, current+","+k, tw)...)
	}

	return res
}
func CountSubString(data, substring string) int {
	sum := 0
	for pos := 0; pos < len(data)-len(substring); pos++ {
		if data[pos:pos+len(substring)] == substring {
			sum++
		}
	}
	return sum
}
func readData() map[string]map[string]bool {
	inp := c.ReadInputFile()

	conns := map[string]map[string]bool{}
	for _, v := range inp {
		s := strings.Split(v, "-")
		if _, ok := conns[s[0]]; !ok {
			conns[s[0]] = make(map[string]bool)
		}
		if s[1] != "start" {
			conns[s[0]][s[1]] = true
		}
		if _, ok := conns[s[1]]; !ok {
			conns[s[1]] = make(map[string]bool)
		}
		if s[0] != "start" {
			conns[s[1]][s[0]] = true
		}
	}
	return conns
}
