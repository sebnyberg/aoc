package day19

import (
	"aoc/ax"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day19part2 int

func BenchmarkDay19Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day19part1 = Part2(ax.MustReadFineLines("input"))
	}
}

func TestDay19Part2(t *testing.T) {
	assert.Equal(t, 3621, Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, 13192, Part2(ax.MustReadFineLines("input")))
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

	// Two scanners shares enough space if there exists a pair of beacons from
	// each scanner such that there are at least 11 shared vectors from those
	// beacons.
	sharesSpace := func(v1, v2 map[vectorHash]struct{}) bool {
		var count int
		for vec := range v1 {
			if _, exists := v2[vec]; exists {
				count++
				if count == 11 {
					return true
				}
			}
		}
		return false
	}

	// Perform BFS, starting with zeroth node, finding matching scanners and
	// adjusting their orientation such that the first set of points corresponds
	// to the orientation of the first scanner
	seen := make([]bool, nscanner)
	seenCount := 1
	seen[0] = true
	cur := []int{0}
	next := []int{}

	// For each pair of scanners, for each beacon, check if there is an
	// orientation for which there is a group of 12 shared beacons.
	scannerPos := make([]point, nscanner)
	scannerPos[0] = point{0, 0, 0}
	for seenCount != nscanner {
		next = next[:0]
		for _, rootScanner := range cur {
			for otherScanner := 0; otherScanner < nscanner; otherScanner++ {
				if seen[otherScanner] {
					continue
				}

				// For each beacon in first
				for rootBeacon := range vectors[rootScanner] {
					// If there exists a beacon + orientation in second such that there
					// are 11 shared vectors, then there is a match
					// Orientation of first doesn't matter, the second beacon is
					// exhaustively searched for all orientations
					firstVecs := vectors[rootScanner][rootBeacon][0]
					for otherBeacon := range vectors[otherScanner] {
						for orient := range vectors[otherScanner][otherBeacon] {
							if !sharesSpace(firstVecs, vectors[otherScanner][otherBeacon][orient]) {
								continue
							}
							// root and other scanner are within the same space. Shift the
							// orientation of the other scanner so that its aligned with root.

							// The root and other scanner are matching. The 'other' scanner
							// will now become a root for further iterations, so we adjust the
							// first orientation of each beacon so that it matches the root,
							// and also the position of each point as well. This will ensure
							// a shared field in the end.
							// Also adjust the point locations to align with the root
							p1 := points[rootScanner][rootBeacon][0]
							p2 := points[otherScanner][otherBeacon][orient]
							dx, dy, dz := p2.x-p1.x, p2.y-p1.y, p2.z-p1.z

							for otherBeacon := range vectors[otherScanner] {
								// Use the right orientation for vectors/points
								vectors[otherScanner][otherBeacon][0] = vectors[otherScanner][otherBeacon][orient]
								points[otherScanner][otherBeacon][0] = points[otherScanner][otherBeacon][orient]

								// Adjust locations
								p := points[otherScanner][otherBeacon][0]
								points[otherScanner][otherBeacon][0] = point{
									x: p.x - dx,
									y: p.y - dy,
									z: p.z - dz,
								}
							}

							// Part 2: keep track of scanner positions
							scannerPos[otherScanner] = point{-dx, -dy, -dz}

							next = append(next, otherScanner)
							seen[otherScanner] = true
							seenCount++
							goto ContinueSearch
						}
					}
				}
			ContinueSearch:
			}
		}

		cur, next = next, cur
	}

	// Calculate maximum manhattan distance
	var maxDist int16
	for first := 0; first < nscanner-1; first++ {
		for second := 0; second < nscanner; second++ {
			p1, p2 := scannerPos[first], scannerPos[second]
			maxDist = ax.Max(maxDist, ax.Abs(p2.x-p1.x)+ax.Abs(p2.y-p1.y)+ax.Abs(p2.z-p1.z))
		}
	}

	return int(maxDist)
}
