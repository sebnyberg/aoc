package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	lines := ax.MustReadFileLines(inf)
	n := 50
	m := 6
	state := make([][]byte, m)
	for i := range state {
		state[i] = make([]byte, n)
	}
	for i := range state {
		for j := range state[i] {
			state[i][j] = '.'
		}
	}
	for _, l := range lines {
		fields := strings.Fields(l)

		// Rectangle
		if len(fields) == 2 {
			var w, h int
			fmt.Sscanf(fields[1], "%dx%d", &w, &h)
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					state[i][j] = '#'
				}
			}
			continue
		}

		amt := ax.Atoi(fields[4])
		target := ax.Atoi(strings.Split(fields[2], "=")[1])
		if fields[1] == "row" {
			next := make([]byte, n)
			for i := range next {
				next[(i+amt)%n] = state[target][i]
			}
			state[target] = next
		} else {
			next := make([]byte, m)
			for i := range state {
				next[(i+amt)%m] = state[i][target]
			}
			for i := range state {
				state[i][target] = next[i]
			}
		}
	}
	var res int
	for i := 0; i < 6; i++ {
		for j, v := range state[i] {
			if v == '#' {
				res++
			}
			if j%5 == 0 {
				fmt.Print("  ")
			}
			fmt.Print(string(state[i][j]))
		}
		fmt.Print("\n")
	}
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
}
