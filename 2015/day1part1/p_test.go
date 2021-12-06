package p_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	res := run()
	require.Equal(t, 138, res)
}

func run() int {
	input := <-ax.MustReadFineLinesChan("input")
	floor := 0
	for _, ch := range input {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		} else {
			log.Fatalln("invalid character", string(ch))
		}
	}
	return floor
}
