package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var buttonPat = regexp.MustCompile(`Button [AB]: X\+(\d+), Y\+(\d+)`)
var prizePat = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

type game struct {
	a     [2]int
	b     [2]int
	prize [2]int
}

func parseGames(lines []string) []game {
	// parse games
	var games []game
	for i := 0; i < len(lines); {
		var g game
		a := buttonPat.FindStringSubmatch(lines[i])
		g.a = [2]int{ax.Atoi(a[1]), ax.Atoi(a[2])}
		i++
		b := buttonPat.FindStringSubmatch(lines[i])
		g.b = [2]int{ax.Atoi(b[1]), ax.Atoi(b[2])}
		i++
		p := prizePat.FindStringSubmatch(lines[i])
		g.prize = [2]int{ax.Atoi(p[1]), ax.Atoi(p[2])}
		i += 2
		games = append(games, g)
	}
	return games
}

func findMinCost(minGame game) int {
	// Let's consider one axis only, since the other axis is easily verifiable.
	// For a combination to be valid, there must exist a combination (ka, kb) such
	// that ka*a + kb*b == c
	a := minGame.a
	b := minGame.b
	c := minGame.prize
	nom := a[1]*c[0] - a[0]*c[1]
	denom := a[1]*b[0] - a[0]*b[1]
	if nom%denom != 0 {
		return math.MaxInt32
	}
	y := nom / denom
	if (c[0]-b[0]*y)%a[0] != 0 {
		return math.MaxInt32
	}
	x := (c[0] - b[0]*y) / a[0]
	return x*3 + y
}

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	games := parseGames(lines)
	// The goal is to find k_a*A + k_b*B = prize such that k_a+k_b is minimized.
	// Seeing as the inputs are not very large, we can pick an axis, count the
	// number of As or Bs, then reduce and combine, counting the min count.
	var res int
	for i := range games {
		minCost := findMinCost(games[i])
		if minCost != math.MaxInt32 {
			res += minCost
		}
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	games := parseGames(lines)
	var res int
	for i := range games {
		games[i].prize[0] += 10000000000000
		games[i].prize[1] += 10000000000000
		minCost := findMinCost(games[i])
		if minCost != math.MaxInt32 {
			res += minCost
		}
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
