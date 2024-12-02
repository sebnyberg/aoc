package day24

import (
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

type pairCond struct {
	// first and second index part of this pair
	first, second int
	// how much larger the first must be compared to the second index
	delta int
}

// parsePairConds parses conditions based on the input rows.
// Returns a list of conditions and pairwise mappings.
func parsePairConds(rows []string) []pairCond {
	stackIndices := ax.Stack[[2]int]{}
	pairs := make([]pairCond, 0, 7)
	for i := 0; i < 14; i++ {
		a := ax.MustParseInt[int](strings.Split(rows[i*18+5], " ")[2])
		b := ax.MustParseInt[int](strings.Split(rows[i*18+15], " ")[2])
		if rows[i*18+4] == "div z 1" {
			// Push
			stackIndices.Push([2]int{i, b})
		} else {
			// Pop
			item := stackIndices.Pop()
			pairs = append(pairs, pairCond{
				first:  item[0],
				second: i,
				delta:  item[1] + a,
			})
		}
	}
	return pairs
}
