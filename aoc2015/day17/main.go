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
	m := len(rs)
	mem := make(map[[2]int]int)
	var dfs func(i, rem int) int
	dfs = func(i, rem int) int {
		if rem == 0 {
			return 1
		}
		if rem < 0 {
			return 0
		}
		if i == m {
			return 0
		}
		k := [2]int{i, rem}
		if v, exists := mem[k]; exists {
			return v
		}
		res := dfs(i+1, rem) + dfs(i+1, rem-rs[i].x)
		mem[k] = res
		return res
	}
	res := dfs(0, 150)
	return sprint(res)
}

func Solve2(rs []parsedRow) string {
	m := len(rs)
	minContainers := math.MaxInt32
	minContainerCount := 0
	var dfs func(i, rem, n int) int
	dfs = func(i, rem, n int) int {
		if rem == 0 {
			if n < minContainers {
				minContainers = n
				minContainerCount = 1
			} else if n == minContainers {
				minContainerCount++
			}
			return 1
		}
		if rem < 0 {
			return 0
		}
		if i == m {
			return 0
		}
		res := dfs(i+1, rem, n) + dfs(i+1, rem-rs[i].x, n+1)
		return res
	}
	dfs(0, 150, 0)
	fmt.Println(minContainerCount)
	return sprint(minContainerCount)
}

type parsedRow struct {
	x int
}

func Parse(s string) parsedRow {
	var r parsedRow
	r.x = toi(s)
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
