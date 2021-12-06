package p_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 5},
		{"input", 1262},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

func run(lines []string) int {
	lineInts := make([]int, len(lines))
	for i := range lines {
		lineInts[i] = ax.MustParseIntBase(lines[i], 10)
	}
	var count int
	var cur, prev int
	for i := 0; i < len(lineInts); i++ {
		cur += lineInts[i]
		if i > 2 {
			cur -= lineInts[i-3]
			if prev < cur {
				count++
			}
		}
		prev = cur
	}
	return count
}
