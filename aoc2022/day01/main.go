package main

import (
	"bufio"
	"os"
	"sort"

	"github.com/sebnyberg/aoc/ax"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	cals := []int{0}
	for sc.Scan() {
		s := sc.Text()
		if s == "" {
			cals = append(cals, 0)
			continue
		}
		cals[len(cals)-1] += ax.MustParseInt[int](s)
	}
	sort.Ints(cals)
	println(ax.Sum(cals[len(cals)-3:]))
}
