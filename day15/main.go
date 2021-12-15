package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra"
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

	graph := dijkstra.NewGraph()

	for y, yv := range inp {
		for x := range strings.Split(yv, "") {
			graph.AddMappedVertex(fmt.Sprintf("%d,%d", x, y))
		}
	}
	for y, yv := range inp {
		for x := range strings.Split(yv, "") {
			cx, cy, err := getCoord(len(yv), len(inp), x, y, 0, -1)
			if err == nil {
				val := c.GetInt(string(inp[cy][cx]))
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", cx, cy), int64(val))
			}
			cx, cy, err = getCoord(len(yv), len(inp), x, y, -1, 0)
			if err == nil {
				val := c.GetInt(string(inp[cy][cx]))
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", cx, cy), int64(val))
			}
			cx, cy, err = getCoord(len(yv), len(inp), x, y, 1, 0)
			if err == nil {
				val := c.GetInt(string(inp[cy][cx]))
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", cx, cy), int64(val))
			}
			cx, cy, err = getCoord(len(yv), len(inp), x, y, 0, 1)
			if err == nil {
				val := c.GetInt(string(inp[cy][cx]))
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", cx, cy), int64(val))
			}
		}
	}
	start, err := graph.GetMapping("0,0")
	if err != nil {
		panic(err.Error())
	}
	end, err := graph.GetMapping(fmt.Sprintf("%d,%d", len(inp[0])-1, len(inp)-1))
	if err != nil {
		panic(err.Error())
	}
	res, err := graph.Shortest(start, end)
	if err != nil {
		panic(err.Error())
	}

	return int(res.Distance)
}
func getCoord(w, h, x, y, dx, dy int) (int, int, error) {
	if (x == 0 && dx < 0) || (x == w-1 && dx > 0) || (y == 0 && dy < 0) || (y == h-1 && dy > 0) {
		return 0, 0, errors.New("outOfBounds")
	}
	return x + dx, y + dy, nil

}
func part2() int {
	inp := c.ReadInputFile()

	m := map[int]map[int]int{}

	graph := dijkstra.NewGraph()

	for y, yv := range inp {
		for x, xv := range strings.Split(yv, "") {
			if _, ok := m[y]; !ok {
				m[y] = map[int]int{}
			}
			m[y][x] = c.GetInt(xv)
			graph.AddMappedVertex(fmt.Sprintf("%d,%d", x, y))
			m[y][x+len(yv)] = (m[y][x] % 9) + 1
			graph.AddMappedVertex(fmt.Sprintf("%d,%d", x+len(yv), y))
			m[y][x+len(yv)*2] = (m[y][x+len(yv)] % 9) + 1
			graph.AddMappedVertex(fmt.Sprintf("%d,%d", x+len(yv)*2, y))
			m[y][x+len(yv)*3] = (m[y][x+len(yv)*2] % 9) + 1
			graph.AddMappedVertex(fmt.Sprintf("%d,%d", x+len(yv)*3, y))
			m[y][x+len(yv)*4] = (m[y][x+len(yv)*3] % 9) + 1
			graph.AddMappedVertex(fmt.Sprintf("%d,%d", x+len(yv)*4, y))
		}
	}
	for i := 1; i < 5; i++ {
		for y := range inp {
			for x, xv := range m[(i-1)*len(inp)+y] {
				if _, ok := m[(i)*len(inp)+y]; !ok {
					m[(i)*len(inp)+y] = map[int]int{}
				}
				m[(i)*len(inp)+y][x] = (xv % 9) + 1
			}
		}
	}

	for y, yv := range m {
		for x := range yv {
			if _, ok := m[y-1]; ok {
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", x, y-1), int64(m[y-1][x]))
			}
			if _, ok := m[y+1]; ok {
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", x, y+1), int64(m[y+1][x]))
			}
			if _, ok := m[y][x-1]; ok {
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", x-1, y), int64(m[y][x-1]))
			}
			if _, ok := m[y][x+1]; ok {
				graph.AddMappedArc(fmt.Sprintf("%d,%d", x, y), fmt.Sprintf("%d,%d", x+1, y), int64(m[y][x+1]))
			}
		}
	}
	start, err := graph.GetMapping("0,0")
	if err != nil {
		panic(err.Error())
	}
	end, err := graph.GetMapping(fmt.Sprintf("%d,%d", len(m[0])-1, len(m)-1))
	if err != nil {
		panic(err.Error())
	}
	res, err := graph.Shortest(start, end)
	if err != nil {
		panic(err.Error())
	}

	return int(res.Distance)
}
