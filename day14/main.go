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
	inp := c.ReadInputFile()

	sum := 0
	str := inp[0]
	rules := map[string]string{}
	for i := 2; i < len(inp); i++ {
		s := strings.Split(inp[i], " -> ")
		rules[s[0]] = s[1]
	}
	for i := 0; i < 10; i++ {
		wstr := ""
		for l := 0; l < len(str)-1; l++ {
			wstr += string(str[l])
			sub := str[l : l+2]
			for k, v := range rules {
				if k == sub {
					wstr += v
				}
			}
		}
		str = wstr + string(str[len(str)-1])
	}
	res := map[rune]int{}
	for _, v := range str {
		res[v]++
	}
	min := 0
	max := 0
	for _, v := range res {
		if v > max {
			max = v
		}
		if v < min || min == 0 {
			min = v
		}
	}
	sum = max - min
	return sum
}

type rule struct {
	next  string
	next2 string
	count int
}

func part2() int {
	inp := c.ReadInputFile()

	sum := 0
	str := inp[0]
	rules := map[string]*rule{}
	for i := 2; i < len(inp); i++ {
		s := strings.Split(inp[i], " -> ")
		rules[s[0]] = &rule{next: string(s[0][0]) + s[1], next2: s[1] + string(s[0][1])}
	}
	rules[str[len(str)-1:]] = &rule{count: 1}
	for l := 0; l < len(str)-1; l++ {
		sub := str[l : l+2]
		rules[sub].count++
	}
	for i := 0; i < 40; i++ {
		wrules := map[string]*rule{}
		for k, v := range rules {
			wrules[k] = &rule{next: v.next, count: v.count, next2: v.next2}
		}
		for k, v := range wrules {
			if v.count > 0 && len(k) == 2 {
				rules[k].count -= v.count
				rules[v.next].count += v.count
				rules[v.next2].count += v.count
			}
		}
		/*		s := map[string]int{}
				for k, v := range rules {
					if v.count > 0 {
						s[string(k[0])] += v.count
					}
				}
				for k, v := range s {
					c.Print(k, v)
				}
				c.Print("------")
		*/
	}
	res := map[string]int{}
	for k, v := range rules {
		res[string(k[0])] += v.count
	}
	min := 0
	max := 0
	for _, v := range res {
		if v > max {
			max = v
		}
		if v < min || min == 0 {
			min = v
		}
	}
	sum = max - min
	return sum
}
