package day17

import (
	"regexp"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day17part1 int

func BenchmarkDay17Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day17part1 = Part1(ax.MustReadFileLines("input")[0])
	}
}

func TestDay17Part1(t *testing.T) {
	assert.Equal(t, 45, Part1("target area: x=20..30, y=-10..-5"))
	assert.Equal(t, 11175, Part1(ax.MustReadFileLines("input")[0]))
}

func Part1(row string) int {
	var pat = regexp.MustCompile(`^target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)$`)
	parts := pat.FindStringSubmatch(row)
	x0 := ax.MustParseInt[int](parts[1])
	x1 := ax.MustParseInt[int](parts[2])
	y0 := ax.MustParseInt[int](parts[3])
	y1 := ax.MustParseInt[int](parts[4])
	var maxY int
	for x := 0; x < 500; x++ {
		for y := 0; y < 500; y++ {
			curX, curY := 0, 0
			xVel, yVel := x, y
			var curMaxY int
			for curY >= y0 && curX <= x1 {
				curX += xVel
				curY += yVel
				curMaxY = ax.Max(curMaxY, curY)
				if curX >= x0 && curX <= x1 && curY >= y0 && curY <= y1 {
					maxY = ax.Max(maxY, curMaxY)
					break
				}
				if xVel > 0 {
					xVel--
				}
				yVel--
			}
		}
	}

	return maxY
}
