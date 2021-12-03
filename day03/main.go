package main

import (
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
	gamma := 0
	epsilon := 0
	sum := 0
	for i := 0; i < len(inp[0]); i++ {
		o := 0
		z := 0
		for _, v := range inp {
			if v[i] == '1' {
				o++
			} else {
				z++
			}
		}
		if o > z {
			gamma = (gamma * 2) + 1
			epsilon *= 2
		} else {
			gamma *= 2
			epsilon = (epsilon * 2) + 1
		}
	}
	sum = gamma * epsilon
	c.Print("Part1: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
func part2() {
	start := time.Now()
	inp := c.ReadInputFile()
	oxygen := 0
	co2 := 0
	sum := 0
	for i := 0; i < len(inp[0]); i++ {
		o := 0
		z := 0
		for _, v := range inp {
			if v[i] == '1' {
				o++
			} else {
				z++
			}
		}
		if o < z {
			inp = keep(inp, i, false)
		} else {
			inp = keep(inp, i, true)
		}
		//c.Print(inp)
		if len(inp) == 1 {
			break
		}
	}
	oxygen = bin2num(inp[0])
	inp = c.ReadInputFile()
	for i := 0; i < len(inp[0]); i++ {
		o := 0
		z := 0
		for _, v := range inp {
			if v[i] == '1' {
				o++
			} else {
				z++
			}
		}
		if z <= o {
			inp = keep(inp, i, false)
		} else {
			inp = keep(inp, i, true)
		}
		//c.Print(inp)
		if len(inp) == 1 {
			break
		}
	}
	co2 = bin2num(inp[0])
	sum = oxygen * co2

	c.Print("Part2: ", sum)
	c.Print("Took: ", time.Since(start).String())
}
func keep(data []string, pos int, one bool) []string {
	res := []string{}
	for _, v := range data {
		if (v[pos] == '1' && one) || (v[pos] == '0' && !one) {
			res = append(res, v)
		}
	}
	return res
}
func bin2num(data string) int {
	res := 0
	for i := 0; i < len(data); i++ {
		res *= 2
		if data[i] == '1' {
			res++
		}
	}
	return res
}
