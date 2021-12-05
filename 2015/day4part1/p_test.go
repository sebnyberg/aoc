package p_test

import (
	"aoc/ux"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ux.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 282749, res)
}

func run(lines chan string) int {
	key := <-lines

	hash := md5.New()
	for x := 0; x < 1e9; x++ {
		n := strconv.Itoa(x)
		input := key + n
		hash.Write([]byte(input))
		hexOutput := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(hexOutput, "00000") {
			return x
		}
		hash.Reset()
	}
	panic("couldn't find the hash")
}
