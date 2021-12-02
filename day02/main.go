package main

import (
	"strings"
	"time"

	c "github.com/jerpa/adventofcode2021/common"
)

func main() {
	part1()
	part2()

}
func part1() {
	start := time.Now()
	inp := c.ReadInputFile()

	sum := 0
	h := 0
	v := 0
	for _, r := range inp {
		s := strings.Split(r, " ")
		val := c.GetInt(s[1])
		if s[0] == "forward" {
			h += val
		} else if s[0] == "up" {
			v -= val
		} else {
			v += val
		}
	}
	sum = v * h

	c.Print("Part1: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
func part2() {
	start := time.Now()
	inp := c.ReadInputFile()

	sum := 0
	h := 0
	v := 0
	aim := 0
	for _, r := range inp {
		s := strings.Split(r, " ")
		val := c.GetInt(s[1])
		if s[0] == "forward" {
			h += val
			v += val * aim
		} else if s[0] == "up" {
			aim -= val
		} else {
			aim += val
		}
	}
	sum = v * h

	c.Print("Part2: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
