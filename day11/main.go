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

type octopus struct {
	value   int
	flashed bool
}

func part1() int {
	inp := c.ReadInputFile()
	board := []octopus{}
	sum := 0
	for _, v := range inp {
		for _, n := range strings.Split(v, "") {
			board = append(board, octopus{value: c.GetInt(n)})
		}
	}
	for i := 0; i < 100; i++ {
		for l := range board {
			if board[l].value > 9 {
				board[l].value = 0
				board[l].flashed = false

			}
			board[l].value++
		}
		flash := true
		for flash {
			flash = false
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					if board[y*10+x].value > 9 && !board[y*10+x].flashed {
						flash = true
						board[y*10+x].flashed = true
						sum++
						board[getCoord(x, y, -1, -1)].value++
						board[getCoord(x, y, 0, -1)].value++
						board[getCoord(x, y, 1, -1)].value++
						board[getCoord(x, y, -1, 0)].value++
						board[getCoord(x, y, 1, 0)].value++
						board[getCoord(x, y, -1, 1)].value++
						board[getCoord(x, y, 0, 1)].value++
						board[getCoord(x, y, 1, 1)].value++
					}
				}
			}
		}
	}

	return sum
}
func getCoord(x, y, dx, dy int) int {
	if (x == 0 && dx < 0) || (x == 9 && dx > 0) || (y == 0 && dy < 0) || (y == 9 && dy > 0) {
		return y*10 + x
	}
	return (y+dy)*10 + (x + dx)

}
func part2() int {
	inp := c.ReadInputFile()
	board := []octopus{}
	sum := 0
	for _, v := range inp {
		for _, n := range strings.Split(v, "") {
			board = append(board, octopus{value: c.GetInt(n)})
		}
	}
	for i := 0; i < 5000; i++ {
		for l := range board {
			if board[l].value > 9 {
				board[l].value = 0
				board[l].flashed = false

			}
			board[l].value++
		}
		flash := true
		rflash := 0
		for flash {
			flash = false
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					if board[y*10+x].value > 9 && !board[y*10+x].flashed {
						flash = true
						rflash++
						board[y*10+x].flashed = true
						sum++
						board[getCoord(x, y, -1, -1)].value++
						board[getCoord(x, y, 0, -1)].value++
						board[getCoord(x, y, 1, -1)].value++
						board[getCoord(x, y, -1, 0)].value++
						board[getCoord(x, y, 1, 0)].value++
						board[getCoord(x, y, -1, 1)].value++
						board[getCoord(x, y, 0, 1)].value++
						board[getCoord(x, y, 1, 1)].value++
					}
				}
			}
		}
		if rflash == 100 {
			return i + 1
		}
	}

	return -1
}
