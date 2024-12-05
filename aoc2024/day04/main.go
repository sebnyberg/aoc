package main

import (
	"fmt"
	"log"

	"github.com/sebnyberg/aoc/ax"
)

func reverse(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%v, %v\n", msg, err)
	}
}

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])

	var directions = [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1},
	}
	directionMatch := func(i, j, di, dj int) int {
		for _, ch := range "XMAS" {
			if i < 0 || j < 0 || i >= m || j >= n || // bounds check
				lines[i][j] != byte(ch) {
				return 0
			}
			i += di
			j += dj
		}
		return 1
	}

	// find "XMAS"
	var result int
	for i := range lines {
		for j := range lines[i] {
			for _, dir := range directions {
				result += directionMatch(i, j, dir[0], dir[1])
			}
		}
	}
	return result
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)

	patterns := [][]string{
		{"M.M", ".A.", "S.S"},
		{"M.S", ".A.", "M.S"},
		{"S.M", ".A.", "S.M"},
		{"S.S", ".A.", "M.M"},
	}

	blockMatch := func(i, j int, pattern []string) int {
		if i+2 >= len(lines) || j+2 >= len(lines[0]) { // bounds check
			return 0
		}
		for di := 0; di < 3; di++ {
			for dj := 0; dj < 3; dj++ {
				want := pattern[di][dj]
				if want == '.' {
					continue
				}
				if lines[i+di][j+dj] != want {
					return 0
				}
			}
		}
		return 1
	}

	var result int
	for i := range lines {
		for j := range lines[i] {
			for _, pattern := range patterns {
				result += blockMatch(i, j, pattern)
			}
		}
	}
	return result
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
