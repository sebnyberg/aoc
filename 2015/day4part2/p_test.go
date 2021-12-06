package p_test

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
	res := part(lines)
	require.Equal(t, 9962624, res)
}

func part(lines chan string) int {
	key := <-lines

	hash := md5.New()
	for x := 0; x < 1e9; x++ {
		n := strconv.Itoa(x)
		input := key + n
		hash.Write([]byte(input))
		hexOutput := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(hexOutput, "000000") {
			return x
		}
		hash.Reset()
	}
	panic("couldn't find the hash")
}
