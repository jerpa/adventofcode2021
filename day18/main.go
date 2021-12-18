package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	c "github.com/jerpa/adventofcode2021/common"
)

var reg *regexp.Regexp
var regr *regexp.Regexp
var regl *regexp.Regexp
var splreg *regexp.Regexp
var mreg *regexp.Regexp

func init() {
	reg = regexp.MustCompile(`\d+,\d+`)
	regr = regexp.MustCompile(`\d+`)
	regl = regexp.MustCompile(`\d+`)
	splreg = regexp.MustCompile(`\d{2,}`)
	mreg = regexp.MustCompile(`\[\d+,\d+\]`)
}

func main() {
	start := time.Now()

	c.Print("Part1: ", part1())
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2())
	c.Print("Took: ", time.Since(start).String())
}
func add(s1, s2 string) string {
	s := "[" + s1 + "," + s2 + "]"
	m := true
	for m {
		s, m = explode(s)
		if !m {
			s, m = split(s)
		}
	}
	return s
}
func explode(s string) (string, bool) {
	level := 0

	for i, v := range s {
		if v == '[' {
			level++
		} else if v == ']' {
			level--
		} else if level > 4 {
			start := i
			stop := i + 1
			for ; stop < len(s); stop++ {
				if s[stop] == ']' {
					break
				}
			}

			if reg.MatchString(s[start:stop]) {
				nums := c.GetInts(strings.Split(s[start:stop], ","))

				if regr.MatchString(s[stop:]) {
					r := s[stop:]
					rr := regr.FindAllIndex([]byte(r), 1)[0]
					num := c.GetInt(r[rr[0]:rr[1]])
					r = fmt.Sprintf("%s%d%s", r[:rr[0]], num+nums[1], r[rr[1]:])
					s = s[:stop] + r
				}
				s = s[:start-1] + "0" + s[stop+1:]

				if regl.MatchString(s[:start-1]) {
					r := s[:start-1]
					rli := regl.FindAllIndex([]byte(r), -1)
					rl := rli[len(rli)-1]
					num := c.GetInt(r[rl[0]:rl[1]])
					r = fmt.Sprintf("%s%d%s", r[:rl[0]], num+nums[0], r[rl[1]:])
					s = r + s[start-1:]
				}
				return s, true
			}
		}
	}
	return s, false
}
func split(s string) (string, bool) {

	m := splreg.FindString(s)
	if len(m) > 0 {
		v := c.GetInt(m)
		n1 := v / 2
		n2 := v - n1
		return strings.Replace(s, m, fmt.Sprintf("[%d,%d]", n1, n2), 1), true
	}
	return s, false
}
func magnitude(s string) int {

	for mreg.MatchString(s) {
		m := mreg.FindAllString(s, -1)
		for _, v := range m {
			nums := c.GetInts(strings.Split(v[1:len(v)-1], ","))
			n := (nums[0] * 3) + (2 * nums[1])
			s = strings.ReplaceAll(s, v, fmt.Sprintf("%d", n))
		}
	}
	return c.GetInt(s)
}
func part1() int {
	inp := c.ReadInputFile()
	s := inp[0]

	for i := 1; i < len(inp); i++ {
		s = add(s, inp[i])
	}
	c.Print(s)

	return magnitude(s)
}
func part2() int {
	inp := c.ReadInputFile()

	max := 0
	for x := range inp {
		for y := range inp {
			if x == y {
				continue
			}
			n := magnitude(add(inp[x], inp[y]))
			if n > max {
				max = n
			}
		}
	}

	return max
}
