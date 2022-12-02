package main

import (
	"fmt"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) string {
	var res int
	for _, s := range ax.MustReadFileLines(inf) {
		// var a, b, c int
		// fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
		// ss := pat.FindAllString(s, -1)
		// a := ax.Atoi(ss[0])
		_ = s
	}
	return fmt.Sprint(res)
}

func solve2(inf string) string {
	var res int
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
