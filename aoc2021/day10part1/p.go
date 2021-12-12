package day10part1

import (
	"strconv"
	"strings"
)

const (
	Problem = 10
	Part    = 1
)

func Run(rows []string) string {
	scoreRow := func(row string) int {
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
				switch ch {
				case ')':
					return 3
				case ']':
					return 57
				case '}':
					return 1197
				case '>':
					return 25137
				}
			}
			want = want[:n-1] // pop
			n--
		}
		return 0
	}

	var res int
	for _, row := range rows {
		res += scoreRow(row)
	}
	return strconv.Itoa(res)
}
