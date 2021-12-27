package day21

import (
	"aoc/ax"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day21part1 int

func BenchmarkDay21Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day21part1 = Part1(ax.MustReadFileLines("input"))
	}
}

func TestDay21Part1(t *testing.T) {
	assert.Equal(t, 739785, Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, 5044, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	pat := regexp.MustCompile(`^Player (\d+) starting position: (\d+)$`)
	parts1 := pat.FindStringSubmatch(rows[0])
	p1 := ax.MustParseInt[int](parts1[2])
	parts2 := pat.FindStringSubmatch(rows[1])
	p2 := ax.MustParseInt[int](parts2[2])
	pos := []int{p1 - 1, p2 - 1}
	pts := []int{0, 0}
	var i int
	dieVal := 1
	for pts[0] < 1000 && pts[1] < 1000 {
		mvmt := (3*dieVal + 3) % 10
		pos[i%2] = (pos[i%2] + mvmt) % 10
		pts[i%2] += pos[i%2] + 1
		dieVal += 3
		i++
	}

	winner, loser := pts[0], pts[1]
	if loser > winner {
		loser, winner = winner, loser
	}

	return loser * 3 * i
}
