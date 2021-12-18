package day18

import (
	"aoc/ax"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day18part1 int

func BenchmarkDay18Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// day18part1 = Part1(ax.MustReadFineLines("input")[0])
	}
}

func TestDay18Part1(t *testing.T) {
	// assert.Equal(t, 11185, Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, 11185, Part1(ax.MustReadFineLines("small2")))
	// assert.Equal(t, 11185, Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) int {
	vals, levels := ParseRow(rows[0])
	for i := 1; i < len(rows); i++ {
		vals2, levels2 := ParseRow(rows[i])
		vals = append(vals, vals2...)
		levels = append(levels, levels2...)
		for i := range levels {
			levels[i]++
		}
		vals, levels = reduce(vals, levels)
	}
	res := toString(vals, levels)
	fmt.Println(res)
	return 0
}

func toString(vals, levels []int) string {
	var level int
	var sb strings.Builder
	levels = append(levels, 0)
	for i := range vals {
		for levels[i] > level {
			level++
			sb.WriteByte('[')
		}
		sb.WriteByte(byte('0' + vals[i]))
		for levels[i+1] < level {
			sb.WriteByte(']')
			level--
		}
		sb.WriteByte(',')
	}
	ss := sb.String()
	ss = ss[:len(ss)-1]
	for level > 0 {
		ss += "]"
		level--
	}
	return ss
}

func reduce(vals, levels []int) ([]int, []int) {
	var idx int
	n := len(vals)
	for idx < n {
		if levels[idx] == 5 { // explode
			// try to move left
			nextIdx := idx
			if idx > 0 {
				vals[idx-1] += vals[idx]
				nextIdx--
			}
			vals[idx] = 0
			levels[idx]--
			// try to move right
			if idx < n-2 {
				vals[idx+2] += vals[idx+1]
			}
			vals[idx+1] = 0
			levels[idx+1]--
			// Merge if necessary
			mergeLeft := idx > 0 && levels[idx-1] == levels[idx]
			mergeRight := idx < n-1 && levels[idx+2] == levels[idx+1]
			if mergeRight {
				copy(vals[idx+1:], vals[idx+2:])
				copy(levels[idx+1:], levels[idx+2:])
				n--
			}
			if mergeLeft {
				copy(vals[idx:], vals[idx+1:])
				copy(levels[idx:], levels[idx+1:])
				n--
			}
			levels = levels[:n]
			vals = vals[:n]
			idx = nextIdx
		} else if vals[idx] >= 10 { // split
			vals = append(vals, 0)
			levels = append(levels, 0)
			// make space
			copy(vals[idx+1:], vals[idx:])
			copy(levels[idx+1:], levels[idx:])
			vals[idx+1] = vals[idx]/2 + vals[idx]%2
			vals[idx] = vals[idx] / 2
			levels[idx]++
			levels[idx+1]++
			n++
		} else {
			idx++
		}
	}
	return vals, levels
}

func ParseRow(row string) ([]int, []int) {
	vals := make([]int, 0)
	levels := make([]int, 0)
	var level int
	for _, ch := range row {
		switch ch {
		case '[':
			level++
		case ']':
			level--
		case ',':
		default:
			levels = append(levels, level)
			vals = append(vals, int(ch-'0'))
		}
	}
	return vals, levels
}
