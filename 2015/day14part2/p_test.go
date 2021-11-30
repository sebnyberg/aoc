package p_test

import (
	"aoc/ax"
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
		{"small", 1, 16},
		{"small", 10, 160},
		{"small", 11, 176},
		{"small", 138, 176},
		{"small", 145, 252},
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
	var maxDist int
	for _, line := range lines {
		parts := pat.FindStringSubmatch(line)
		// Name SHOULD matter but it doesn't :(
		// name := parts[1]
		speed := ax.MustParseInt(parts[2], 10)
		duration := ax.MustParseInt(parts[3], 10)
		restTime := ax.MustParseInt(parts[4], 10)
		period := duration + restTime
		var dist int
		if time > period {
			iterations := time / period
			dist = iterations * speed * duration
		}
		if d := time % period; d < duration {
			dist += d * speed
		} else {
			dist += speed * duration
		}
		maxDist = ax.Max(maxDist, dist)
	}

	return maxDist
}
