package day19

import (
	"aoc/ax"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day19part2 int

func BenchmarkDay19Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day19part1 = Part2(ax.MustReadFileLines("input"))
	}
}

func TestDay19Part2(t *testing.T) {
	assert.Equal(t, 3621, Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, 13192, Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) int {
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
	var seenCount uint64 = 1
	seen[0] = true
	cur := []int{0}
	next := []int{}
	var nextMtx sync.Mutex

	scannerPos := make([]point, nscanner)
	scannerPos[0] = point{0, 0, 0}

	// For each pair of scanners, for each beacon, check if there is an
	// orientation for which there is a group of 12 shared beacons.
	for seenCount != uint64(nscanner) {
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
						atomic.AddUint64(&seenCount, 1)
						nextMtx.Unlock()
					}
				}(otherScanner)
			}
			wg.Wait()
		}

		cur, next = next, cur
	}

	// Calculate maximum manhattan distance
	var maxDist int
	for first := 0; first < nscanner-1; first++ {
		for second := 0; second < nscanner; second++ {
			p1, p2 := scannerPos[first], scannerPos[second]
			maxDist = ax.Max(maxDist, ax.Abs(p2.x-p1.x)+ax.Abs(p2.y-p1.y)+ax.Abs(p2.z-p1.z))
		}
	}

	return int(maxDist)
}
