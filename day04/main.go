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
	nums, boards := loadGame()

	for _, v := range nums {
		for i := range boards {
			boards[i].markNumber(v)
			if boards[i].checkForBingo() {
				return v * boards[i].sum()
			}
		}
	}
	return -1
}
func part2() int {
	nums, boards := loadGame()

	for _, v := range nums {
		for i := len(boards) - 1; i >= 0; i-- {
			boards[i].markNumber(v)
			if boards[i].checkForBingo() {
				if len(boards) == 1 {
					return v * boards[i].sum()
				}
				boards = append(boards[:i], boards[i+1:]...)
			}
		}
	}
	return -1
}

func loadGame() ([]int, []board) {
	inp := c.ReadInputFile()
	nums := c.GetInts(strings.Split(inp[0], ","))
	boards := []board{}
	for i := range inp {
		if i == 0 {
			continue
		}
		if inp[i] == "" {
			boards = append(boards, board{})
		} else {
			for _, v := range strings.Split(inp[i], " ") {
				if v == "" {
					continue
				}
				boards[len(boards)-1].addNumber(c.GetInt(v))
			}
		}
	}
	return nums, boards
}

type board struct {
	numbers []int
}

func (b *board) addNumber(n int) {
	b.numbers = append(b.numbers, n)
}
func (b *board) markNumber(n int) {
	for i := range b.numbers {
		if b.numbers[i] == n {
			b.numbers[i] = -1
		}
	}
}
func (b *board) checkForBingo() bool {
	for i := 0; i < len(b.numbers); i += 5 {
		bingo := true
		for o := 0; o < 5; o++ {
			if b.numbers[i+o] != -1 {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}
	for o := 0; o < 5; o++ {
		bingo := true
		for i := 0; i < len(b.numbers); i += 5 {
			if b.numbers[i+o] != -1 {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}

	return false
}
func (b board) sum() int {
	sum := 0
	for _, v := range b.numbers {
		if v != -1 {
			sum += v
		}
	}
	return sum
}
