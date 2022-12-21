package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

type equation struct {
	monkey string
	a      string
	b      string
	op     string
	v      int
}

func parse(inf string) map[string]equation {
	lines := ax.MustReadFileLines(inf)
	eqs := make(map[string]equation)

	for _, l := range lines {
		fs := strings.Fields(l)
		monkey := fs[0][:len(fs[0])-1]
		if len(fs) == 2 {
			eqs[monkey] = equation{
				monkey: monkey,
				op:     "=",
				v:      ax.Atoi(fs[1]),
			}
			continue
		}
		m1 := fs[1]
		m2 := fs[3]
		eqs[monkey] = equation{
			monkey: monkey,
			a:      m1,
			b:      m2,
			op:     fs[2],
		}
	}
	return eqs
}

func resolve(eqs map[string]equation, x, xName string) (int, bool) {
	e := eqs[x]
	if x == xName {
		// cannot resolve this variable
		return 0, false
	}
	if e.op == "=" {
		return e.v, true
	}
	x1, ok1 := resolve(eqs, e.a, xName)
	if !ok1 {
		return 0, false
	}
	x2, ok2 := resolve(eqs, e.b, xName)
	if !ok2 {
		return 0, false
	}
	switch e.op {
	case "+":
		return x1 + x2, true
	case "-":
		return x1 - x2, true
	case "*":
		return x1 * x2, true
	case "/":
		return x1 / x2, true
	}
	panic("unresolvable variable")
}

func part2(eqs map[string]equation) int {
	r := eqs["root"]
	curr := r
	var y int
	a, ok1 := resolve(eqs, curr.a, "humn")
	b, _ := resolve(eqs, curr.b, "humn")
	if !ok1 {
		y = b + b
	} else {
		y = a + a
	}
	for curr.monkey != "humn" {
		a, ok1 := resolve(eqs, curr.a, "humn")
		b, ok2 := resolve(eqs, curr.b, "humn")
		if ok1 == ok2 {
			log.Fatalln("wut?", curr)
		}
		if !ok1 {
			switch curr.op {
			case "+":
				y = y - b
			case "-":
				y = y + b
			case "/":
				y = y * b
			case "*":
				y = y / b
			}
			curr = eqs[curr.a]
		} else {
			switch curr.op {
			case "+":
				y = y - a
			case "-":
				y = a - y
			case "/":
				y = a / y
			case "*":
				y = y / a
			}
			curr = eqs[curr.b]
		}
	}
	return y
}

func main() {
	eqs := parse("input")
	p1, _ := resolve(eqs, "root", "")
	fmt.Printf("Part1:\n%v\n", p1)
	fmt.Printf("Part1:\n%v\n", part2(eqs))
}
