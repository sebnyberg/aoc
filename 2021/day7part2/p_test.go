package p_test

import (
	"aoc/ax"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 168},
		{"input", 101571302},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(parseInput(lines)))
		})
	}
}

var benchRes int

func BenchmarkPart(b *testing.B) {
	input := ax.MustReadFineLines("input")
	vals := parseInput(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchRes = run(vals)
	}
}

func parseInput(rows []string) []int {
	valStrs := strings.Split(rows[0], ",")
	res := make([]int, len(valStrs))
	for i, valStr := range valStrs {
		res[i] = ax.MustParseInt[int](valStr)
	}
	return res
}

const SZ int = 2001

func run(vals []int) int {
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

	return int(minCost)
}
