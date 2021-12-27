package day19

import (
	"aoc/ax"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day19part1 int

func BenchmarkDay19Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day19part1 = Part1(ax.MustReadFileLines("input"))
	}
}

func TestDay19Part1(t *testing.T) {
	assert.Equal(t, 79, Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, 315, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	// Parse points with orientations from rows
	points := parsePoints(rows)

	// For an integer number of rotations, the x,y, and z can be in different
	// positions (swapped), which yields 3! => 3*2*1 = 6 different permutations
	// Also, the orientation may be either positive or negative for each number,
	// which yields 2^3 => 8 different perspectives. In total this means 48
	// orientations.

	// Even though the scanner's position is unknown, the beacons relative
	// distances are static.

	// For two scanners to share a field of beacons, there must be a pair of
	// orientations for which at least 11 vectors from a certain point are the
	// same.

	// Create a set of vectors for each point compared to all other points for
	// a given orientation.
	vectors := parseVectors(points)
	nscanner := len(points)

	// Perform BFS, starting with zeroth node, finding matching scanners and
	// adjusting their orientation such that the first set of points corresponds
	// to the orientation of the first scanner
	seen := make([]bool, nscanner)
	seenCount := 1
	seen[0] = true
	cur := []int{0}
	next := []int{}
	var nextMtx sync.Mutex

	scannerPos := make([]point, nscanner)
	scannerPos[0] = point{0, 0, 0}

	// For each pair of scanners, for each beacon, check if there is an
	// orientation for which there is a group of 12 shared beacons.
	for seenCount != nscanner {
		next = next[:0]
		for _, rootScanner := range cur {
			var wg sync.WaitGroup
			for otherScanner := 0; otherScanner < nscanner; otherScanner++ {
				if seen[otherScanner] {
					continue
				}
				wg.Add(1)
				go func(other int) {
					defer wg.Done()
					if compareScanners(vectors, points, scannerPos, rootScanner, other) {
						seen[other] = true
						nextMtx.Lock()
						next = append(next, other)
						seenCount++
						nextMtx.Unlock()
					}
				}(otherScanner)
			}
			wg.Wait()
		}

		cur, next = next, cur
	}

	// Create set of unique points
	uniquePoints := make(map[point]struct{})
	for scanner := range points {
		for beacon := range points[scanner] {
			uniquePoints[points[scanner][beacon][0]] = struct{}{}
		}
	}

	return len(uniquePoints)
}

func TestVariations(t *testing.T) {
	p := point{1, 2, 3}
	qq := p.getOrientations()
	for i, q := range qq {
		if i == 0 {
			continue
		}
		if q == p {
			t.FailNow()
		}
	}
}
