package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := make(map[string]bool)
	var maxLen int
	for _, pat := range strings.Split(lines[0], ",") {
		pat = strings.TrimSpace(pat)
		m[pat] = true
		maxLen = max(maxLen, len(pat))
	}
	lines = lines[2:]
	var res int
	for _, design := range lines {
		n := len(design)
		dp := make([]int, n+1)
		for i := range dp {
			dp[i] = math.MaxInt32
		}
		dp[0] = 1
		for j := 1; j < len(dp); j++ {
			for k := 1; k <= j && k <= maxLen; k++ {
				if m[design[j-k:j]] && dp[j-k] != math.MaxInt32 {
					dp[j] += dp[j-k]
				}
			}
		}
		if dp[n] != math.MaxInt32 {
			res += 1
		}
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	m := make(map[string]bool)
	var maxLen int
	for _, pat := range strings.Split(lines[0], ",") {
		pat = strings.TrimSpace(pat)
		if _, exists := m[pat]; exists {
			panic("double pattern!")
		}
		m[pat] = true
		maxLen = max(maxLen, len(pat))
	}
	lines = lines[2:]
	res := big.NewInt(0)
	for _, design := range lines {
		n := len(design)
		dp := make([]*big.Int, n+1)
		for i := range dp {
			dp[i] = big.NewInt(0)
		}
		dp[0] = big.NewInt(1)
		for j := 1; j < len(dp); j++ {
			for k := 1; k <= j && k <= maxLen; k++ {
				if m[design[j-k:j]] {
					dp[j].Add(dp[j], dp[j-k])
				}
			}
		}
		res.Add(res, dp[n])
	}
	return res
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve1("input"))
	fmt.Printf("Result2:\n%v\n", solve2("input"))
}
