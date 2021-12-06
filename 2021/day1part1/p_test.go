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
		{"small", 7},
		{"input", 1292},
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
	for i := 1; i < len(lineInts); i++ {
		if lineInts[i-1] < lineInts[i] {
			count++
		}
	}
	return count
}
