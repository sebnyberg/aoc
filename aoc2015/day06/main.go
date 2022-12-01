package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	var state [1000][1000]uint8
	rs := in.xs
	toggle := func(x uint8) uint8 {
		return 1 - x
	}
	on := func(x uint8) uint8 {
		return 1
	}
	off := func(x uint8) uint8 {
		return 0
	}
	for _, r := range rs {
		var f func(x uint8) uint8
		switch r.action {
		case "toggle":
			f = toggle
		case "turn on":
			f = on
		case "turn off":
			f = off
		default:
			panic(r.action)
		}
		for x := r.x1; x <= r.x2; x++ {
			for y := r.y1; y <= r.y2; y++ {
				state[x][y] = f(state[x][y])
			}
		}
	}
	var res int
	for i := range state {
		for _, v := range state[i] {
			res += int(v)
		}
	}
	return fmt.Sprint(res)
}

func solve2(in *input) string {
	var state [1000][1000]int
	toggle := func(x int) int {
		return x + 2
	}
	on := func(x int) int {
		return x + 1
	}
	off := func(x int) int {
		return ax.Max(0, x-1)
	}
	for _, r := range in.xs {
		var f func(x int) int
		switch r.action {
		case "toggle":
			f = toggle
		case "turn on":
			f = on
		case "turn off":
			f = off
		default:
			panic(r.action)
		}
		for x := r.x1; x <= r.x2; x++ {
			for y := r.y1; y <= r.y2; y++ {
				state[x][y] = f(state[x][y])
			}
		}
	}
	var res int
	for i := range state {
		for _, v := range state[i] {
			res += int(v)
		}
	}
	return fmt.Sprint(res)
}

type inputItem struct {
	action string
	x1, y1 int
	x2, y2 int
}

type input struct {
	n  int
	xs []inputItem
}

var pat = regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

func (p *input) parse(s string) {
	var x inputItem
	ss := pat.FindStringSubmatch(s)
	x.action = ss[1]
	x.x1 = ax.Atoi(ss[2])
	x.y1 = ax.Atoi(ss[3])
	x.x2 = ax.Atoi(ss[4])
	x.y2 = ax.Atoi(ss[5])
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
