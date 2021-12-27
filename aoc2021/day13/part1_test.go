package day13

import (
	"aoc/ax"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay13Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay13Part1(t *testing.T) {
	assert.Equal(t, "17", Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, "807", Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) string {
	// Parse points
	type point struct {
		x, y int
	}
	points := make([]point, 0, 100)
	for i, row := range rows {
		if row == "" {
			rows = rows[i+1:]
			break
		}
		parts := strings.Split(row, ",")
		x := ax.MustParseInt[int](parts[0])
		y := ax.MustParseInt[int](parts[1])
		points = append(points, point{x, y})
	}

	// Perform first fold
	var pat = regexp.MustCompile(`^fold along (\w+)=(\d+)$`)
	for _, row := range rows[:1] {
		match := pat.FindStringSubmatch(row)
		axis := match[1]
		val := ax.MustParseInt[int](match[2])
		if axis == "y" {
			for i := range points {
				if points[i].y > val {
					points[i].y = 2*val - (points[i].y)
					continue
				}
			}
		} else { // axis == "x"
			for i := range points {
				if points[i].x > val {
					points[i].x = 2*val - points[i].x
					continue
				}
			}
		}
	}

	return strconv.Itoa(len(ax.MapSet(points)))
}
