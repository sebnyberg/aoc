package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

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

func solve1(in *input) string {
	rs := in.xs
	want := map[string]int{}
	for _, row := range strings.Split(wants, "\n") {
		a := strings.Split(row, ": ")
		want[a[0]] = ax.Atoi(a[1])
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
	return fmt.Sprint(res)
}

func solve2(rs *input) string {
	want := map[string]int{}
	for _, row := range strings.Split(wants, "\n") {
		a := strings.Split(row, ": ")
		want[a[0]] = ax.Atoi(a[1])
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
	return fmt.Sprint(res)
}

type parsedRow struct {
	s     string
	idx   int
	props map[string]int
}

type inputItem struct {
	s     string
	idx   int
	props map[string]int
}

type input struct {
	n  int
	xs []inputItem
}

var rrr = regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

func (p *input) parse(s string) {
	var x inputItem

	x.props = make(map[string]int)
	ss := rrr.FindStringSubmatch(s)
	x.idx = ax.Atoi(ss[1])
	for i := 2; i < len(ss)-1; i += 2 {
		x.props[ss[i]] = ax.Atoi(ss[i+1])
	}

	x.s = s
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	rows := ax.ReadLines(os.Stdin)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
