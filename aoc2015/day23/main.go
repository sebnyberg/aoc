package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input, reg [2]int) string {
	var pc int
	n := len(in.xs)
	for pc >= 0 && pc < n {
		parts := strings.Fields(in.xs[pc].s)
		switch parts[0] {
		case "inc":
			reg[parts[1][0]-'a']++
			pc++
		case "tpl":
			reg[parts[1][0]-'a'] *= 3
			pc++
		case "hlf":
			reg[parts[1][0]-'a'] /= 2
			pc++
		case "jmp":
			offset := ax.Atoi(parts[1][1:])
			if parts[1][0] == '-' {
				offset = -offset
			}
			pc += offset
		case "jie":
			i := int(parts[1][0] - 'a')
			if reg[i]%2 == 0 {
				offset := ax.Atoi(parts[2][1:])
				if parts[2][0] == '-' {
					offset = -offset
				}
				pc += offset
			} else {
				pc++
			}
		case "jio":
			i := int(parts[1][0] - 'a')
			if reg[i] == 1 {
				offset := ax.Atoi(parts[2][1:])
				if parts[2][0] == '-' {
					offset = -offset
				}
				pc += offset
			} else {
				pc++
			}
		}
	}
	return fmt.Sprint(reg[1])
}

func solve2(in *input) string {
	return solve1(in, [2]int{1, 0})
}

type inputItem struct {
	s string
}

type input struct {
	n  int
	xs []inputItem
}

func (p *input) parse(s string) {
	var x inputItem
	x.s = s
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	f, _ := os.Open("input")
	rows := ax.ReadLines(f)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in, [2]int{0, 0}))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
