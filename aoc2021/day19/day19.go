package day19

import (
	"aoc/ax"
	"fmt"
	"sort"
	"strings"
)

type vectorHash int

func parseVectors(points [][][48]point) [][][48]map[vectorHash]struct{} {
	nscanner := len(points)
	vectors := make([][][48]map[vectorHash]struct{}, nscanner)
	for scanner, beaconPoints := range points {

		// Create map of vectors per beacon and orientation
		nbeacons := len(points[scanner])
		vectors[scanner] = make([][48]map[vectorHash]struct{}, nbeacons)

		// Initialize maps
		for beacon := 0; beacon < nbeacons; beacon++ {
			for orient := 0; orient < 48; orient++ {
				vectors[scanner][beacon][orient] = make(map[vectorHash]struct{})
			}
		}

		// Calculate / add vectors for each pair of beacons
		for firstBeac := 0; firstBeac < nbeacons-1; firstBeac++ {
			for secondBeac := firstBeac + 1; secondBeac < nbeacons; secondBeac++ {
				// point values per orientation
				p1s, p2s := beaconPoints[firstBeac], beaconPoints[secondBeac]
				for orient := range p1s {
					p1ToP2 := p1s[orient].vecTo(p2s[orient])
					vectors[scanner][firstBeac][orient][p1ToP2.hash()] = struct{}{}
					p2ToP1 := p2s[orient].vecTo(p1s[orient])
					vectors[scanner][secondBeac][orient][p2ToP1.hash()] = struct{}{}
				}
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
		scannerBeaconPoints = append(scannerBeaconPoints, make([][48]point, 0))
		for ; i < len(rows) && rows[i] != ""; i++ {
			parts := strings.Split(rows[i], ",")
			x := ax.MustParseInt[int16](parts[0])
			y := ax.MustParseInt[int16](parts[1])
			z := ax.MustParseInt[int16](parts[2])
			p := point{x, y, z}
			scannerBeaconPoints[j] = append(scannerBeaconPoints[j], p.getOrientations())
		}
		i++
		j++
	}
	return scannerBeaconPoints
}

type point struct {
	x, y, z int16
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
	x, y, z int16
}

const (
	hashOffset1 = 5000
	hashOffset2 = 5000 * 5000
)

func (v vector) hash() vectorHash {
	res := vectorHash(v.x + 2500)
	res += 5000 * vectorHash(v.y+2500)
	res += 5000 * 5000 * vectorHash(v.z+2500)
	return res
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
