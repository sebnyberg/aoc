package day14

import (
	"aoc/ax"
	"math"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day14part2res int

func BenchmarkDay14Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day14part2res = Part2(ax.MustReadFineLines("input"))
	}
}

func TestDay14Part2(t *testing.T) {
	assert.Equal(t, 2188189693529, Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, 3353146900153, Part2(ax.MustReadFineLines("input")))
}

func Part2(rows []string) int {
	const maxDepth = 40

	firstRow := rows[0]

	// Parse pairs into map
	rows = rows[2:]
	pat := regexp.MustCompile(`^(\w{2}) -> (\w)$`)
	pairs := make(map[[2]byte]byte, len(rows))
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		pair := parts[1]
		insert := parts[2]
		pairs[[2]byte{pair[0], pair[1]}] = insert[0]
	}

	// Helper to merge to [26]ints together
	// [26]int is used to count occurrences of characters
	merge := func(a, b *[26]int) *[26]int {
		for i, count := range b {
			a[i] += count
		}
		return a
	}

	// Create memoization map, capturing (pair,depth) -> character count
	type memKey struct {
		pair  [2]byte
		depth int
	}
	mem := make(map[memKey][26]int)

	// countBetween visits a pair, counting occurrences of characters between
	// the characters in the pair for the current and subsequent levels.
	var countBetween func(depth int, pair [2]byte) [26]int
	countBetween = func(depth int, pair [2]byte) [26]int {
		if depth == maxDepth {
			return [26]int{}
		}

		// Use memoized value if possible
		k := memKey{pair, depth}
		if v, exists := mem[k]; exists {
			return v
		}

		// Otherwise, count occurrences of characters for this level and all levels
		// below the current.
		v := pairs[pair]
		leftRes := countBetween(depth+1, [2]byte{pair[0], v})
		leftRes[v-'A']++
		rightRes := countBetween(depth+1, [2]byte{v, pair[1]})

		// Memoize result and return
		mem[k] = *merge(&leftRes, &rightRes)
		return mem[k]
	}

	// Count occurrences of characters between pairs as a result of extensions
	res := [26]int{}
	for i := 0; i < len(firstRow)-1; i++ {
		pairRes := countBetween(0, [2]byte{firstRow[i], firstRow[i+1]})
		res = *merge(&res, &pairRes)
	}

	// Don't forget to add the contents of the first row!
	for i := range firstRow {
		res[firstRow[i]-'A']++
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
