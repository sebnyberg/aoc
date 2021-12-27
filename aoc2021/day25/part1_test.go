package day25

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day23part1 int

func BenchmarkDay23Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day23part1 = Part1(ax.MustReadFileLines("input"))
	}
}

func TestDay23Part1(t *testing.T) {
	assert.Equal(t, 50, Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, 582, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	// Move east first, then south
	n, m := len(rows), len(rows[0])
	cur := make([][]byte, n)
	for i := range cur {
		cur[i] = []byte(rows[i])
	}
	nmoved := 1
	makeNext := func() [][]byte {
		next := make([][]byte, n)
		for i := range next {
			next[i] = make([]byte, m)
			for j := range next[i] {
				next[i][j] = '.'
			}
		}
		return next
	}

	var i int
	for ; nmoved > 0; i++ {
		nmoved = 0
		// East
		next := makeNext()
		for i := range cur {
			for j := range cur[i] {
				k := (j + 1) % m
				if cur[i][j] == '>' && cur[i][k] == '.' {
					next[i][k] = cur[i][j]
					if cur[i][j] != '.' {
						nmoved++
					}
				} else if next[i][j] == '.' {
					next[i][j] = cur[i][j]
				}
			}
		}
		cur = next
		// South
		next = makeNext()
		for i := range cur {
			for j := range cur[i] {
				k := (i + 1) % n
				if cur[i][j] == 'v' && cur[k][j] == '.' {
					next[k][j] = cur[i][j]
					if cur[i][j] != '.' {
						nmoved++
					}
				} else if next[i][j] == '.' {
					next[i][j] = cur[i][j]
				}
			}
		}
		cur = next
	}
	return i
}
