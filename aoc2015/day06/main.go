package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var minf = ax.Min[float64]
var mini = ax.Min[int]
var minu = ax.Min[uint]
var maxf = ax.Max[float64]
var maxi = ax.Max[int]
var maxu = ax.Max[uint]
var print = fmt.Print
var printf = fmt.Printf
var println = fmt.Println
var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var sprintln = fmt.Sprintln
var tof = ax.MustParseFloat[float64]
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint]

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

func Solve1(rs []parsedRow) string {
	var state [1000][1000]uint8
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
	return sprint(res)
}

func Solve2(rs []parsedRow) string {
	var state [1000][1000]int
	toggle := func(x int) int {
		return x + 2
	}
	on := func(x int) int {
		return x + 1
	}
	off := func(x int) int {
		return maxi(0, x-1)
	}
	for _, r := range rs {
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
	return sprint(res)
}

type parsedRow struct {
	s      string
	action string
	x1, y1 int
	x2, y2 int
}

var r = regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

func Parse(s string) parsedRow {
	var row parsedRow
	ss := r.FindStringSubmatch(s)
	row.action = ss[1]
	row.x1 = toi(ss[2])
	row.y1 = toi(ss[3])
	row.x2 = toi(ss[4])
	row.y2 = toi(ss[5])
	return row
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
