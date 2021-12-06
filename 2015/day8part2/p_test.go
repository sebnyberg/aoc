package p_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 2074, res)
}

func run(lines chan string) int {
	var origSize int
	var encodedSize int
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789"
	for line := range lines {
		origSize += len(line)
		encodedSize += len(line) + 4
		line = strings.TrimSpace(line)
		line = line[1 : len(line)-1] // Skip quotes
		// We know the string is ASCII, so iterating as bytes is OK
		for _, ch := range line {
			if strings.ContainsRune(alphabet, ch) {
				continue
			}
			encodedSize += 1
		}
	}
	return encodedSize - origSize
}
