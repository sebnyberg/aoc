package day11part1

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
		{"small", 1656},
		{"input", 1562},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, Run(lines, 100))
		})
	}
}
