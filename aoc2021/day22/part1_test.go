package day22

import (
	"regexp"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day22part1 int

func BenchmarkDay22Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day22part1 = Part1(ax.MustReadFileLines("input"))
	}
}

func TestDay22Part1(t *testing.T) {
	assert.Equal(t, 590784, Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, 527915, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	pat := regexp.MustCompile(`^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)
	var state [101][101][101]bool
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		x1 := ax.MustParseInt[int](parts[2]) + 50
		x2 := ax.MustParseInt[int](parts[3]) + 50
		y1 := ax.MustParseInt[int](parts[4]) + 50
		y2 := ax.MustParseInt[int](parts[5]) + 50
		z1 := ax.MustParseInt[int](parts[6]) + 50
		z2 := ax.MustParseInt[int](parts[7]) + 50
		if x2 < 0 || x1 > 100 || y2 < 0 || y1 > 100 || z2 < 0 || z1 > 100 {
			continue
		}
		var val bool
		if parts[1] == "on" {
			val = true
		}
		for x := ax.Max(0, x1); x <= ax.Min(100, x2); x++ {
			for y := ax.Max(0, y1); y <= ax.Min(100, y2); y++ {
				for z := ax.Max(0, z1); z <= ax.Min(100, z2); z++ {
					state[x][y][z] = val
				}
			}
		}
	}
	var count int
	for i := range state {
		for j := range state[i] {
			for _, val := range state[i][j] {
				if val {
					count++
				}
			}
		}
	}
	return count
}
