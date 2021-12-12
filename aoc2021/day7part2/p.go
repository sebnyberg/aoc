package day7part2

import (
	"aoc/ax"
	"strconv"
	"strings"
)

const (
	Problem = 7
	Part    = 2

	SZ = 2001
)

func Run(rows []string) string {
	valStrs := strings.Split(rows[0], ",")
	vals := make([]int, len(valStrs))
	for i, valStr := range valStrs {
		vals[i] = ax.MustParseInt[int](valStr)
	}

	var posCount [SZ]int
	for _, val := range vals {
		posCount[val]++
	}

	// Calculate cost of moving all crabs from the left to a given position
	costLeft := make([]int, SZ)
	for i := int(0); i < SZ; i++ {
		var cost int
		for j := i + 1; j < SZ; j++ {
			cost += (j - i) * posCount[i]
			costLeft[j] += cost
		}
	}

	// Calculate cost of moving all crabs on the right side
	// Minimum cost is the combination of costLeft + costRight
	minCost := costLeft[SZ-1]
	for i := SZ - 2; i >= 0; i-- {
		var count, costRight int
		for j := SZ - 1; j >= i; j-- {
			costRight += count * (j - i + 1)
			count += posCount[j]
		}
		minCost = ax.Min(minCost, costLeft[i]+costRight)
	}

	return strconv.Itoa(minCost)
}