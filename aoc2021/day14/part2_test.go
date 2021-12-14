package day13

import (
	"aoc/ax"
	"math"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay13Part2(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part2(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay13Part2(t *testing.T) {
	assert.Equal(t, "2188189693529", Part2(ax.MustReadFineLines("small")))
	assert.Equal(t, "3353146900153", Part2(ax.MustReadFineLines("input")))
}

func Part2(rows []string) string {
	firstRow := rows[0]

	rows = rows[2:]
	pat := regexp.MustCompile(`^(\w{2}) -> (\w)$`)
	pairs := make(map[[2]byte]byte, len(rows))
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		pair := parts[1]
		insert := parts[2]
		pairs[[2]byte{pair[0], pair[1]}] = insert[0]
	}

	const maxDepth = 40

	type memKey struct {
		pair  [2]byte
		depth int
	}
	merge := func(a, b *[26]int) *[26]int {
		for i, count := range b {
			a[i] += count
		}
		return a
	}
	mem := make(map[memKey][26]int)

	var visit func(int, [2]byte) [26]int
	visit = func(depth int, pair [2]byte) [26]int {
		if depth == maxDepth {
			return [26]int{}
		}
		k := memKey{pair, depth}
		if v, exists := mem[k]; exists {
			return v
		}
		v := pairs[pair]
		res := [26]int{}
		leftRes := visit(depth+1, [2]byte{pair[0], v})
		res = *merge(&res, &leftRes)
		rightRes := visit(depth+1, [2]byte{v, pair[1]})
		res = *merge(&res, &rightRes)
		res[v-'A']++
		mem[k] = res
		return mem[k]
	}

	res := [26]int{}
	for i := 0; i < len(firstRow)-1; i++ {
		pairRes := visit(0, [2]byte{firstRow[i], firstRow[i+1]})
		res = *merge(&res, &pairRes)
	}
	for i := range firstRow {
		res[firstRow[i]-'A']++
	}

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

	return strconv.Itoa(maxCount - minCount)
}
