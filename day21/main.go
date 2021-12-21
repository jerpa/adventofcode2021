package main

import (
	"strings"
	"time"

	c "github.com/jerpa/adventofcode2021/common"
)

type player struct {
	pos   int
	point int
}

func (p *player) move(num int) {

	p.pos += num
	for p.pos > 10 {
		p.pos -= 10
	}
}
func (p *player) addPoint() {
	p.point += p.pos
}

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

	dice := 0
	diceCount := 0
	for {
		dice = (dice % 1000) + 1
		data[0].move(dice)
		dice = (dice % 1000) + 1
		data[0].move(dice)
		dice = (dice % 1000) + 1
		data[0].move(dice)
		data[0].addPoint()
		diceCount += 3
		if data[0].point >= 1000 {
			break
		}
		dice = (dice % 1000) + 1
		data[1].move(dice)
		dice = (dice % 1000) + 1
		data[1].move(dice)
		dice = (dice % 1000) + 1
		data[1].move(dice)
		data[1].addPoint()
		diceCount += 3
		if data[1].point >= 1000 {
			break
		}

	}
	return c.MinInt([]int{data[0].point, data[1].point}) * diceCount

}

func part2() int {

	return 0
}

func loadData() []player {
	res := []player{}

	for _, v := range c.ReadInputFile() {
		p := player{}
		p.pos = c.GetInt(strings.TrimSpace(string(v[len(v)-2:])))
		res = append(res, p)
	}
	return res
}
