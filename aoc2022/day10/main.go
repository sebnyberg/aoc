package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve(inf string) any {
	lines := ax.MustReadFileLines(inf)
	screen := make([][]byte, 6)
	for i := range screen {
		screen[i] = make([]byte, 40)
	}
	var pc int
	var signalStrength int
	x := 1
	cycle := func(v int) {
		for _, signalCycle := range []int{20, 60, 100, 140, 180, 220} {
			if pc+1 == signalCycle {
				signalStrength += x * signalCycle
			}
		}
		if ax.Abs(x-(pc%40)) <= 1 {
			screen[pc/40][pc%40] = '#'
		} else {
			screen[pc/40][pc%40] = '.'
		}
		x += v
		pc++
	}
	for _, l := range lines {
		fs := strings.Fields(l)
		switch fs[0] {
		case "addx":
			cycle(0)
			cycle(ax.Atoi(fs[1]))
		case "noop":
			cycle(0)
		}
	}
	var render []string
	for i := range screen {
		render = append(render, string(screen[i]))
	}
	return fmt.Sprintf("signal: %v\n%v",
		signalStrength, strings.Join(render, "\n"),
	)
}

func main() {
	f := "input"
	fmt.Printf("Result:\n%v\n", solve(f))
}
