package day05

import (
	"aoc/ax"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkPart2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "12", Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, "20484", Part2(ax.MustReadFineLines("input")))
}

func Part2(rows []string) string {
	pat := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)`)
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
		x1, y1 := ax.MustParseInt[int](parts[1]), ax.MustParseInt[int](parts[2])
		x2, y2 := ax.MustParseInt[int](parts[3]), ax.MustParseInt[int](parts[4])
		dx, dy := ax.Abs(x2-x1), ax.Abs(y2-y1)
		dirX, dirY := dir(x1, x2), dir(y1, y2)
		if dx != dy && dx != 0 && dy != 0 {
			continue
		}
		delta := ax.Max(dx, dy)
		for i := 0; i <= delta; i++ {
			visit(x1, y1)
			x1 += dirX
			y1 += dirY
		}
	}

	return strconv.Itoa(count)
}
