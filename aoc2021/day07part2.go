package aoc2021

import (
	"aoc/ax"
	"strconv"
	"strings"
)

func Day07Part2(rows []string) string {
	const sz = 2001
	valStrs := strings.Split(rows[0], ",")
	vals := make([]int, len(valStrs))
	for i, valStr := range valStrs {
		vals[i] = ax.MustParseInt[int](valStr)
	}

	var posCount [sz]int
	for _, val := range vals {
		posCount[val]++
	}

	// Calculate cost of moving all crabs from the left to a given position
	costLeft := make([]int, sz)
	for i := int(0); i < sz; i++ {
		var cost int
		for j := i + 1; j < sz; j++ {
			cost += (j - i) * posCount[i]
			costLeft[j] += cost
		}
	}

	// Calculate cost of moving all crabs on the right side
	// Minimum cost is the combination of costLeft + costRight
	minCost := costLeft[sz-1]
	for i := sz - 2; i >= 0; i-- {
		var count, costRight int
		for j := sz - 1; j >= i; j-- {
			costRight += count * (j - i + 1)
			count += posCount[j]
		}
		minCost = ax.Min(minCost, costLeft[i]+costRight)
	}

	return strconv.Itoa(minCost)
}
