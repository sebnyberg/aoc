package day23

import (
	"container/heap"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day23part1 int

func BenchmarkDay23Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day23part1 = Part1(ax.MustReadFileLines("input"))
	}
}

func TestDay23Part1(t *testing.T) {
	assert.Equal(t, 12521, Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, 16244, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	var start stateEnergy
	start.rooms[0] = string([]byte{rows[2][3], rows[3][3]})
	start.rooms[1] = string([]byte{rows[2][5], rows[3][5]})
	start.rooms[2] = string([]byte{rows[2][7], rows[3][7]})
	start.rooms[3] = string([]byte{rows[2][9], rows[3][9]})
	start.hallway = "..........."

	h := StateHeap{start}
	seen := make(ax.Set[state], 100)
	minDist := make(map[state]int)
	minDist[start.state] = 0
	for len(h) > 0 {
		s := heap.Pop(&h).(stateEnergy)
		if seen.Has(s.state) {
			continue
		}
		if s.valid() {
			return s.energy
		}
		moves := append(getMovesFromHallway(s), getMovesFromRooms(s)...)
		for _, m := range moves {
			if _, exists := minDist[m.state]; !exists {
				minDist[m.state] = m.energy
				heap.Push(&h, m)
				continue
			}
			if seen.Has(m.state) || minDist[m.state] <= m.energy {
				continue
			}
			minDist[m.state] = m.energy
			heap.Push(&h, m)
		}
	}

	return 0
}
