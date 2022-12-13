package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_less(t *testing.T) {
	for i, tc := range []struct {
		s1   string
		s2   string
		want bool
	}{
		{"[[1],[2,3,4]]", "[[1],4]", true},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", false},
		{"[[[]]]", "[[]]", false},
		{"[[]]", "[[[]]]", true},
		{"[]", "[3]", true},
		{"[7,7,7,7]", "[7,7,7]", false},
		{"[[4,4],4,4]", "[[4,4],4,4,4]", true},
		{"[9]", "[8,7,6]", false},
		{"[1,1,3,1,1]", "[1,1,5,1,1]", true},
	} {
		t.Run(fmt.Sprint(i, tc.s1, tc.s2), func(t *testing.T) {
			l1 := parseList(tc.s1)
			l2 := parseList(tc.s2)
			require.Equal(t, tc.want, l1.less(l2))
		})
	}
}
