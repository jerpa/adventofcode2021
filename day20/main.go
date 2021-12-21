package main

import (
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
	key, data := loadData()

	for l := 0; l < 2; l++ {
		res := []string{}
		res = append(res, data[0])
		for y := 1; y < len(data)-1; y++ {
			row := "."
			for x := 1; x < len(data[y])-1; x++ {
				g := data[y-1][x-1 : x+2]
				g += data[y][x-1 : x+2]
				g += data[y+1][x-1 : x+2]
				n := str2bin2dec(g)
				row += string(key[n])
			}
			row += "."
			res = append(res, row)
		}
		res = append(res, data[0])
		data = res
	}

	sum := 0
	for y := 4; y < len(data)-4; y++ {
		for x := 4; x < len(data[y])-4; x++ {
			if data[y][x] == '#' {
				sum++
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
func str2bin2dec(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 2
		if s[i] == '#' {
			res += 1
		}
	}
	return res
}

func loadData() (string, []string) {
	inp := c.ReadInputFile()
	key := inp[0]
	inp = inp[2:]
	for i := range inp {
		inp[i] = "............" + inp[i] + "............"
	}
	n := len(inp[0])
	s := ""
	for i := 0; i < n; i++ {
		s += "."
	}
	for i := 0; i < 12; i++ {
		inp = append([]string{s}, inp...)
		inp = append(inp, s)
	}
	return key, inp
}
func loadDataV2() (map[int]bool, map[int]map[int]bool) {
	inp := c.ReadInputFile()
	key := map[int]bool{}
	res := map[int]map[int]bool{}
	for k, v := range inp[0] {
		key[k] = (v == '#')
	}

	inp = inp[2:]
	for y := range inp {
		for x := range inp[y] {
			if _, ok := res[y]; !ok {
				res[y] = map[int]bool{}
			}
			res[y][x] = inp[y][x] == '#'
		}
	}
	return key, res
}
