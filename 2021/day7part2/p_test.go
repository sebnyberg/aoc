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
			require.Equal(t, tc.want, run(lines))
		})
	}
}

const SZ = 2000

func run(rows []string) int {
	row := rows[0]
	var posCount [SZ]int
	for _, valStr := range strings.Split(row, ",") {
		posCount[ax.MustParseInt[int](valStr)]++
	}

	// Calculate cost of moving all crabs from the left to a given position
	costLeft := make([]int, SZ)
	for i := 1; i < SZ; i++ {
		var count int
		for j := 0; j <= i; j++ {
			costLeft[i] += count * (i - j + 1)
			count += posCount[j]
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

	return minCost
}
