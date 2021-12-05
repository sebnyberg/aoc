package p_test

import (
	"aoc/ux"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	line := <-ux.MustReadFineLinesChan("input")
	res := run(line)
	require.Equal(t, 252594, res)
}

func run(line string) int {
	cur := make([]int8, len(line))
	for i := range line {
		cur[i] = int8(line[i] - '0')
	}

	// Perform transformation
	next := make([]int8, 0, len(line))
	for i := 0; i < 40; i++ {
		next = next[:0]
		cur = append(cur, -1) // sentinel value to pop last run of digits
		var count int
		for i := range cur {
			if i > 0 && cur[i] != cur[i-1] {
				next = append(next, int8(count), cur[i-1])
				count = 1
			} else {
				count++
			}
		}

		cur, next = next, cur
	}
	return len(cur)
}
