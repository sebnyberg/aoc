package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	rs := in.xs
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

func solve2(in *input) string {
	in.xs[0].s = solve1(in)
	return solve1(in)
}

type inputItem struct {
	s       string
	t       string
	a, b, c int
	x, y    int
	x1, y1  int
	x2, y2  int
}

type input struct {
	n  int
	xs []inputItem
}

var pat = regexp.MustCompile(``)

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
