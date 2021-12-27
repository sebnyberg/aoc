package day17

import (
	"aoc/ax"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day17part2res int

func BenchmarkDay17Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day17part2res = Part2(ax.MustReadFileLines("input")[0])
	}
}

func TestDay17Part2(t *testing.T) {
	assert.Equal(t, 112, Part2("target area: x=20..30, y=-10..-5"))
	assert.Equal(t, 3540, Part2(ax.MustReadFileLines("input")[0]))
}

func Part2(row string) int {
	var pat = regexp.MustCompile(`^target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)$`)
	parts := pat.FindStringSubmatch(row)
	x0 := ax.MustParseInt[int](parts[1])
	x1 := ax.MustParseInt[int](parts[2])
	y0 := ax.MustParseInt[int](parts[3])
	y1 := ax.MustParseInt[int](parts[4])
	var uniq int
	for x := 0; x < 500; x++ {
		for y := -500; y < 500; y++ {
			curX, curY := 0, 0
			dx, dy := x, y
			var curMaxY int
			for curY >= y0 && curX <= x1 {
				curX += dx
				curY += dy
				curMaxY = ax.Max(curMaxY, curY)
				if curX >= x0 && curX <= x1 && curY >= y0 && curY <= y1 {
					uniq++
					break
				}
				if dx > 0 {
					dx--
				}
				dy--
			}
		}
	}

	return uniq
}
