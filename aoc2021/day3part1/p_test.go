package day3part1

import (
	"aoc/ax"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  string
	}{
		{"small", "198"},
		{"input", "775304"},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, Run(lines))
		})
	}
}

func BenchmarkRun(b *testing.B) {
	lines := ax.MustReadFineLines("input")
	for n := 0; n < b.N; n++ {
		Run(lines)
	}
}
