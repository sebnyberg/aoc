package day14

import (
	"math"
	"regexp"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day14part1 int

func BenchmarkDay14Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day14part1 = Part1(ax.MustReadFileLines("input"))
	}
}

func TestDay14Part1(t *testing.T) {
	assert.Equal(t, 1588, Part1(ax.MustReadFileLines("small")))
	assert.Equal(t, 2915, Part1(ax.MustReadFileLines("input")))
}

func Part1(rows []string) int {
	const maxDepth = 10

	firstRow := rows[0]

	// Parse pairs into map
	rows = rows[2:]
	pat := regexp.MustCompile(`^(\w{2}) -> (\w)$`)
	var pairs [26][26]byte
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		pair := parts[1]
		insert := parts[2]
		pairs[pair[0]-'A'][pair[1]-'A'] = insert[0] - 'A'
	}

	var pairCount [26][26]int
	var res [26]int
	res[firstRow[0]-'A']++
	for i := 1; i < len(firstRow); i++ {
		a, b := firstRow[i-1]-'A', firstRow[i]-'A'
		pairCount[a][b]++
		res[b]++
	}

	for i := 0; i < maxDepth; i++ {
		var nextPairCount [26][26]int
		for left := range pairCount {
			for right, count := range pairCount[left] {
				if count == 0 {
					continue
				}
				v := pairs[left][right]
				nextPairCount[left][v] += count
				nextPairCount[v][right] += count
				res[v] += count
			}
		}
		pairCount = nextPairCount
	}

	// Find max/min count
	var maxCount int
	minCount := math.MaxInt64
	for _, cnt := range res {
		if cnt > maxCount {
			maxCount = cnt
		}
		if cnt > 0 && cnt < minCount {
			minCount = cnt
		}
	}

	return maxCount - minCount
}
