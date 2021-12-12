package day6part2

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
		{"small", "26984457539"},
		{"input", "1767323539209"},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, Run(lines))
		})
	}
}

func BenchmarkPart(b *testing.B) {
	input := ax.MustReadFineLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Run(input)
	}
}
