package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	s := in.xs[0].s
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

		return ax.Min(len(s), i+1)
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
	return fmt.Sprint(sum)
}

func solve2(in *input) string {
	s := in.xs[0].s
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
	return fmt.Sprint(stack[0])
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
	rows := ax.ReadLines(os.Stdin)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
