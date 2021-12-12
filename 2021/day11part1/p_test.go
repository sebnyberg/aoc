package p_test

import (
	"aoc/ax"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"small", 1656},
		{"input", 1562},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines, 100))
		})
	}
}

func run(rows []string, steps int) int {
	type pos struct{ i, j int }
	ok := func(p pos) bool {
		return p.i >= 0 && p.i < 10 && p.j >= 0 && p.j < 10
	}
	var board [10][10]byte
	for i, row := range rows {
		for j := range row {
			board[i][j] = byte(row[j] - '0')
		}
	}
	willFlash := make([]pos, 0, 100)
	var sum int
	for step := 0; step < steps; step++ {
		willFlash = willFlash[:0]

		// Collect 9's
		var seen [10][10]bool
		for i := range board {
			for j := range board[i] {
				if board[i][j] != 9 {
					continue
				}
				willFlash = append(willFlash, pos{i, j})
				board[i][j]++
			}
		}
		// While there are places to visit
		for i := 0; i < len(willFlash); i++ {
			p := willFlash[i]
			seen[p.i][p.j] = true
			for _, q := range []pos{
				{p.i + 1, p.j}, {p.i - 1, p.j}, {p.i, p.j + 1}, {p.i, p.j - 1},
				{p.i - 1, p.j - 1}, {p.i - 1, p.j + 1}, {p.i + 1, p.j - 1}, {p.i + 1, p.j + 1},
			} {
				if !ok(q) {
					continue
				}
				board[q.i][q.j]++
				if board[q.i][q.j] == 9 {
					willFlash = append(willFlash, q)
				}
			}
		}
		sum += len(willFlash)
		for i := range board {
			for j := range board[i] {
				if board[i][j] >= 9 {
					board[i][j] = 0
				} else {
					board[i][j]++
				}
			}
		}
	}
	return sum
}
