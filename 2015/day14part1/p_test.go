package p_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		time  int
		want  int
	}{
		{"small", 1, 1},
		{"small", 1000, 689},
		{"input", 2503, 601},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines, tc.time))
		})
	}
}

var pat = regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.$`)

func run(lines []string, time int) int {
	deltas := make([][]int8, len(lines))
	periods := make([]int, len(lines))
	for i, line := range lines {
		parts := pat.FindStringSubmatch(line)
		speed := ax.MustParseIntBase(parts[2], 10)
		duration := ax.MustParseIntBase(parts[3], 10)
		restTime := ax.MustParseIntBase(parts[4], 10)
		period := duration + restTime
		periods[i] = period
		deltas[i] = make([]int8, period)
		for t := 0; t < duration; t++ {
			deltas[i][t] = int8(speed)
		}
	}

	points := make([]int, len(lines))
	positions := make([]int, len(lines))
	for t := 0; t < time; t++ {
		var maxPos int
		for reindeer, delta := range deltas {
			positions[reindeer] += int(delta[t%periods[reindeer]])
			maxPos = ax.Max(maxPos, positions[reindeer])
		}
		for reindeer, pos := range positions {
			if pos == maxPos {
				points[reindeer]++
			}
		}
	}
	var maxPoints int
	for _, p := range points {
		maxPoints = ax.Max(maxPoints, p)
	}

	return maxPoints
}
