package day24

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day23part2 int

func BenchmarkDay23Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day23part2 = Part2(ax.MustReadFileLines("input"))
	}
}

func TestDay23Part2(t *testing.T) {
	assert.Equal(t, 13191913571211, Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) int {
	pairs := parsePairConds(rows)
	var res [14]byte
	for _, p := range pairs {
		if p.delta < 0 {
			p.first, p.second = p.second, p.first
			p.delta = -p.delta
		}
		// Now delta is guaranteed to be positive, so the first must be smaller than
		// the second
		res[p.first] = '1'
		res[p.second] = byte(p.delta + '1')
	}
	return ax.MustParseInt[int](string(res[:]))
}
