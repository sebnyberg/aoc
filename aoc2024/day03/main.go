package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var mulpat = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var dopat = regexp.MustCompile(`do\(\)`)
var dontpat = regexp.MustCompile(`don't\(\)`)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var res int
	program := strings.Join(lines, "")
	for _, s := range mulpat.FindAllStringSubmatch(program, -1) {
		a := ax.MustParseInt[int](s[1])
		b := ax.MustParseInt[int](s[2])
		res += a * b
	}
	return res
}

const (
	typMul  = 0
	typDo   = 1
	typDont = 2
)

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var res int
	program := strings.Join(lines, "")
	type event struct {
		typ  int
		idx  int
		args interface{}
	}
	var events []event
	doidx := dopat.FindAllStringIndex(program, -1)
	for _, idxs := range doidx {
		events = append(events, event{typ: typDo, idx: idxs[0], args: nil})
	}
	dontidx := dontpat.FindAllStringIndex(program, -1)
	for _, idxs := range dontidx {
		events = append(events, event{typ: typDont, idx: idxs[0], args: nil})
	}
	mulidx := mulpat.FindAllStringSubmatchIndex(program, -1)
	for _, idxs := range mulidx {
		a := ax.MustParseInt[int](program[idxs[2]:idxs[3]])
		b := ax.MustParseInt[int](program[idxs[4]:idxs[5]])
		events = append(events, event{typ: typMul, idx: idxs[0], args: []int{a, b}})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].idx < events[j].idx
	})
	var off bool
	for _, e := range events {
		switch e.typ {
		case typDo:
			off = false
		case typDont:
			off = true
		case typMul:
			if !off {
				intArgs := e.args.([]int)
				res += intArgs[0] * intArgs[1]
			}
		}
	}

	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
