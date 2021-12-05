package p_test

import (
	"aoc/ux"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ux.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 15343601, res)
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
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isNice(tc.s))
		})
	}
}

var forbidden = map[string]struct{}{
	"ab": {},
	"cd": {},
	"pq": {},
	"xy": {},
}

func isNice(s string) bool {
	var twiceInARow bool
	var vowelCount int
	for i, ch := range s {
		if strings.ContainsRune("aeiou", ch) {
			vowelCount++
		}
		if i == 0 {
			continue
		}
		if s[i-1] == s[i] {
			twiceInARow = true
		}
		if _, exists := forbidden[s[i-1:i+1]]; exists {
			return false
		}
	}
	return twiceInARow && vowelCount >= 3
}
