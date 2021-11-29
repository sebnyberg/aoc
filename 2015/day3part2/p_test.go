package p_test

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFileLines("input")
	res := run(lines)
	require.Equal(t, 2631, res)
}

func run(lines chan string) int {
	line := <-lines

	visited := make(map[[2]int]struct{})
	visited[[2]int{0, 0}] = struct{}{}
	res := 1
	var x, y [2]int
	for i, ch := range line {
		switch ch {
		case '>':
			x[i%2]++
		case '^':
			y[i%2]++
		case '<':
			x[i%2]--
		case 'v':
			y[i%2]--
		}
		k := [2]int{x[i%2], y[i%2]}
		if _, exists := visited[k]; !exists {
			res++
			visited[k] = struct{}{}
		}
	}
	return res
}
