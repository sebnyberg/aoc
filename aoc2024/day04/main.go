package main

import (
	"fmt"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var mulpat = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var dopat = regexp.MustCompile(`do\(\)`)
var dontpat = regexp.MustCompile(`don't\(\)`)

func reverse(s string) string {
	bs := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}
	return string(bs)
}

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])

	// reverse
	rev := make([]string, m)
	for i := range lines {
		rev[i] = reverse(lines[i])
	}

	// downward
	downward := make([]string, n)
	for i := range lines[0] {
		for j := range lines {
			downward[i] += string(lines[j][i])
		}
	}

	getDiag := func(ss []string) []string {
		var res []string
		for j := range ss[0] {
			res = append(res, "")
			mm := len(res) - 1
			for k := 0; j+k < n && k < m; k++ {
				res[mm] += string(ss[k][j+k])
			}
		}
		for i := 1; i < m; i++ {
			res = append(res, "")
			mm := len(res) - 1
			for k := 0; i+k < m && k < n; k++ {
				res[mm] += string(ss[i+k][k])
			}
		}
		return res
	}

	countMatches := func(ss []string, word string) int {
		var count int
		for i := range ss {
			for j := range ss[i] {
				if len(ss[i][j:]) < len(word) {
					break
				}
				if ss[i][j:j+4] == word {
					count++
				}
			}
		}
		return count
	}

	// diagonals
	diagfwd := getDiag(lines)
	diagBack := getDiag(rev)

	var res int
	res += countMatches(lines, "XMAS")
	res += countMatches(lines, "SAMX")
	res += countMatches(downward, "XMAS")
	res += countMatches(downward, "SAMX")
	res += countMatches(diagfwd, "XMAS")
	res += countMatches(diagfwd, "SAMX")
	res += countMatches(diagBack, "XMAS")
	res += countMatches(diagBack, "SAMX")
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := len(lines)
	n := len(lines[0])

	// reverse
	rev := make([]string, m)
	for i := range lines {
		rev[i] = reverse(lines[i])
	}
	// y-flipped
	upward := make([]string, m)
	for j := range lines[0] {
		for i := range lines {
			upward[j] += string(lines[i][m-1-j])
		}
	}
	// y-flipped and reverse
	upwardRev := make([]string, m)
	for i := range upward {
		upwardRev[i] = reverse(upward[i])
	}

	match := func(ss []string) int {
		var count int
		for i := 0; i < m-2; i++ {
			for j := 0; j < n-2; j++ {
				if ss[i][j] == 'M' &&
					ss[i][j+2] == 'S' &&
					ss[i+1][j+1] == 'A' &&
					ss[i+2][j] == 'M' &&
					ss[i+2][j+2] == 'S' {
					count++
				}
			}
		}
		return count
	}
	var res int
	res += match(lines)
	res += match(rev)
	res += match(upward)
	res += match(upwardRev)
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
