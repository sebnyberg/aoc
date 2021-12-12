package aoc2021

import "strconv"

func Day09Part1(rows []string) string {
	m := len(rows)
	n := len(rows[0])
	grid := make([][]int, m)
	for i, row := range rows {
		grid[i] = make([]int, n)
		for j := range row {
			grid[i][j] = int(row[j] - '0')
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	minPoint := func(i, j int) bool {
		for _, p := range [][2]int{{i + 1, j}, {i - 1, j}, {i, j - 1}, {i, j + 1}} {
			ii, jj := p[0], p[1]
			if ok(ii, jj) && grid[i][j] >= grid[ii][jj] {
				return false
			}
		}
		return true
	}

	var sum int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if minPoint(i, j) {
				sum += grid[i][j] + 1
			}
		}
	}
	return strconv.Itoa(sum)
}
