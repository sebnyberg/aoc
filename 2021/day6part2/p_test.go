package p_test

import (
	"aoc/ux"
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
		{"small", 26984457539},
		{"input", 1767323539209},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ux.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines, 256))
		})
	}
}

func BenchmarkPart(b *testing.B) {
	input := ux.MustReadFineLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		run(input, 256)
	}
}

func run(rows []string, ndays int) int {
	var fishCount [9]int
	for _, valStr := range strings.Split(rows[0], ",") {
		val := ux.MustParseInt(valStr)
		fishCount[val]++
	}
	var nextCount [9]int
	for day := 0; day <= ndays; day++ {
		for i := 0; i < 8; i++ {
			nextCount[i] = fishCount[(i+1)%9]
		}
		nextCount[6] += fishCount[0]
		nextCount[8] = fishCount[0]
		nextCount, fishCount = fishCount, nextCount
	}
	var sum int
	for _, count := range nextCount {
		sum += count
	}
	return sum
}
