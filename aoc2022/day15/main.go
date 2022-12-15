package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

func solve1(inf string, yy int) any {
	lines := ax.MustReadFileLines(inf)
	// beacons := make(map[[2]int]bool)
	seen := make(map[int]bool)
	sensors := [][2][2]int{}
	for _, l := range lines {
		// fmt.Println(i)
		g := pat.FindStringSubmatch(l)
		// fmt.Println(g)
		// Find region of line that is closer to the beacon than its sensors
		// Closest possible distance is a straight line
		sx := ax.Atoi(g[1])
		sy := ax.Atoi(g[2])

		bx := ax.Atoi(g[3])
		by := ax.Atoi(g[4])

		sensors = append(sensors, [2][2]int{{sx, sy}, {bx, by}})
		if by == yy {
			seen[bx] = true
		}
	}

	var res int
	for _, s := range sensors {
		sx := s[0][0]
		sy := s[0][1]
		bx := s[1][0]
		by := s[1][1]
		maxDist := abs(sx-bx) + abs(sy-by)
		dy := abs(yy - sy)
		for dx := 0; dy+dx <= maxDist; dx++ {
			if !seen[sx-dx] {
				seen[sx-dx] = true
				res++
			}
			if !seen[sx+dx] {
				seen[sx+dx] = true
				res++
			}
		}
	}

	return res
}

func solve2(inf string, yy int) any {
	lines := ax.MustReadFileLines(inf)
	seen := make(map[int]bool)
	var sensors [][2][2]int
	for _, l := range lines {
		g := pat.FindStringSubmatch(l)
		sx := ax.Atoi(g[1])
		sy := ax.Atoi(g[2])

		bx := ax.Atoi(g[3])
		by := ax.Atoi(g[4])

		sensors = append(sensors, [2][2]int{{sx, sy}, {bx, by}})
		if by == yy {
			seen[bx] = true
		}
	}

	var res int
	type delta struct {
		x      int
		change int
	}
	for y := 0; y <= yy; y++ {
		// Collect intervals covered by sensors on this line
		var deltas []delta
		minX := math.MaxInt32
		maxX := math.MinInt32
		for _, s := range sensors {
			// The boundary of the sensor area is given by dbx + dby (distance
			// to beacon in x and y).
			// The interval covered on this line is given by the remaining
			// distance after removing dyy (y-distance to line).
			sx := s[0][0]
			sy := s[0][1]
			bx := s[1][0]
			by := s[1][1]
			bd := abs(sx-bx) + abs(sy-by)
			if abs(sy-y) > bd {
				continue
			}
			dx := bd - abs(sy-y)
			minX = ax.Min(minX, sx-dx)
			maxX = ax.Max(minX, sx+dx)
			deltas = append(deltas,
				delta{sx - dx, 1},
				delta{sx + dx + 1, -1},
			)
		}
		// Add sentinel values
		deltas = append(deltas,
			delta{minX - 2, 1},
			delta{minX - 1, -1},
			delta{maxX + 1, 1},
			delta{maxX + 2, -1},
		)
		sort.Slice(deltas, func(i, j int) bool {
			if deltas[i].x == deltas[j].x {
				return deltas[i].change < deltas[j].change
			}
			return deltas[i].x < deltas[j].x
		})
		var cur int
		for i, d := range deltas {
			cur += d.change
			if d.x < 0 || d.x > yy || cur > 0 {
				continue
			}
			if deltas[i+1].x != deltas[i].x {
				return d.x*4000000 + y
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	f := "input"
	yy := 4000000
	fmt.Printf("Result1:\n%v\n", solve1(f, yy))
	fmt.Printf("Result2:\n%v\n", solve2(f, yy))
}
