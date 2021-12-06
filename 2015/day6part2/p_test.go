package p_test

import (
	"aoc/ax"
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 15343601, res)
}

var prefix = regexp.MustCompile(`^(turn off|turn on|toggle) (\d+),(\d+) through (\d+),(\d+)$`)

func run(lines chan string) int {
	var brightness [1000][1000]int16
	for line := range lines {
		matches := prefix.FindStringSubmatch(line)
		if len(matches) != 6 {
			log.Fatalln("invalid number of matches")
		}
		action := matches[1]
		x1 := ax.MustParseIntBase[int](matches[2], 10)
		y1 := ax.MustParseIntBase[int](matches[3], 10)
		x2 := ax.MustParseIntBase[int](matches[4], 10)
		y2 := ax.MustParseIntBase[int](matches[5], 10)
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
					brightness[x][y]++
				case "turn off":
					brightness[x][y] = ax.Max(0, brightness[x][y]-1)
				case "toggle":
					brightness[x][y] += 2
				default:
					log.Fatalln(action)
				}
			}
		}
	}
	var res int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			res += int(brightness[x][y])
		}
	}
	return res
}
