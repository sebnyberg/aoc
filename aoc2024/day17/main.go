package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

const (
	opAdv = 0
	opBxl = 1
	opBst = 2
	opJnz = 3
	opBxc = 4
	opOut = 5
	opBdv = 6
	opCdv = 7
)

const (
	regA = 0
	regB = 1
	regC = 2
)

type computer struct {
	program []int
	reg     [3]int
	out     []int
	ip      int
}

func (c *computer) loop() int {
	c.ip = c.step(c.ip)
	for c.ip > 0 && c.ip < len(c.program) {
		c.ip = c.step(c.ip)
	}
	return c.out[len(c.out)-1]
}

func (c *computer) step(ip int) int {
	op := c.program[c.ip]
	literal := c.program[c.ip+1]
	combo := literal
	if literal >= 4 {
		combo = c.reg[literal-4] // combo
	}
	if op == opJnz && c.reg[regA] != 0 {
		return literal // return to start of program
	}
	switch op {
	case opAdv:
		c.reg[regA] = c.reg[regA] >> combo
	case opBxl:
		c.reg[regB] ^= literal
	case opBst:
		c.reg[regB] = combo % 8
	case opBxc:
		c.reg[regB] ^= c.reg[regC]
	case opOut:
		c.out = append(c.out, combo%8)
	case opBdv:
		c.reg[regB] = c.reg[regA] >> combo
	case opCdv:
		c.reg[regC] = c.reg[regA] >> combo
	}
	return ip + 2
}

func parse(inf string) computer {
	lines := ax.MustReadFileLines(inf)
	var c computer
	for i := 0; i < 3; i++ {
		ss := strings.Fields(lines[i])
		c.reg[i] = ax.Atoi(ss[2])
	}
	for _, s := range strings.Split(strings.Fields(lines[4])[1], ",") {
		c.program = append(c.program, ax.Atoi(s))
	}
	return c
}

func (c *computer) run() []int {
	for c.ip != len(c.program) {
		c.loop()
	}
	return c.out
}

func solve1(inf string) any {
	c := parse(inf)
	res := []string{}
	for c.ip != len(c.program) {
		res = append(res, strconv.Itoa(c.loop()))
	}
	return res
}

func solve2(inf string) any {
	comp := parse(inf)
	curr := []int{0}
	next := []int{}

	binstr := func(x int) string {
		return fmt.Sprintf("%03b", x)
	}
	_ = binstr
	reset := func() {
		comp.ip = 0
		comp.out = comp.out[:0]
		comp.reg = [3]int{0, 0, 0}
	}

	for i := 0; i < len(comp.program); i++ {
		fmt.Println("iteration", i)
		if len(curr) == 0 {
			panic("no candidates left")
		}
		next = next[:0]
		for _, cand := range curr {
			for a := 0; a < 8; a++ {
				x := cand<<3 + a

				// Verify
				reset()
				comp.reg[0] = x
				comp.run()
				if !slices.Equal(comp.out, comp.program[len(comp.program)-len(comp.out):]) {
					continue
				}

				// append
				next = append(next, x)
			}
		}

		curr, next = next, curr
	}
	res := math.MaxInt64
	for _, x := range curr {
		res = min(res, x)
	}
	return res
}

func main() {
	// fmt.Printf("Result1:\n%v\n", solve1("input"))
	fmt.Printf("Result2:\n%v\n", solve2("input"))
}
