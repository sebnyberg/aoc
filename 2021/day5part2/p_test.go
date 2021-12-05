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
		{"small", 12},
		{"input", 20484},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

var pat = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)`)

func run(rows []string) int {
	var grid [1001][1001]uint16
	var count int
	visit := func(x, y int) {
		grid[y][x]++
		if grid[y][x] == 2 {
			count++
		}
	}
	dir := func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	}

	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		x1, y1 := ax.MustParseInt(parts[1]), ax.MustParseInt(parts[2])
		x2, y2 := ax.MustParseInt(parts[3]), ax.MustParseInt(parts[4])
		dx, dy := ax.Abs(x2-x1), ax.Abs(y2-y1)
		dirX, dirY := dir(x1, x2), dir(y1, y2)
		if dx != dy && dx != 0 && dy != 0 {
			continue
		}
		delta := max(dx, dy)
		for i := 0; i <= delta; i++ {
			visit(x1, y1)
			x1 += dirX
			y1 += dirY
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
