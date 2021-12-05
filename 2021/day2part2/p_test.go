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
		{"small", 900},
		{"input", 1592426537},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ux.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

var pat = regexp.MustCompile(`^(\w+) (\d+)`)

func run(lines []string) int {
	var horz, depth, aim int
	for _, line := range lines {
		parts := pat.FindStringSubmatch(line)
		dir := parts[1]
		val := ux.MustParseInt(parts[2])
		switch dir {
		case "forward":
			horz += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return depth * horz
}
