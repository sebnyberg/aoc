package day19

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

type vectorHash uint64

func parseVectors(points [][][48]point) [][][48][]vectorHash {
	nscanner := len(points)
	vectors := make([][][48][]vectorHash, nscanner)
	minVectors := math.MaxInt32
	maxVectors := 0
	for scanner, beaconPoints := range points {

		// Create map of vectors per beacon and orientation
		nbeacons := len(points[scanner])
		vectors[scanner] = make([][48][]vectorHash, nbeacons)

		// Initialize maps
		for beacon := 0; beacon < nbeacons; beacon++ {
			for orient := 0; orient < 48; orient++ {
				vectors[scanner][beacon][orient] = make([]vectorHash, 0, 27)
			}
		}

		// Calculate / add vectors for each pair of beacons
		minVectors = ax.Min(minVectors, nbeacons-1)
		maxVectors = ax.Max(maxVectors, nbeacons-1)
		for firstBeac := 0; firstBeac < nbeacons-1; firstBeac++ {
			for secondBeac := firstBeac + 1; secondBeac < nbeacons; secondBeac++ {
				// point values per orientation
				p1s, p2s := beaconPoints[firstBeac], beaconPoints[secondBeac]
				for orient := range p1s {
					p1ToP2 := p1s[orient].vecTo(p2s[orient])
					// vectors[scanner][firstBeac][orient][p1ToP2.hash()] = struct{}{}
					vectors[scanner][firstBeac][orient] = append(vectors[scanner][firstBeac][orient], p1ToP2.hash())
					p2ToP1 := p2s[orient].vecTo(p1s[orient])
					// vectors[scanner][secondBeac][orient][p2ToP1.hash()] = struct{}{}
					vectors[scanner][secondBeac][orient] = append(vectors[scanner][secondBeac][orient], p2ToP1.hash())
				}
			}
		}
		// Sort vectors for faster comparisons
		for beacon := 0; beacon < nbeacons; beacon++ {
			for orient := 0; orient < 48; orient++ {
				sort.Slice(vectors[scanner][beacon][orient], func(i, j int) bool {
					return vectors[scanner][beacon][orient][i] < vectors[scanner][beacon][orient][j]
				})
			}
		}
	}
	return vectors
}

func parsePoints(rows []string) [][][48]point {
	var i, j int
	scannerBeaconPoints := make([][][48]point, 0)
	for i < len(rows) {
		i++ // skip scanner id
		scannerBeaconPoints = append(scannerBeaconPoints, make([][48]point, 0, 28))
		for ; i < len(rows) && rows[i] != ""; i++ {
			parts := strings.Split(rows[i], ",")
			x := ax.MustParseInt[int](parts[0])
			y := ax.MustParseInt[int](parts[1])
			z := ax.MustParseInt[int](parts[2])
			p := point{x, y, z}
			scannerBeaconPoints[j] = append(scannerBeaconPoints[j], p.getOrientations())
		}
		i++
		j++
	}
	return scannerBeaconPoints
}

func compareScanners(
	vectors [][][48][]vectorHash,
	points [][][48]point,
	scannerPos []point,
	rootScanner, otherScanner int,
) bool {
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
				return true
			}
		}
	}
	return false
}

// Two scanners shares enough space if there exists a pair of beacons from
// each scanner such that there are at least 11 shared vectors from those
// beacons.
func sharesSpace(v1, v2 []vectorHash) bool {
	var count int
	var l, r int
	for l != len(v1) && r != len(v2) && ax.Min(len(v1), len(v2))+count >= 11 {
		if v1[l] == v2[r] {
			count++
			if count == 11 {
				return true
			}
			l++
			r++
		} else {
			if v1[l] < v2[r] {
				l++
			} else if v2[r] < v1[l] {
				r++
			}
		}
	}
	return false
}

type point struct {
	x, y, z int
}

// Note that vecTo is not commutative
func (p point) vecTo(q point) vector {
	return vector{q.x - p.x, q.y - p.y, q.z - p.z}
}

// getOrientations returns 48 orientations of point (could be reduced to 24)
func (p point) getOrientations() [48]point {
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

const (
	hashOffset1 = 5000
	hashOffset2 = 5000 * 5000
)

func (v vector) hash() vectorHash {
	return vectorHash((v.x + 2500) + 5000*(v.y+2500) + 5000*5000*(v.z+2500))
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
