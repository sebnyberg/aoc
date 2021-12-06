package day5part1

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
		{"small", 5},
		{"input", 5092},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ux.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

func BenchmarkRun(b *testing.B) {
	input := ux.MustReadFineLines("input")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		run(input)
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
		x1, y1 := ux.MustParseInt(parts[1]), ux.MustParseInt(parts[2])
		x2, y2 := ux.MustParseInt(parts[3]), ux.MustParseInt(parts[4])
		dx, dy := ux.Abs(x2-x1), ux.Abs(y2-y1)
		dirX, dirY := dir(x1, x2), dir(y1, y2)
		if dx != 0 && dy != 0 {
			continue
		}
		delta := ux.Max(dx, dy)
		for i := 0; i <= delta; i++ {
			visit(x1, y1)
			x1 += dirX
			y1 += dirY
		}
	}

	return count
}
