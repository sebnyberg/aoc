package day10part2

import (
	"sort"
	"strings"
)

const (
	Problem = 10
	Part    = 2
)

func Run(rows []string) int {
	scoreRow := func(row string) int {
		// Corrupted lines close a param that never opened
		want := make([]byte, 0)
		n := 0
		for _, ch := range row {
			if strings.ContainsRune("[({<", ch) {
				if ch == '(' {
					want = append(want, ')')
				} else {
					want = append(want, byte(ch+2))
				}
				n++
				continue
			}
			if byte(ch) != want[n-1] {
				return 0
			}
			want = want[:n-1] // pop
			n--
		}
		var score int
		for i := len(want) - 1; i >= 0; i-- {
			score *= 5
			switch want[i] {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}
		return score
	}

	scores := make([]int, 0, len(rows))
	for _, row := range rows {
		if score := scoreRow(row); score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
