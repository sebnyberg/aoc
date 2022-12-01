package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	var res int
	rs := in.xs
	for _, r := range rs {
		s := r.speed
		ft := r.flyTime
		rt := r.restTime

		var t int
		var pos int
		for t+ft <= 2503 {
			t += ft
			pos += s * ft
			t += rt
		}
		if t <= 2503 {
			pos += ax.Min(ft, 2503-t) * s
		}
		if pos > res {
			res = pos
		}
	}
	return fmt.Sprint(res)
}

func solve2(in *input) string {
	n := len(rs)
	pos := make([]int, n)
	points := make([]int, n)
	for t := 0; t < 2503; t++ {
		winners := []int{}
		var maxPos int
		for i, r := range rs {
			s := r.speed
			ft := r.flyTime
			rt := r.restTime
			if t%(ft+rt) < ft {
				// make progress
				pos[i] += s
			}
			if pos[i] > maxPos {
				winners = winners[:0]
				winners = append(winners, i)
				maxPos = pos[i]
			} else if pos[i] == maxPos {
				winners = append(winners, i)
			}
		}
		for _, i := range winners {
			points[i]++
		}
	}
	var maxPoints int
	for i := range points {
		if points[i] > maxPoints {
			maxPoints = points[i]
		}
	}
	return fmt.Sprint(maxPoints)
}

type inputItem struct {
	s        string
	name     string
	speed    int
	flyTime  int
	restTime int
}

type input struct {
	n  int
	xs []inputItem
}

var rrr = regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

func (p *input) parse(s string) {
	var x inputItem
	ss := rrr.FindStringSubmatch(s)
	x.name = ss[1]
	x.speed = ax.Atoi(ss[2])
	x.flyTime = ax.Atoi(ss[3])
	x.restTime = ax.Atoi(ss[4])
	x.s = s
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	rows := ax.ReadLines(os.Stdin)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
