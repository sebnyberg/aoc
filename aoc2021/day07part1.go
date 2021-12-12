package aoc2021

import (
	"aoc/ax"
	"math"
	"strconv"
	"strings"
)

func Day07Part1(rows []string) string {
	const size = 2000
	row := rows[0]

	// Capture frequency of crabs per position
	var posCount [size]int
	for _, valStr := range strings.Split(row, ",") {
		posCount[ax.MustParseInt[int](valStr)]++
	}

	// Calculate cost of moving all crabs from the left to a given position
	costLeft := make([]int, size)
	count := posCount[0]
	for i := 1; i < size; i++ {
		costLeft[i] = costLeft[i-1] + count
		count += posCount[i]
	}

	// Calculate cost of moving all crabs on the right side
	// Minimum cost is the combination of costLeft + costRight
	minCost := math.MaxInt32
	count = posCount[size-1]
	costRight := 0
	for i := size - 2; i >= 0; i-- {
		costRight += count
		count += posCount[i]
		minCost = ax.Min(minCost, costLeft[i]+costRight)
	}

	return strconv.Itoa(minCost)
}
