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
			pos += mini(ft, 2503-t) * s
		}
		if pos > res {
			res = pos
		}
	}
	return sprint(res)
}

func Solve2(rs []parsedRow) string {
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
	return sprint(maxPoints)
}

type parsedRow struct {
	s        string
	name     string
	speed    int
	flyTime  int
	restTime int
}

var rrr = regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

func Parse(s string) parsedRow {
	var r parsedRow
	ss := rrr.FindStringSubmatch(s)
	r.name = ss[1]
	r.speed = toi(ss[2])
	r.flyTime = toi(ss[3])
	r.restTime = toi(ss[4])
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
