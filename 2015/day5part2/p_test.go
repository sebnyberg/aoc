package p_test

import (
	"aoc/ax"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFileLines("input")
	res := run(lines)
	require.Equal(t, 55, res)
}

func run(lines chan string) int {
	var res int
	for line := range lines {
		if isNice(line) {
			res++
		}
	}
	return res
}

func Test_test(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"xxyxx", true},
		{"qjhvhtzxzqqjkmpb", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isNice(tc.s))
		})
	}
}

func isNice(s string) bool {
	var weaved bool
	var doubleTwice bool
	seen := make(map[string]struct{})
	for i := range s {
		if i > 1 {
			seen[s[i-2:i]] = struct{}{}
		}
		if i > 1 && i < len(s)-1 {
			if _, exists := seen[s[i:i+2]]; exists {
				doubleTwice = true
			}
		}
		if i > 1 {
			if s[i-2] == s[i] {
				weaved = true
			}
		}
	}
	return weaved && doubleTwice
}
