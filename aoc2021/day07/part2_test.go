package day07

import (
	"aoc/ax"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay07Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay07Part2(t *testing.T) {
	assert.Equal(t, "168", Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, "101571302", Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) string {
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
	var costLeft [2001]int
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
