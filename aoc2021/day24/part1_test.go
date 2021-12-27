package day24

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
	assert.Equal(t, 79197919993985, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	pairs := parsePairConds(rows)
	var res [14]byte
	for _, p := range pairs {
		if p.delta > 0 {
			p.first, p.second = p.second, p.first
			p.delta = -p.delta
		}
		// Now delta is guaranteed to be negative, so the first must be greater than
		// the second
		res[p.first] = '9'
		res[p.second] = byte(9 + p.delta + '0')
	}
	return ax.MustParseInt[int](string(res[:]))
}
