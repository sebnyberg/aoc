package p_test

import (
	"aoc/ux"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	line := <-ux.MustReadFineLinesChan("input")
	res := run(line)
	require.Equal(t, 1, res)
}

var digitPat = regexp.MustCompile(`([-]?\d+)`)

func run(line string) int {
	var sum int
	for _, match := range digitPat.FindAllString(line, -1) {
		sum += ux.MustParseIntBase(match, 10)
	}
	return sum
}
