package day21

import (
	"aoc/ax"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day21part2 int

func BenchmarkDay21Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day21part1 = Part2(ax.MustReadFileLines("input"))
	}
}

func TestDay21Part2(t *testing.T) {
	assert.Equal(t, 444356092776315, Part2(ax.MustReadFileLines("small")))
	assert.Equal(t, 146854918035875, Part2(ax.MustReadFileLines("input")))
}

func Part2(rows []string) int {
	pat := regexp.MustCompile(`^Player (\d+) starting position: (\d+)$`)
	parts1 := pat.FindStringSubmatch(rows[0])
	p1 := ax.MustParseInt[int](parts1[2])
	parts2 := pat.FindStringSubmatch(rows[1])
	p2 := ax.MustParseInt[int](parts2[2])

	var dp [21][10][21][10]int
	dp[0][p1-1][0][p2-1] = 1
	var empty [21][10][21][10]int
	var w1, w2 int
	for t := 0; dp != empty; t++ {
		var next [21][10][21][10]int
		first := t%2 == 0
		for pt1 := 0; pt1 < 21; pt1++ {
			for pos1 := 0; pos1 < 10; pos1++ {
				for pt2 := 0; pt2 < 21; pt2++ {
					for pos2 := 0; pos2 < 10; pos2++ {
						if dp[pt1][pos1][pt2][pos2] == 0 {
							continue
						}
						for die1 := 1; die1 <= 3; die1++ {
							for die2 := 1; die2 <= 3; die2++ {
								for die3 := 1; die3 <= 3; die3++ {
									if first {
										nextPos := (pos1 + die1 + die2 + die3) % 10
										nextPt := pt1 + nextPos + 1
										if nextPt >= 21 {
											w1 += dp[pt1][pos1][pt2][pos2]
										} else {
											next[nextPt][nextPos][pt2][pos2] += dp[pt1][pos1][pt2][pos2]
										}
									} else {
										nextPos := (pos2 + die1 + die2 + die3) % 10
										nextPt := pt2 + nextPos + 1
										if nextPt >= 21 {
											w2 += dp[pt1][pos1][pt2][pos2]
										} else {
											next[pt1][pos1][nextPt][nextPos] += dp[pt1][pos1][pt2][pos2]
										}
									}
								}
							}
						}
					}
				}
			}
		}
		dp = next
	}

	return ax.Max(w1, w2)
}
