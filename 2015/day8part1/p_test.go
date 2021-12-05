package p_test

import (
	"aoc/ux"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ux.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 1342, res)
}

func run(lines chan string) int {
	var memSize int
	var codeSize int
	for line := range lines {
		codeSize += len(line)
		memSize += len(line) - 2
		line = strings.TrimSpace(line)
		line = line[1 : len(line)-1] // Skip quotes
		// We know the string is ASCII, so iterating as bytes is OK
		var pos int
		for pos < len(line) {
			if line[pos] != '\\' {
				pos++
				continue
			}
			if line[pos+1] == 'x' {
				memSize -= 3
				pos += 4
			} else {
				memSize -= 1
				pos += 2
			}
		}
	}
	return codeSize - memSize
}
