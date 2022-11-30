package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint]
var tof = ax.MustParseFloat[float64]
var mini = ax.Min[int]
var minf = ax.Min[float64]
var minu = ax.Min[uint]

func Solve1(input []string) string {
	var floor int
	for _, ch := range input[0] {
		switch ch {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return sprint(floor)
}

func Solve2(input []string) string {
	var floor int
	for i, ch := range input[0] {
		switch ch {
		case '(':
			floor++
		case ')':
			if floor == 0 {
				return sprint(i + 1)
			}
			floor--
		}
	}
	return ""
}

func Parse(s string) string {
	return s
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var p ax.Problem[string]
	p.HeadN = 3
	p.TailN = 3
	for sc.Scan() {
		s := sc.Text()
		p.Input = append(p.Input, s)
		p.Parsed = append(p.Parsed, Parse(s))
	}
	p.Result1 = Solve1(p.Parsed)
	p.Result2 = Solve2(p.Parsed)
	fmt.Fprint(os.Stdout, p)
}
