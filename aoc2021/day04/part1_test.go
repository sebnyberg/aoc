package day04

import (
	"aoc/ax"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkPart1(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = Part1(ax.MustReadFineLines("input"))
	}
	_ = res
}

func TestPart1(t *testing.T) {
	assert.Equal(t, "4512", Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, "65325", Part1(ax.MustReadFineLines("input")))
}

func Part1(inputRows []string) string {
	const boardSize = 5

	// Parse nums
	numsLine := strings.Split(inputRows[0], ",")
	nums := make([]int, len(numsLine))
	for i, numStr := range numsLine {
		nums[i] = ax.MustParseInt[int](numStr)
	}

	// Parse boards
	inputRows = inputRows[1:]
	nboards := len(inputRows) / 6
	pos := make([]map[int][2]int, nboards)
	cols := make([][]int, nboards)
	rows := make([][]int, nboards)
	vals := make([][][]int, nboards)
	for i := 0; i < nboards; i++ {
		pos[i] = make(map[int][2]int, boardSize)
		cols[i] = make([]int, boardSize)
		rows[i] = make([]int, boardSize)
		vals[i] = make([][]int, boardSize)
		for row := 0; row < 5; row++ {
			vals[i][row] = make([]int, boardSize)
			lineIdx := i*6 + 1 + row
			for col, valStr := range strings.Fields(inputRows[lineIdx]) {
				val := ax.MustParseInt[int](valStr)
				pos[i][val] = [2]int{row, col}
				vals[i][row][col] = val
			}
		}
	}

	// A win fills the bitmask for the column / row with ones
	win := (1 << boardSize) - 1

	// Add a number and check if it resulted in bingo
	markAndCheck := func(i, num int) bool {
		pos, exists := pos[i][num]
		if !exists {
			return false
		}
		row, col := pos[0], pos[1]
		cols[i][col] |= 1 << row
		rows[i][row] |= 1 << col
		return cols[i][col] == win ||
			rows[i][row] == win
	}

	for _, num := range nums {
		for i := 0; i < nboards; i++ {
			if !markAndCheck(i, num) {
				continue
			}
			// Found winner
			var sum int
			for row := 0; row < boardSize; row++ {
				for col := 0; col < boardSize; col++ {
					if rows[i][row]&(1<<col) == 0 {
						sum += vals[i][row][col]
					}
				}
			}
			return strconv.Itoa(sum * num)
		}
	}
	panic("not found")
}
