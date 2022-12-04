package main

import (
	"fmt"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) string {
	var res int
	count := func(s string) [256]int {
		var res [256]int
		for _, ch := range s {
			res[ch]++
		}
		return res
	}
	for _, s := range ax.MustReadFileLines(inf) {
		n := len(s)
		left := s[:n/2]
		right := s[n/2:]

		l := count(left)
		r := count(right)
		var item byte
		for ch := range l {
			if l[ch] > 0 && r[ch] > 0 {
				item = byte(ch)
				break
			}
		}
		if item >= 'a' && item <= 'z' {
			res += int(item - 'a' + 1)
		} else {
			res += int(item - 'A' + 27)
		}
	}
	return fmt.Sprint(res)
}

func solve2(inf string) string {
	var res int
	count := func(s string) [256]int {
		var res [256]int
		for _, ch := range s {
			res[ch]++
		}
		return res
	}
	lines := ax.MustReadFileLines(inf)
	for i := 0; i < len(lines); i += 3 {
		first := count(lines[i])
		second := count(lines[i+1])
		third := count(lines[i+2])
		for i := range second {
			first[i] = ax.Min(first[i], second[i])
			first[i] = ax.Min(first[i], third[i])
		}

		var item byte
		for ch := range first {
			if first[ch] > 0 {
				item = byte(ch)
				break
			}
		}
		if item >= 'a' && item <= 'z' {
			res += int(item - 'a' + 1)
		} else {
			res += int(item - 'A' + 27)
		}
	}
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
