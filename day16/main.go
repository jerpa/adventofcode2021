package main

import (
	"strconv"
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

func part1() int {
	inp := c.ReadInputFile()
	data := []bool{}
	sum := 0
	for _, v := range inp[0] {
		n, err := strconv.ParseInt(string(v), 16, 64)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, n&0x8 == 0x8)
		data = append(data, n&0x4 == 0x4)
		data = append(data, n&0x2 == 0x2)
		data = append(data, n&0x1 == 0x1)

	}
	v, t, sum := parse(data)
	//sum = bin2num(data[0:3])
	return sum + v + t
}
func parse(data []bool) (int, int, int) {
	ver := bin2num(data[0:3])
	typ := bin2num(data[3:6])
	stop := false
	res := []bool{}
	for i := 6; i < len(data) && !stop; i += 5 {
		stop = !data[i]
		res = append(res, data[i+1:i+5]...)
	}

	return ver, typ, bin2num(res)
}
func bin2num(data []bool) int {
	res := 0
	for _, v := range data {
		res *= 2
		if v {
			res += 1
		}
	}
	return res
}
func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	for _, v := range c.GetInts(inp) {
		sum += v
	}

	return sum
}
