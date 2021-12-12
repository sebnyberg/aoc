package p_test

import (
	"aoc/ax"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 26},
		{"input", 375},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

func run(rows []string) int {
	var count int
	for _, row := range rows {
		parts := strings.Split(row, "|")
		outputFields := strings.Fields(parts[1])
		for _, field := range outputFields {
			if len(field) == 2 || len(field) == 4 || len(field) == 3 || len(field) == 7 {
				count++
			}
		}
	}
	return count
}
