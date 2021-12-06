package y2015day3_test

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 2572, res)
}

func run(lines chan string) int {
	line := <-lines

	visited := make(map[[2]int]struct{})
	visited[[2]int{0, 0}] = struct{}{}
	res := 1
	var x, y int
	for _, ch := range line {
		switch ch {
		case '>':
			x++
		case '^':
			y++
		case '<':
			x--
		case 'v':
			y--
		}
		k := [2]int{x, y}
		if _, exists := visited[k]; !exists {
			res++
			visited[k] = struct{}{}
		}
	}
	return res
}
