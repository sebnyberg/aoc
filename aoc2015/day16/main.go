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

const wants = `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`

func Solve1(rs []parsedRow) string {
	want := map[string]int{}
	for _, row := range strings.Split(wants, "\n") {
		a := strings.Split(row, ": ")
		want[a[0]] = toi(a[1])
	}
	res := -1
	for _, aunt := range rs {
		ok := true
		for prop, c := range aunt.props {
			if want[prop] != c {
				ok = false
				break
			}
		}
		if ok {
			res = aunt.idx
			break
		}
	}
	return sprint(res)
}

func Solve2(rs []parsedRow) string {
	want := map[string]int{}
	for _, row := range strings.Split(wants, "\n") {
		a := strings.Split(row, ": ")
		want[a[0]] = toi(a[1])
	}
	res := -1
	for _, aunt := range rs {
		ok := true
		for prop, c := range aunt.props {
			if prop == "trees" || prop == "cats" {
				if want[prop] >= c {
					ok = false
					break
				}
				continue
			}
			if prop == "pomeranians" || prop == "goldfish" {
				if want[prop] <= c {
					ok = false
					break
				}
				continue
			}
			if want[prop] != c {
				ok = false
				break
			}
		}
		if ok {
			res = aunt.idx
			break
		}
	}
	return sprint(res)
}

type parsedRow struct {
	s     string
	idx   int
	props map[string]int
}

var rrr = regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

func Parse(s string) parsedRow {
	var r parsedRow
	r.props = make(map[string]int)
	ss := rrr.FindStringSubmatch(s)
	r.idx = toi(ss[1])
	for i := 2; i < len(ss)-1; i += 2 {
		r.props[ss[i]] = toi(ss[i+1])
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
