package main

import (
	"bufio"
	"encoding/hex"
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
	var code int
	var n int
	for i := range rs {
		code += len(rs[i].s)
		n += len(rs[i].inmem)
	}
	return sprint(code - n)
}

func Solve2(rs []parsedRow) string {
	var code int
	var n int
	for i := range rs {
		code += len(rs[i].s)
		n += len(rs[i].encoded) + 2
	}
	return sprint(n - code)
}

type parsedRow struct {
	s       string
	inmem   string
	encoded string
}

func Parse(s string) parsedRow {
	var r parsedRow
	r.s = s
	s = s[1 : len(s)-1]
	var i int
	var res []byte
	for i < len(s) {
		if s[i] != '\\' {
			res = append(res, s[i])
			i++
			continue
		}
		if s[i+1] == 'x' {
			chs, err := hex.DecodeString(s[i+2 : i+4])
			ax.Check(err, "invalid hex "+s[i:i+4])
			ch := chs[0]
			res = append(res, ch)
			i += 4
			continue
		}
		res = append(res, s[i+1])
		i += 2
	}
	r.inmem = string(res)
	// Encode by adding \ in front of any " or \
	res = []byte{}
	i = 0
	for i := range r.s {
		if r.s[i] == '"' || r.s[i] == '\\' {
			res = append(res, '\\')
		}
		res = append(res, r.s[i])
		i++
	}
	r.encoded = string(res)
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
