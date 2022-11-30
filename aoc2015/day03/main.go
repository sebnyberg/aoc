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
	m := ax.Set[[2]int]{}
	var x, y int
	m.Add([2]int{x, y})
	for _, s := range rs[0].s {
		switch s {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		m.Add([2]int{x, y})
	}
	return sprint(len(m))
}

func Solve2(rs []parsedRow) string {
	m := ax.Set[[2]int]{}
	var xy [2][2]int
	x := 0
	y := 1
	m.Add([2]int{0, 0})
	for i := 0; i < len(rs[0].s); i += 2 {
		for j := 0; j < 2; j++ {
			switch rs[0].s[i+j] {
			case '^':
				xy[j][y]++
			case 'v':
				xy[j][y]--
			case '>':
				xy[j][x]++
			case '<':
				xy[j][x]--
			}
			m.Add([2]int{xy[j][x], xy[j][y]})
		}
	}
	return sprint(len(m))
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
