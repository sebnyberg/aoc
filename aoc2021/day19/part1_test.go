package day19

import (
	"aoc/ax"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day19part1 int

func BenchmarkDay19Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day19part1 = Part1(ax.MustReadFineLines("input"))
	}
}

func TestDay19Part1(t *testing.T) {
	assert.Equal(t, 79, Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, 315, Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) int {
	// Parse scanner beacon points
	parseScannerBeaconPoints := func(rows []string) [][][48]point {
		var i, j int
		scannerBeaconPoints := make([][][48]point, 0)
		for i < len(rows) {
			i++ // skip scanner id
			scannerBeaconPoints = append(scannerBeaconPoints, make([][48]point, 0))
			for ; i < len(rows) && rows[i] != ""; i++ {
				parts := strings.Split(rows[i], ",")
				x := ax.MustParseInt[int](parts[0])
				y := ax.MustParseInt[int](parts[1])
				z := ax.MustParseInt[int](parts[2])
				p := point{x, y, z}
				scannerBeaconPoints[j] = append(scannerBeaconPoints[j], p.getVariations())
			}
			i++
			j++
		}
		return scannerBeaconPoints
	}
	points := parseScannerBeaconPoints(rows)

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

	// The first step is to create such a set of vectors:
	// [scanner][beacon][orient][vector] = count
	nscanner := len(points)
	vectors := make([][][48]map[vector]struct{}, nscanner)
	for scanner, beaconPoints := range points {

		// Create map of vectors per beacon and orientation
		nbeacons := len(points[scanner])
		vectors[scanner] = make([][48]map[vector]struct{}, nbeacons)

		// Initialize maps
		for beacon := 0; beacon < nbeacons; beacon++ {
			for orient := 0; orient < 48; orient++ {
				vectors[scanner][beacon][orient] = make(map[vector]struct{})
			}
		}

		// Calculate / add vectors for each pair of beacons
		for firstBeac := 0; firstBeac < nbeacons-1; firstBeac++ {
			for secondBeac := firstBeac + 1; secondBeac < nbeacons; secondBeac++ {
				// point values per orientation
				p1s, p2s := beaconPoints[firstBeac], beaconPoints[secondBeac]
				for orient := range p1s {
					p1ToP2 := p1s[orient].vecTo(p2s[orient])
					vectors[scanner][firstBeac][orient][p1ToP2] = struct{}{}
					p2ToP1 := p2s[orient].vecTo(p1s[orient])
					vectors[scanner][secondBeac][orient][p2ToP1] = struct{}{}
				}
			}
		}
	}

	// Two scanners shares enough space if there exists a pair of beacons from
	// each scanner such that there are at least 11 shared vectors from those
	// beacons.
	sharesSpace := func(v1, v2 map[vector]struct{}) bool {
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

	// Time to pair up scanners.
	// belongs := make([][]bool, nscanner)
	// for i := range belongs {
	// 	belongs[i] = make([]bool, nscanner)
	// }

	// Perform BFS, starting with zeroth node, finding matching scanners and
	// adjusting their orientation such that the first set of points corresponds
	// to the orientation of the first scanner
	seen := (1 << 0) // can just be 1 but shifting for clarity - this is a bitmask
	cur := []int{0}
	next := []int{}

	// For each pair of scanners, for each beacon, check if there is an
	// orientation for which there is a group of 12 shared beacons.
	for seen != (1<<nscanner)-1 {
		next = next[:0]
		for _, rootScanner := range cur {
			for otherScanner := 0; otherScanner < nscanner; otherScanner++ {
				if seen&(1<<otherScanner) > 0 {
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
							// root and other scanner shares space, shift the orientation of
							// the other scanner so that its orientation is aligned with the
							// root

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

							next = append(next, otherScanner)
							seen |= (1 << otherScanner) // mark as seen
							goto ContinueSearch
						}
					}
				}
			ContinueSearch:
			}
		}

		cur, next = next, cur
	}

	// Debugging time
	// Flatten points
	scannerBeacons := make([][]point, nscanner)
	for scanner := range points {
		nbeacon := len(points[scanner])
		scannerBeacons[scanner] = make([]point, nbeacon)
		for beacon := range points[scanner] {
			scannerBeacons[scanner][beacon] = points[scanner][beacon][0]
		}
		// Sort
		sortPoints(scannerBeacons[scanner])
	}

	// Create set of unique points
	uniquePoints := make(map[point]struct{})
	for scanner := range scannerBeacons {
		for _, p := range scannerBeacons[scanner] {
			uniquePoints[p] = struct{}{}
		}
	}
	uniquePointsList := ax.Keys(uniquePoints)
	sortPoints(uniquePointsList)
	// for _, p := range uniquePointsList {
	// 	// fmt.Printf("%v,%v,%v\n", p.x, p.y, p.z)
	// }

	return len(uniquePointsList)
}

type distTo struct {
	otherIdx int
	dist     int
}

func TestVariations(t *testing.T) {
	p := point{1, 2, 3}
	qq := p.getVariations()
	for i, q := range qq {
		if i == 0 {
			continue
		}
		if q == p {
			t.FailNow()
		}
	}
}

type point struct {
	x, y, z int
}

func (p point) intDistTo(q point) int {
	return p.x*q.x + p.y*q.y + p.z*q.z
}

// Note that vecTo is not commutative
func (p point) vecTo(q point) vector {
	return vector{q.x - p.x, q.y - p.y, q.z - p.z}
}

func (p point) getVariations() [48]point {
	var res [48]point
	x, y, z := p.x, p.y, p.z
	for i := 0; i < 8; i++ {
		j := i * 6
		res[j] = point{x, y, z}
		res[j+1] = point{x, z, y}
		res[j+2] = point{y, x, z}
		res[j+3] = point{y, z, x}
		res[j+4] = point{z, x, y}
		res[j+5] = point{z, y, x}
		if i&1 == 1 {
			for k := j; k < j+6; k++ {
				res[k].x = -res[k].x
			}
		}
		if i&2 == 2 {
			for k := j; k < j+6; k++ {
				res[k].y = -res[k].y
			}
		}
		if i&4 == 4 {
			for k := j; k < j+6; k++ {
				res[k].z = -res[k].z
			}
		}
	}
	return res
}

type vector struct {
	x, y, z int
}

func sortPoints(a []point) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			if a[i].y == a[j].y {
				return a[i].z < a[j].z
			}
			return a[i].y < a[j].y
		}
		return a[i].x < a[j].x
	})
}

func sortVectors(a []vector) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].x == a[j].x {
			if a[i].y == a[j].y {
				return a[i].z < a[j].z
			}
			return a[i].y < a[j].y
		}
		return a[i].x < a[j].x
	})
}

func debugPoints(vectors [][][24]map[vector]struct{}, sc1, b1, sc2, b2 int) {
	a := ax.Keys(vectors[sc1][b1][0])
	b := ax.Keys(vectors[sc2][b2][0])
	sortVectors(a)
	sortVectors(b)
	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Println()
	for i := range b {
		fmt.Println(b[i])
	}
}
