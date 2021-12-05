package p_test

import (
	"aoc/ux"
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ux.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 400410, res)
}

var prefix = regexp.MustCompile(`^(turn off|turn on|toggle) (\d+),(\d+) through (\d+),(\d+)$`)

func run(lines chan string) int {
	var grid [1000][1000]bool
	for line := range lines {
		matches := prefix.FindStringSubmatch(line)
		if len(matches) != 6 {
			log.Fatalln("invalid number of matches")
		}
		action := matches[1]
		x1 := ux.MustParseIntBase(matches[2], 10)
		y1 := ux.MustParseIntBase(matches[3], 10)
		x2 := ux.MustParseIntBase(matches[4], 10)
		y2 := ux.MustParseIntBase(matches[5], 10)
		if x2 < x1 {
			x1, x2 = x2, x1
		}
		if y2 < y1 {
			y1, y2 = y2, y1
		}
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				switch action {
				case "turn on":
					grid[x][y] = true
				case "turn off":
					grid[x][y] = false
				case "toggle":
					grid[x][y] = !grid[x][y]
				default:
					log.Fatalln(action)
				}
			}
		}
	}
	var res int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				res++
			}
		}
	}
	return res
}
