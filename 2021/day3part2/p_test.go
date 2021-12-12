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
		{"small", 198},
		{"input", 775304},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

func run(lines []string) int {
	n := len(strings.TrimSpace(lines[0]))
	m := len(lines)
	bitLines := make([]int, m)
	for i, line := range lines {
		bitLines[i] = ax.MustParseIntBase(line, 2)
	}
	oxyLines := make(map[int]struct{})
	coLines := make(map[int]struct{})
	for _, k := range bitLines {
		oxyLines[k] = struct{}{}
		coLines[k] = struct{}{}
	}
	var coRes, oxyRes int
	for bit := 1 << (n - 1); bit > 0; bit >>= 1 {
		var oxyCount int
		for oxyLine := range oxyLines {
			if oxyLine&bit > 0 {
				oxyCount++
			}
		}
		if oxyCount*2 >= len(oxyLines) {
			for k := range oxyLines {
				if k&bit == 0 {
					delete(oxyLines, k)
				}
			}
		} else {
			for k := range oxyLines {
				if k&bit > 0 {
					delete(oxyLines, k)
				}
			}
		}
		if oxyCount == 1 {
			for k := range oxyLines {
				oxyRes = k
			}
		}
		var coCount int
		for coLine := range coLines {
			if coLine&bit > 0 {
				coCount++
			}
		}
		if coCount*2 >= len(coLines) {
			for k := range coLines {
				if k&bit > 0 {
					delete(coLines, k)
				}
			}
		} else {
			for k := range coLines {
				if k&bit == 0 {
					delete(coLines, k)
				}
			}
		}
		if coCount == 1 {
			for k := range coLines {
				coRes = k
			}
		}
		fmt.Printf("co:\n")
		for k := range coLines {
			fmt.Printf("%0b\n", k)
		}
		fmt.Println()
	}
	return oxyRes * coRes
}
