package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/sebnyberg/aoc/ax"
)

func parse(lines []string) [][]byte {
	stacks := make([][]byte, 9)
	for i := 0; i < 8; i++ {
		for k := 1; k < len(lines[i]); k += 4 {
			j := (k - 1) / 4
			if unicode.IsLetter(rune(lines[i][k])) {
				stacks[j] = append(stacks[j], lines[i][k])
			}
		}
	}
	for i := range stacks {
		for l, r := 0, len(stacks[i])-1; l < r; l, r = l+1, r-1 {
			stacks[i][l], stacks[i][r] = stacks[i][r], stacks[i][l]
		}
	}
	return stacks
}

func solve1(inf string) string {
	lines := ax.MustReadFileLines(inf)
	stacks := parse(lines)
	for _, s := range lines[10:] {
		fs := strings.Fields(s)
		n := ax.Atoi(fs[1])
		from := ax.Atoi(fs[3]) - 1
		to := ax.Atoi(fs[5]) - 1
		m := len(stacks[from])
		for i := 0; i < n; i++ {
			stacks[to] = append(stacks[to], stacks[from][m-1])
			stacks[from] = stacks[from][:m-1]
			m--
		}
	}
	var res []byte
	for i := range stacks {
		res = append(res, stacks[i][len(stacks[i])-1])
	}
	return fmt.Sprint(string(res))
}

func solve2(inf string) string {
	lines := ax.MustReadFileLines(inf)
	stacks := parse(lines)
	for _, s := range lines[10:] {
		fs := strings.Fields(s)
		n := ax.Atoi(fs[1])
		from := ax.Atoi(fs[3]) - 1
		to := ax.Atoi(fs[5]) - 1
		m := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][m-n:]...)
		stacks[from] = stacks[from][:m-n]
	}
	var res []byte
	for i := range stacks {
		res = append(res, stacks[i][len(stacks[i])-1])
	}
	return fmt.Sprint(string(res))
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
