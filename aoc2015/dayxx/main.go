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

func Solve1(rs []parsedRow) string {
	return ""
}

func Solve2(rs []parsedRow) string {
	return ""
}

type parsedRow struct {
	s string
}

func Parse(s string) parsedRow {
	var r parsedRow
	r.s = s
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var p ax.Problem[parsedRow]
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
