package p_test

import (
	"aoc/ux"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ux.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 3737498, res)
}

func run(lines chan string) int {
	var res int
	for line := range lines {
		var l, w, h int
		fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		if l == 0 || w == 0 || h == 0 {
			panic("invalid line")
		}
		ribbonLen := 2 * ux.Min(l+w, ux.Min(w+h, h+l))
		volume := l * w * h
		res += ribbonLen + volume
	}
	return res
}
