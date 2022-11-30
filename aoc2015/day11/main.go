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

	rev := func(b []byte) []byte {
		cpy := make([]byte, len(b))
		copy(cpy, b)
		b = cpy
		for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
			b[l], b[r] = b[r], b[l]
		}
		return b
	}

	incr := func(b []byte) []byte {
		b = rev(b)
		var carry byte
		b[0] += 1
		for i := range b {
			a := b[i] + carry
			carry = a / 26
			b[i] = a % 26
		}
		if carry > 0 {
			b = append(b, carry)
		}
		return rev(b)
	}

	illegal := [256]bool{
		'i' - 'a': true,
		'l' - 'a': true,
		'o' - 'a': true,
	}
	isValid := func(b []byte) bool {
		n := len(b)
		for i := range b {
			if illegal[b[i]] {
				return false
			}
		}
		var incr bool
		for i := 0; i < n-2; i++ {
			if b[i] == b[i+1]-1 && b[i+1] == b[i+2]-1 {
				incr = true
				break
			}
		}
		if !incr {
			return false
		}
		var count int
		for i := 0; i < len(b)-1; i++ {
			if b[i] == b[i+1] {
				count++
				if count == 2 {
					return true
				}
				i++
			}
		}
		return false
	}
	unshift := func(b []byte) []byte {
		cpy := make([]byte, len(b))
		copy(cpy, b)
		b = cpy
		for i := range b {
			b[i] -= 'a'
		}
		return b
	}
	shift := func(b []byte) []byte {
		cpy := make([]byte, len(b))
		copy(cpy, b)
		b = cpy
		for i := range b {
			b[i] += 'a'
		}
		return b
	}

	b := []byte(rs[0].s)
	b = unshift(b)
	b = incr(b)
	for !isValid(b) {
		b = incr(b)
	}
	return string(shift(b))
}

func Solve2(rs []parsedRow) string {
	rs[0].s = Solve1(rs)
	return Solve1(rs)
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
