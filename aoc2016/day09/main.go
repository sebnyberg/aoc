package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	var res int
	l := ax.MustReadFileLines(inf)[0]
	n := len(l)
	for i := 0; i < n; {
		if l[i] != '(' {
			res++
			i++
			continue
		}
		j := i + 1
		for l[j] != ')' {
			j++
		}
		ps := strings.Split(l[i+1:j], "x")
		nchars := ax.Atoi(ps[0])
		repeats := ax.Atoi(ps[1])
		res += nchars * repeats
		i = j + 1 + nchars
	}
	return fmt.Sprint(res)
}

func solve2(inf string) string {
	var res int
	// First let's check whether "nchars" can contain a non-complete parenthesis
	countOpen := func(s string) int {
		var open int
		for i := range s {
			if s[i] == '(' {
				open++
			} else if s[i] == ')' {
				open--
			}
		}
		return open
	}
	l := ax.MustReadFileLines(inf)[0]
	n := len(l)
	for i := 0; i < n; {
		if l[i] != '(' {
			res++
			i++
			continue
		}
		j := i + 1
		for l[j] != ')' {
			j++
		}
		ps := strings.Split(l[i+1:j], "x")
		nchars := ax.Atoi(ps[0])
		repeats := ax.Atoi(ps[1])
		thingToRepeat := l[j+1 : j+1+nchars]
		if countOpen(thingToRepeat) > 0 {
			panic(thingToRepeat)
		}
		res += nchars * repeats
		i = j + 1 + nchars
	}
	// According to the test, there is no decompressed region which does not
	// contain parts of another decompression. This helps.
	//
	// What decompression does is to multiply the repeats of the encompassed
	// decompressions.

	// The second question is whether a decompression contains a decompression
	// which reaches outside its range. Let's check for the case on the
	// recursion.
	a := calcLen(l)
	return fmt.Sprint(a)
}

func calcLen(s string) int {
	n := len(s)
	var res int
	for i := 0; i < n; {
		if s[i] != '(' {
			i++
			res++
			continue
		}
		j := i
		for s[j] != ')' {
			j++
		}
		ps := strings.Split(s[i+1:j], "x")
		nchars := ax.Atoi(ps[0])
		repeats := ax.Atoi(ps[1])
		i = j + 1
		a := repeats * calcLen(s[i:i+nchars])
		res += a
		if i+nchars > n {
			panic("not good")
		}
		i += nchars
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
