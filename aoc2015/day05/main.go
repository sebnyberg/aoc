package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var print = fmt.Print
var printf = fmt.Printf
var println = fmt.Println
var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint]
var tof = ax.MustParseFloat[float64]
var mini = ax.Min[int]
var minf = ax.Min[float64]
var minu = ax.Min[uint]

func Solve1(rs []parsedRow) string {
	illegal := map[string]bool{
		"ab": true,
		"cd": true,
		"pq": true,
		"xy": true,
	}
	nice := func(s string) int {
		var vowelCount int
		var hasTwo bool
		for i := range s {
			if strings.ContainsRune("aeiou", rune(s[i])) {
				vowelCount++
			}
			if i > 0 && s[i] == s[i-1] {
				hasTwo = true
			}
			if i > 0 && illegal[s[i-1:i+1]] {
				return 0
			}
		}
		if hasTwo && vowelCount >= 3 {
			return 1
		}
		return 0
	}
	var res int
	for i := range rs {
		res += nice(rs[i].s)
	}
	return sprint(res)
}

func Solve2(rs []parsedRow) string {
	nice := func(s string) int {
		n := len(s)
		var ok bool
		for i := 0; i < n-1; i++ {
			for j := i + 2; j+1 < n; j++ {
				if s[i:i+2] == s[j:j+2] {
					ok = true
					break
				}
			}
			if ok {
				break
			}
		}
		if !ok {
			return 0
		}

		// Second condition
		ok = false
		for i := 2; i < n; i++ {
			if s[i] == s[i-2] && s[i] != s[i-1] {
				ok = true
				break
			}
		}

		if !ok {
			return 0
		}
		return 1
	}
	var res int
	k := 5
	kk := 5
	for i := range rs {
		x := nice(rs[i].s)
		if x == 0 && k > 0 {
			println("not ok", rs[i].s)
			k--
		}
		if x == 1 && kk > 0 {
			println("ok", rs[i].s)
			kk--
		}
		if k > 0 {

		}
		res += x
	}
	return sprint(res)
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
