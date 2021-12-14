package day14

import (
	"aoc/ax"
	"math"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day14part1 int

func BenchmarkDay14Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day14part1 = Part1(ax.MustReadFineLines("input"))
	}
}

func TestDay14Part1(t *testing.T) {
	assert.Equal(t, 1588, Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, 807, Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) int {
	type listNode struct {
		val  byte
		next *listNode
	}

	// Parse first row as a list
	root := &listNode{}
	cur := root
	for i := range rows[0] {
		cur.next = &listNode{val: rows[0][i]}
		cur = cur.next
	}

	// Parse pairs
	rows = rows[2:]
	pat := regexp.MustCompile(`^(\w{2}) -> (\w)$`)
	pairs := make(map[[2]byte]byte, len(rows))
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		pair := parts[1]
		ext := parts[2]
		pairs[[2]byte{pair[0], pair[1]}] = ext[0]
	}

	// Insert extensions into list
	for i := 0; i < 10; i++ {
		cur := root.next
		for cur.next != nil {
			k := [2]byte{cur.val, cur.next.val}
			if v, exists := pairs[k]; exists {
				n := &listNode{val: v, next: cur.next}
				cur.next = n
				cur = n.next
				continue
			}
			panic("no matching pair")
		}
	}

	// Count results
	var count [26]int
	cur = root.next
	for cur != nil {
		count[cur.val-'A']++
		cur = cur.next
	}

	// Find min/max count
	var maxCount int
	minCount := math.MaxInt32
	for _, cnt := range count {
		if cnt > maxCount {
			maxCount = cnt
		}
		if cnt > 0 && cnt < minCount {
			minCount = cnt
		}
	}

	return maxCount - minCount
}
