package main

import (
	"bufio"
	"fmt"
	"math"
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
	names := []string{}
	idx := map[string]int{}
	for _, r := range rs {
		if _, exists := idx[r.a]; !exists {
			idx[r.a] = len(names)
			names = append(names, r.a)
		}
		if _, exists := idx[r.b]; !exists {
			idx[r.b] = len(names)
			names = append(names, r.b)
		}
	}
	n := len(names)
	deltas := make([][]int, n)
	for i := range deltas {
		deltas[i] = make([]int, n)
	}
	for _, r := range rs {
		deltas[idx[r.a]][idx[r.b]] = r.delta
	}

	res := math.MinInt32
	var tried int
	cost := func(arr []int) {
		var x int
		for i, j := range arr {
			left := arr[(i-1+n)%n]
			right := arr[(i+1)%n]
			x += deltas[j][left] + deltas[j][right]
		}
		tried++
		if x > res {
			res = x
		}
	}

	var dfs func([]bool, []int, int)
	dfs = func(seen []bool, arr []int, j int) {
		if j == n {
			cost(arr)
			return
		}
		for i := range seen {
			if seen[i] {
				continue
			}
			seen[i] = true
			arr[j] = i
			dfs(seen, arr, j+1)
			seen[i] = false
		}
	}

	seen := make([]bool, n)
	arr := make([]int, n)
	dfs(seen, arr, 0)

	return sprint(res)
}

func Solve2(rs []parsedRow) string {
	rs = append(rs, parsedRow{
		a:     "Me",
		b:     "Carol",
		delta: 0,
	})
	return Solve1(rs)
}

type parsedRow struct {
	s     string
	a     string
	b     string
	delta int
}

var rrr = regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)`)

func Parse(s string) parsedRow {
	var r parsedRow
	ss := rrr.FindStringSubmatch(s)
	r.a = ss[1]
	r.b = ss[4]
	r.delta = toi(ss[3])
	if ss[2] == "lose" {
		r.delta *= -1
	}
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
