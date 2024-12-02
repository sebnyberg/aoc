package day22

import (
	"regexp"
	"sort"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day22part2 int

func BenchmarkDay22Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day22part1 = Part2(ax.MustReadFileLines("input"))
	}
}

func TestDay22Part2(t *testing.T) {
	assert.Equal(t, 39769202357779, Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, 1218645427221987, Part2(ax.MustReadFileLines("input")))
}

type cube struct {
	x1, x2,
	y1, y2,
	z1, z2 int
}

func Part2(rows []string) int {
	pat := regexp.MustCompile(`^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)
	cubes := make([]cube, len(rows))
	on := make([]bool, len(rows))
	xs := ax.Set[int]{}
	ys := ax.Set[int]{}
	zs := ax.Set[int]{}
	for i, row := range rows {
		parts := pat.FindStringSubmatch(row)
		x1 := ax.MustParseInt[int](parts[2])
		x2 := ax.MustParseInt[int](parts[3])
		y1 := ax.MustParseInt[int](parts[4])
		y2 := ax.MustParseInt[int](parts[5])
		z1 := ax.MustParseInt[int](parts[6])
		z2 := ax.MustParseInt[int](parts[7])
		xs.Add(x1, x2+1)
		ys.Add(y1, y2+1)
		zs.Add(z1, z2+1)
		cubes[i] = cube{x1, x2, y1, y2, z1, z2}
		on[i] = parts[1] == "on"
	}
	xsList := ax.Keys(xs)
	ysList := ax.Keys(ys)
	zsList := ax.Keys(zs)
	sort.Ints(xsList)
	sort.Ints(ysList)
	sort.Ints(zsList)
	xIdx := make(map[int]int, len(xsList))
	yIdx := make(map[int]int, len(ysList))
	zIdx := make(map[int]int, len(zsList))
	for i, val := range xsList {
		xIdx[val] = i
	}
	for j, val := range ysList {
		yIdx[val] = j
	}
	for k, val := range zsList {
		zIdx[val] = k
	}
	grid := make([][][]bool, len(xs))
	for x := range grid {
		grid[x] = make([][]bool, len(ys))
		for y := range grid[x] {
			grid[x][y] = make([]bool, len(zs))
		}
	}
	for i, c := range cubes {
		for xi := xIdx[c.x1]; xi < xIdx[c.x2+1]; xi++ {
			for yi := yIdx[c.y1]; yi < yIdx[c.y2+1]; yi++ {
				for zi := zIdx[c.z1]; zi < zIdx[c.z2+1]; zi++ {
					grid[xi][yi][zi] = on[i]
				}
			}
		}
	}
	var res int
	for xi := 0; xi < len(xsList); xi++ {
		for yi := 0; yi < len(ysList); yi++ {
			for zi := 0; zi < len(zsList); zi++ {
				if grid[xi][yi][zi] {
					dx := xsList[xi+1] - xsList[xi]
					dy := ysList[yi+1] - ysList[yi]
					dz := zsList[zi+1] - zsList[zi]
					res += dx * dy * dz
				}
			}
		}
	}
	return res
}
