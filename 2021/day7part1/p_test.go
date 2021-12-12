package p_test

import (
	"aoc/ax"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 37},
		{"input", 352997},
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

	return minCost
}
