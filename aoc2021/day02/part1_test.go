package day02

import (
	"aoc/ax"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay02Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay02Part1(t *testing.T) {
	assert.Equal(t, "150", Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, "1524750", Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) string {
	pat := regexp.MustCompile(`^(\w+) (\d+)`)
	var horz, depth int
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		dir := parts[1]
		val := ax.MustParseInt[int](parts[2])
		switch dir {
		case "forward":
			horz += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return strconv.Itoa(depth * horz)
}
