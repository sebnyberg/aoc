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
	apply := func(s string) string {
		res := []byte{}
		val := s[0]
		var count int
		n := len(s)
		for i := 0; i <= n; i++ {
			if i == n || (i > 0 && s[i] != s[i-1]) {
				res = append(res, sprint(count)...)
				res = append(res, val)
				if i == n {
					break
				}
				count = 0
			}
			val = s[i]
			count++
		}
		return string(res)
	}
	s := rs[0].s
	for i := 0; i < 40; i++ {
		s = apply(s)
	}
	return sprint(len(s))
}

func Solve2(rs []parsedRow) string {
	apply := func(s string) string {
		res := []byte{}
		val := s[0]
		var count int
		n := len(s)
		for i := 0; i <= n; i++ {
			if i == n || (i > 0 && s[i] != s[i-1]) {
				res = append(res, sprint(count)...)
				res = append(res, val)
				if i == n {
					break
				}
				count = 0
			}
			val = s[i]
			count++
		}
		return string(res)
	}
	s := rs[0].s
	for i := 0; i < 50; i++ {
		s = apply(s)
	}
	return sprint(len(s))
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
