package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/sebnyberg/aoc/ax"
)

func solve(inf string, cval int) string {
	lines := ax.MustReadFileLines(inf)
	var regs [4]int
	regs[2] = cval
	n := len(lines)
	for pc := 0; pc < n; pc++ {
		fs := strings.Split(lines[pc], " ")
		switch fs[0] {
		case "cpy":
			from := fs[1]
			to := fs[2][0] - 'a'
			if ax.StringIsFunc(from, unicode.IsLetter) {
				regs[to] = regs[from[0]-'a']
			} else {
				regs[to] = ax.Atoi(from)
			}
		case "jnz":
			a := fs[1]
			delta := ax.Atoi(fs[2])
			if ax.StringIsFunc(a, unicode.IsLetter) {
				if regs[a[0]-'a'] != 0 {
					pc += delta - 1
				}
			} else {
				if ax.Atoi(fs[1]) != 0 {
					pc += delta - 1
				}
			}
		case "inc":
			regs[fs[1][0]-'a']++
		case "dec":
			regs[fs[1][0]-'a']--
		}
	}

	return fmt.Sprint(regs[00])
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, 0))
	fmt.Printf("Result2:\n%v\n\n", solve(f, 1))
}
