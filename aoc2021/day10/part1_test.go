package day10

import (
	"strconv"
	"strings"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay10Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay10Part1(t *testing.T) {
	assert.Equal(t, "26397", Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, "366027", Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) string {
	scoreRow := func(row string) int {
		want := make([]byte, 0)
		n := 0
		for _, ch := range row {
			if strings.ContainsRune("[({<", ch) {
				if ch == '(' {
					want = append(want, ')')
				} else {
					want = append(want, byte(ch+2))
				}
				n++
				continue
			}
			if byte(ch) != want[n-1] {
				switch ch {
				case ')':
					return 3
				case ']':
					return 57
				case '}':
					return 1197
				case '>':
					return 25137
				}
			}
			want = want[:n-1] // pop
			n--
		}
		return 0
	}

	var res int
	for _, row := range rows {
		res += scoreRow(row)
	}
	return strconv.Itoa(res)
}
