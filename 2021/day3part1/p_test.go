package p_test

import (
	"aoc/ax"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 198},
		{"input", 775304},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

func BenchmarkRun(b *testing.B) {
	lines := ax.MustReadFineLines("input")
	for n := 0; n < b.N; n++ {
		run(lines)
	}
}

func run(lines []string) int {
	n := len(lines[0])
	m := len(lines)
	oneCount := make([]int, n)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				oneCount[i]++
			}
		}
	}
	var gamma, eps int
	for _, count := range oneCount {
		gamma <<= 1
		eps <<= 1
		if count*2 > m {
			gamma += 1
		} else {
			eps += 1
		}
	}
	return gamma * eps
}
