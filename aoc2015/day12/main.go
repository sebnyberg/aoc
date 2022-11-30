package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

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
	s := rs[0].s
	s += "  " // sentinel
	sign := 1
	var val int
	var sum int
	findMatchingParen := func(i int) int {
		var parens int
		i++
		for i < len(s) {
			if s[i] == '{' {
				parens++
			} else if s[i] == '}' {
				parens--
				if parens == 0 {
					break
				}
			}
			i++
		}

		return mini(len(s), i+1)
	}
	for i := range s {
		if s[i] == '{' {
			j := findMatchingParen(i)
			if strings.Contains(s[i:j], ":\"red\"") {
				i = j - 1
				continue
			}
		}
		if s[i] == '-' {
			sign = -1
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			val = val*10 + int(s[i]-'0')
			continue
		}
		sum += val * sign
		val = 0
		sign = 1
	}
	return sprint(sum)
}

func Solve2(rs []parsedRow) string {
	s := rs[0].s
	s += "  " // sentinel
	sign := 1
	var val int
	stack := []int{0}
	hasRed := []bool{false}
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			sign = -1
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			val = val*10 + int(s[i]-'0')
			continue
		}
		stack[len(stack)-1] += val * sign
		val = 0
		sign = 1
		if s[i] == '{' {
			stack = append(stack, 0)
			hasRed = append(hasRed, false)
			continue
		}
		if i+6 <= len(s) && s[i:i+6] == `:"red"` {
			hasRed[len(hasRed)-1] = true
		}
		if s[i] == '}' {
			if !hasRed[len(hasRed)-1] {
				stack[len(stack)-2] += stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			hasRed = hasRed[:len(hasRed)-1]
		}
	}
	return sprint(stack[0])
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
