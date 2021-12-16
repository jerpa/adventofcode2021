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
	c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}

func part1() int {
	data := loadData()
	sum := 0
	sum, _, _ = parse(data, 0, len(data), -1)

	return sum
}
func parse(data []bool, start, stop, numberOfPackets int) (sum, newStop int, res []int) {
	res = []int{}
	typ := -1
	for idx := 0; start < stop-6 && (numberOfPackets == -1 || idx < numberOfPackets); idx++ {
		ver := bin2num(data[start : start+3])
		sum += ver
		typ = bin2num(data[start+3 : start+6])

		if typ == 4 {
			halt := false
			r := []bool{}
			i := 6
			for i = start + 6; i < len(data) && !halt; i += 5 {
				halt = !data[i]
				r = append(r, data[i+1:i+5]...)
			}
			res = append(res, bin2num(r))
			start = i
		} else {

			if data[start+6] {
				n := bin2num(data[start+7 : start+18])
				s := 0
				r := []int{}
				s, newStop, r = parse(data, start+18, stop, n)
				res = append(res, checkType(typ, r)...)
				if newStop > start {
					start = newStop
				}
				sum += s
			} else {
				n := 0
				if start < stop-22 {
					n = bin2num(data[start+7 : start+22])
					s := 0
					r := []int{}
					s, _, r = parse(data, start+22, start+22+n, -1)
					res = append(res, checkType(typ, r)...)
					sum += s
				}
				start = start + 22 + n
			}
		}
	}
	newStop = start
	return
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
func checkType(typ int, data []int) []int {
	res := []int{}
	if typ == 0 {
		s := 0
		for _, v := range data {
			s += v
		}
		res = append(res, s)
	} else if typ == 1 {
		s := 1
		for _, v := range data {
			s *= v
		}
		res = append(res, s)
	} else if typ == 2 {
		res = append(res, c.MinInt(data))
	} else if typ == 3 {
		res = append(res, c.MaxInt(data))
	} else if typ == 5 {
		if data[0] > data[1] {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	} else if typ == 6 {
		if data[0] < data[1] {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	} else if typ == 7 {
		if data[0] == data[1] {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}
	return res
}
func part2() int {
	data := loadData()

	_, _, sum := parse(data, 0, len(data), -1)

	return sum[0]
}
func loadData() []bool {
	inp := c.ReadInputFile()
	data := []bool{}
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
	return data
}
