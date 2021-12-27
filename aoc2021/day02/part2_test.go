package day02

import (
	"aoc/ax"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay02Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay02Part2(t *testing.T) {
	assert.Equal(t, "900", Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, "1592426537", Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) string {
	pat := regexp.MustCompile(`^(\w+) (\d+)`)
	var horz, depth, aim int
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		dir := parts[1]
		val := ax.MustParseInt[int](parts[2])
		switch dir {
		case "forward":
			horz += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return strconv.Itoa(depth * horz)
}
