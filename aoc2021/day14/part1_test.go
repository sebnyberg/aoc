package day13

import (
	"aoc/ax"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkDay13Part1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestDay13Part1(t *testing.T) {
	assert.Equal(t, "1588", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "807", Part1(ax.MustReadFineLines("input")))
}

func Part1(rows []string) string {
	type listNode struct {
		val  byte
		next *listNode
	}

	root := &listNode{}
	cur := root
	for i := range rows[0] {
		cur.next = &listNode{val: rows[0][i]}
		cur = cur.next
	}

	rows = rows[2:]
	pat := regexp.MustCompile(`^(\w{2}) -> (\w)$`)
	pairs := make(map[[2]byte]byte, len(rows))
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		pair := parts[1]
		insert := parts[2]
		pairs[[2]byte{pair[0], pair[1]}] = insert[0]
	}

	// print := func(root *listNode) {
	// 	cur := root.next
	// 	for cur != nil {
	// 		fmt.Print(string(cur.val))
	// 		cur = cur.next
	// 	}
	// 	fmt.Print("\n")
	// }

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
			panic("no matching pair!!")
		}
	}
	var count [26]int
	cur = root.next
	for cur != nil {
		count[cur.val-'A']++
		cur = cur.next
	}

	var maxCount int
	var maxChar byte = 'A'
	minCount := math.MaxInt32
	var minChar byte = 'A'
	for i, cnt := range count {
		if cnt > maxCount {
			maxChar = byte(i + 'A')
			maxCount = cnt
		}
		if cnt > 0 && cnt < minCount {
			minChar = byte(i + 'A')
			minCount = cnt
		}
	}
	fmt.Println(string(minChar), minCount)
	fmt.Println(string(maxChar), maxCount)

	return strconv.Itoa(maxCount - minCount)
}
