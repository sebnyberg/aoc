package p_test

import (
	"aoc/ux"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", -1},
		{"input", -1},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ux.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

var pat = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)`)

func run(rows []string) int {
	return 0
}
