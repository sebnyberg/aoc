package p_test

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	line := <-ax.MustReadFineLinesChan("input")
	res := run(line)
	require.Equal(t, "hepxxyzz", res[0])
}

func run(line string) string {
	// Exactly 8 lowercase letters
	// Increment old password until it is valid
	// Must include one increasing straight of at least three letters
	// May not contain iol
	// Must contain at least two different non-overlapping pairs of letters like
	// 'aa', 'bb' or 'zz'
	// Step 1: consider whether the current input can be considered a single int
	// 26~=32==2^5
	// 8 characters => 40 bits => OK with 64-bit integer
	next := func(s []byte) []byte {
		var carry bool
		s[len(s)-1]++
		for i := len(s) - 1; i >= 0; i-- {
			if carry {
				s[i]++
				carry = false
			}
			if s[i] == 26+'a' {
				s[i] = 'a'
				carry = true
			}
		}
		return s
	}
	valid := func(s []byte) bool {
		var straight, seenDouble, twoPair bool
		for i := range s {
			// first criterion
			if i > 1 && s[i] == s[i-1]+1 && s[i-1] == s[i-2]+1 {
				straight = true
			}
			// second criterion
			if s[i] == 'i' || s[i] == 'o' || s[i] == 'l' {
				return false
			}
			// third criterion
			if i > 1 && s[i-1] == s[i-2] {
				seenDouble = true
			}
			if i < len(s)-1 && seenDouble && s[i] == s[i+1] {
				twoPair = true
			}
		}
		return straight && twoPair
	}
	pw := []byte(line)
	for pw = next(pw); !valid(pw); pw = next(pw) {
	}
	return string(pw)
}
