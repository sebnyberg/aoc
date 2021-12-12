package day4part1

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
		{"small", "4512"},
		{"input", "65325"},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, Run(lines))
		})
	}
}

func BenchmarkTest(b *testing.B) {
	lines := ax.MustReadFineLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Run(lines)
	}
}
