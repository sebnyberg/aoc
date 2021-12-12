package day7part1

import (
	"aoc/ax"
	"math"
	"strconv"
	"strings"
)

const (
	Problem = 7
	Part    = 1

	SZ = 2000
)

func Run(rows []string) string {
	row := rows[0]

	// Capture frequency of crabs per position
	var posCount [SZ]int
	for _, valStr := range strings.Split(row, ",") {
		posCount[ax.MustParseInt[int](valStr)]++
	}

	// Calculate cost of moving all crabs from the left to a given position
	costLeft := make([]int, SZ)
	count := posCount[0]
	for i := 1; i < SZ; i++ {
		costLeft[i] = costLeft[i-1] + count
		count += posCount[i]
	}

	// Calculate cost of moving all crabs on the right side
	// Minimum cost is the combination of costLeft + costRight
	minCost := math.MaxInt32
	count = posCount[SZ-1]
	costRight := 0
	for i := SZ - 2; i >= 0; i-- {
		costRight += count
		count += posCount[i]
		minCost = ax.Min(minCost, costLeft[i]+costRight)
	}

	return strconv.Itoa(minCost)
}
