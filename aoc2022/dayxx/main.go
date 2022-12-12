package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	for _, l := range lines {
		fs := strings.Fields(l)
		_ = fs
	}
	return ""
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
}
