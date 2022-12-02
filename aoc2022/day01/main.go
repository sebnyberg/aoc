package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/sebnyberg/aoc/ax"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var cals []int
	for _, s := range ax.MustReadFileLines("input") {
		if s == "" {
			cals = append(cals, 0)
			continue
		}
		cals[len(cals)-1] += ax.MustParseInt[int](s)
	}
	sort.Ints(cals)
	fmt.Println(ax.Sum(cals[len(cals)-3:]))
}
