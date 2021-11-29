package p_test

import (
	"aoc/ax"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFileLines("input")
	res := run(lines)
	require.Equal(t, 1586300, res)
}

func run(lines chan string) int {
	var res int
	for line := range lines {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		if l == 0 || w == 0 || h == 0 {
			panic("invalid line")
		}
		area := 2*l*w + 2*w*h + 2*h*l
		area += ax.Min(l*w, ax.Min(w*h, h*l))
		res += area
	}
	return res
}
