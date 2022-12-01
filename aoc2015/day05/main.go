package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
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
	for _, x := range in.xs {
		res += nice(x.s)
	}
	return fmt.Sprint(res)
}

func solve2(in *input) string {
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
	for i := range in.xs {
		x := nice(in.xs[i].s)
		if x == 0 && k > 0 {
			println("not ok", in.xs[i].s)
			k--
		}
		if x == 1 && kk > 0 {
			println("ok", in.xs[i].s)
			kk--
		}
		if k > 0 {

		}
		res += x
	}
	return fmt.Sprint(res)
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
