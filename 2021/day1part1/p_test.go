package p_test

import (
	"aoc/ax"
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
		{"small", 1},
		{"input", 1},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

var pat = regexp.MustCompile(``)

func run(lines []string) int {
	for _, line := range lines {
		parts := pat.FindStringSubmatch(line)
		_ = parts
	}
	return 0
}
