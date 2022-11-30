package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var absf = ax.Abs[float64]
var absi = ax.Abs[int]
var minf = ax.Min[float64]
var mini = ax.Min[int]
var minu = ax.Min[uint16]
var maxf = ax.Max[float64]
var maxi = ax.Max[int]
var maxu = ax.Max[uint16]
var print = fmt.Print
var printf = fmt.Printf
var println = fmt.Println
var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var sprintln = fmt.Sprintln
var tof = ax.MustParseFloat[float64]
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint16]

func pprint(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmt.Printf(fmtStr, a...)
}
func pprintln(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmtStr += "\n"
	fmt.Printf(fmtStr, a...)
}

var intr = regexp.MustCompile(`[1-9][0-9]*|0`)

func isnum(s string) bool {
	return intr.MatchString(s)
}

func Solve1(rs []parsedRow) string {
	var res int
	var dfs func([5]int, int, int)
	dfs = func(stats [5]int, rem, i int) {
		if i == len(rs) {
			if rem != 0 {
				return
			}
			a := maxi(0, stats[0]) * maxi(0, stats[1]) *
				maxi(0, stats[2]) * maxi(0, stats[3])
			res = maxi(res, a)
			return
		}
		for x := 0; x <= rem; x++ {
			c := x * rs[i].capacity
			d := x * rs[i].durability
			f := x * rs[i].flavour
			t := x * rs[i].texture
			stats[0] += c
			stats[1] += d
			stats[2] += f
			stats[3] += t
			dfs(stats, rem-x, i+1)
			stats[0] -= c
			stats[1] -= d
			stats[2] -= f
			stats[3] -= t
		}
	}
	dfs([5]int{}, 100, 0)
	return sprint(res)
}

func Solve2(rs []parsedRow) string {
	var res int
	var dfs func([5]int, int, int)
	dfs = func(stats [5]int, rem, i int) {
		if i == len(rs) {
			if rem != 0 || stats[4] != 500 {
				return
			}
			a := maxi(0, stats[0]) * maxi(0, stats[1]) *
				maxi(0, stats[2]) * maxi(0, stats[3])
			res = maxi(res, a)
			return
		}
		for x := 0; x <= rem; x++ {
			stats[0] += x * rs[i].capacity
			stats[1] += x * rs[i].durability
			stats[2] += x * rs[i].flavour
			stats[3] += x * rs[i].texture
			stats[4] += x * rs[i].calories
			dfs(stats, rem-x, i+1)
			stats[0] -= x * rs[i].capacity
			stats[1] -= x * rs[i].durability
			stats[2] -= x * rs[i].flavour
			stats[3] -= x * rs[i].texture
			stats[4] -= x * rs[i].calories
		}
	}
	dfs([5]int{}, 100, 0)
	return sprint(res)
}

type parsedRow struct {
	s          string
	name       string
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

var rrr = regexp.MustCompile(`(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

func Parse(s string) parsedRow {
	var r parsedRow
	// r.s = s
	ss := rrr.FindStringSubmatch(s)
	r.name = ss[1]
	r.capacity = toi(ss[2])
	r.durability = toi(ss[3])
	r.flavour = toi(ss[4])
	r.texture = toi(ss[5])
	r.calories = toi(ss[6])
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
