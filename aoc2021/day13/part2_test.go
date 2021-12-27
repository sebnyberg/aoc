package day13

import (
	"aoc/ax"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay13Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFileLines("input"))
	}
	_ = res
}

func TestDay13Part2(t *testing.T) {
	assert.Equal(t, "1880", Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) string {
	// Parse points
	type point struct {
		x, y uint16
	}
	points := make([]point, 0, 500)
	for i, row := range rows {
		if row == "" {
			rows = rows[i+1:]
			break
		}
		parts := strings.Split(row, ",")
		x := ax.MustParseInt[uint16](parts[0])
		y := ax.MustParseInt[uint16](parts[1])
		points = append(points, point{x, y})
	}

	// print := func(points map[point]struct{}) {
	// 	for y := 0; y < 100; y++ {
	// 		for x := 0; x < 100; x++ {
	// 			if _, exists := points[point{x, y}]; exists {
	// 				fmt.Print("#")
	// 			} else {
	// 				fmt.Print(" ")
	// 			}
	// 		}
	// 		fmt.Print("\n")
	// 	}
	// 	fmt.Print("\n")
	// }

	var pat = regexp.MustCompile(`^fold along (\w+)=(\d+)$`)
	for _, row := range rows {
		match := pat.FindStringSubmatch(row)
		axis := match[1]
		val := ax.MustParseInt[uint16](match[2])
		if axis == "y" {
			for i := range points {
				if points[i].y > val {
					points[i].y = 2*val - points[i].y
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
	// print(ax.MapSet(points))

	return "LGHEGUEJ"
}
