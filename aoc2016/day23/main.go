package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func isnum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func solve(inf string, initA int) string {
	lines := ax.MustReadFileLines(inf)
	var pc int
	var reg [4]int
	val := func(s string) int {
		if isnum(s) {
			return ax.Atoi(s)
		}
		return reg[s[0]-'a']
	}
	const (
		a, b, c, d = 0, 1, 2, 3
	)
	reg[a] = initA
	toggle := func(s string) string {
		fs := strings.Fields(s)
		switch fs[0] {
		case "inc":
			fs[0] = "dec"
			return strings.Join(fs, " ")
		case "dec":
			fs[0] = "inc"
			return strings.Join(fs, " ")
		case "tgl":
			fs[0] = "inc"
			return strings.Join(fs, " ")
		case "jnz":
			fs[0] = "cpy"
			return strings.Join(fs, " ")
		case "cpy":
			fs[0] = "jnz"
			return strings.Join(fs, " ")
		default:
			panic(fs)
		}
	}
	dodebug := false
	p := func(instr string) {
		if dodebug {
			fmt.Printf("%v: [%v,%v,%v,%v]: %s\n",
				pc, reg[a], reg[b], reg[c], reg[d], instr)
		}
	}
	n := len(lines)
	hash := func() string {
		return strings.Join(lines[:pc+1], "")
	}
	prevHash := make([]string, n)
	prevReg := make([][4]int, n)
	for pc < len(lines) {
		fs := strings.Fields(lines[pc])
		switch fs[0] {
		case "cpy":
			p(lines[pc])
			if !isnum(fs[2]) {
				reg[fs[2][0]-'a'] = val(fs[1])
			}
			pc++
		case "dec":
			p(lines[pc])
			if !isnum(fs[1]) {
				reg[fs[1][0]-'a']--
			}
			pc++
		case "inc":
			p(lines[pc])
			if !isnum(fs[1]) {
				reg[fs[1][0]-'a']++
			}
			pc++
		case "jnz":
			p(lines[pc])
			h := hash()
			if isnum(fs[2]) && prevHash[pc] == h {
				// Every time we reach this instruction (with the current
				// program up to this point), we end up jumping backwards.
				// Also, the jump length is static (not affected by iterating).
				//
				// Instead of jumping backwards, we may count the difference in
				// register values during each iteration. Then we may
				// fast-forward to where the value has become zero.
				var deltas [4]int
				for i := range reg {
					deltas[i] = reg[i] - prevReg[pc][i]
				}
				r := int(fs[1][0] - 'a')
				iters := ax.Abs(reg[r] / deltas[r])
				for i := range reg {
					reg[i] += iters * deltas[i]
				}
			}
			if val(fs[1]) != 0 {
				prevHash[pc] = h
				prevReg[pc] = reg
				pc += val(fs[2])
			} else {
				prevHash[pc] = ""
				pc++
			}
		case "tgl":
			p(lines[pc])
			target := pc + val(fs[1])
			if target < n {
				lines[target] = toggle(lines[target])
			}
			pc++
		default:
			panic(fs)
		}
	}
	return fmt.Sprint(reg[a])
}

func solve2(inf string) string {
	var res int
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, 7))
	fmt.Printf("Result2:\n%v\n\n", solve(f, 12))
}
