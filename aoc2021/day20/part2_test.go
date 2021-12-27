package day20

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day20part2 int

func BenchmarkDay20Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day20part1 = Part2(ax.MustReadFileLines("input"))
	}
}

func TestDay20Part2(t *testing.T) {
	assert.Equal(t, 3351, Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, 18074, Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) int {
	// Parse pixel enhancement
	enhance := make([]bool, len(rows[0]))
	for i := range rows[0] {
		if rows[0][i] == '#' {
			enhance[i] = true
		}
	}

	lights := make(ax.Set[pos], len(rows[2]))
	for i := 2; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] == '#' {
				lights.Add(pos{i - 2, j})
			}
		}
	}

	calc := func(lights ax.Set[pos], p pos) int {
		var val int
		for i := p.i - 1; i <= p.i+1; i++ {
			for j := p.j - 1; j <= p.j+1; j++ {
				val <<= 1
				if lights.Has(pos{i, j}) {
					val++
				}
			}
		}
		return val
	}

	inv := enhance[0] == true

	for i := 0; i < 50; i++ {
		// Since we cannot light an infinite grid, we need to sometimes invert what
		// we track from lights to black pixels. Then in the next iteration, the
		// enhancement grid needs to be changed to caulculate whether a pixel should
		// be lit based on surrounding black pixels rather than light pixels.

		// Based on the current grid, add pixels that may become non-zero to a list
		// of pixels to visit.
		toVisit := make(ax.Set[pos], len(lights))
		for p := range lights {
			// Add all pixels around the current one to the list of pixels to consider
			for i := p.i - 1; i <= p.i+1; i++ {
				for j := p.j - 1; j <= p.j+1; j++ {
					toVisit.Add(pos{i, j})
				}
			}
		}
		// Create new set of lights
		newLights := make(ax.Set[pos], len(lights))
		for p := range toVisit {
			val := calc(lights, p)
			if !inv {
				if enhance[val] {
					newLights.Add(p)
				}
			} else {
				// Check whether the light should become the opposite of what it is
				// right now
				if i%2 == 0 {
					if !enhance[val] {
						newLights.Add(p)
					}
				} else {
					// inverse the calculation - it was counting black pixels instead of
					// light ones
					mask := (1 << 9) - 1
					invVal := ^val & mask
					if enhance[invVal] {
						newLights.Add(p)
					}
				}
			}
		}

		lights = newLights
		// print(lights)
	}
	return len(lights)
}
