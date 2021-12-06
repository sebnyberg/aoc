package p_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
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
		ribbonLen := 2 * ax.Min(l+w, ax.Min(w+h, h+l))
		volume := l * w * h
		res += ribbonLen + volume
	}
	return res
}
