package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	for i, tc := range []struct {
		in   string
		want any
	}{
		// {}
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, Parse(tc.in))
		})
	}
}

func Test_Solve1(t *testing.T) {
	for i, tc := range []struct {
		in   []any
		want any
	}{
		// {}
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, Solve1(tc.in))
		})
	}
}

func Test_Solve2(t *testing.T) {
	for i, tc := range []struct {
		in   []any
		want any
	}{
		// {}
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, Solve2(tc.in))
		})
	}
}
