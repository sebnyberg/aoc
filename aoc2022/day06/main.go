package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func findNDiff(s string, n int) int {
	var ndistinct int
	var count [256]int
	for i := range s {
		if count[s[i]] == 0 {
			ndistinct++
		}
		count[s[i]]++
		if i >= n {
			if count[s[i-n]] == 1 {
				ndistinct--
			}
			count[s[i-n]]--
		}
		if ndistinct == n {
			return i + 1
		}
	}
	return -1
}

func solve1(inf string) string {
	line := ax.MustReadFileLines(inf)[0]
	return fmt.Sprint(findNDiff(line, 4))
}

func solve2(inf string) string {
	line := ax.MustReadFileLines(inf)[0]
	return fmt.Sprint(findNDiff(line, 13))
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
