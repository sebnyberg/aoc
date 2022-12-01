package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	rs := in.xs
	type gate struct {
		name   string
		typ    string
		indeg  int
		inputs [2]uint16
		out    string
	}
	type wire struct {
		value   uint16
		gates   []*gate
		indices []int
	}

	wires := map[string]*wire{}
	getWire := func(s string) *wire {
		if _, exists := wires[s]; !exists {
			wires[s] = &wire{}
		}
		return wires[s]
	}
	todos := []string{}
	for _, r := range rs {
		switch r.op {
		case opNot:
			a := r.a.(string)
			if ax.IsInt(a) {
				w := getWire(r.out)
				w.value = ^ax.Atou16(a)
				todos = append(todos, r.out)
				continue
			}
			g := &gate{
				name:  r.s,
				typ:   opNot,
				indeg: 1,
				out:   r.out,
			}
			w := getWire(a)
			w.gates = append(w.gates, g)
			w.indices = append(w.indices, 0)
		case opLShift, opRShift, opAnd, opOr:
			g := &gate{
				name:  r.s,
				typ:   r.op,
				out:   r.out,
				indeg: 2,
			}
			a := r.a.(string)
			if ax.IsInt(a) {
				g.indeg--
				g.inputs[0] = ax.Atou16(a)
			} else {
				w := getWire(a)
				w.gates = append(w.gates, g)
				w.indices = append(w.indices, 0)
			}
			b := r.b.(string)
			if ax.IsInt(b) {
				g.indeg--
				g.inputs[1] = ax.Atou16(b)
			} else {
				w := getWire(b)
				w.gates = append(w.gates, g)
				w.indices = append(w.indices, 1)
			}
			if ax.IsInt(a) && ax.IsInt(b) {
				panic("hehe")
			}
		case opAssign:
			a := r.a.(string)
			if ax.IsInt(a) {
				w := getWire(r.out)
				w.value = ax.Atou16(a)
				todos = append(todos, r.out)
			} else {
				g := &gate{
					name:  r.s,
					typ:   r.op,
					indeg: 1,
					out:   r.out,
				}
				w := getWire(a)
				w.gates = append(w.gates, g)
				w.indices = append(w.indices, 0)
			}
		}
	}

	next := []string{}

	for len(todos) > 0 {
		next = next[:0]
		for _, x := range todos {
			w := wires[x]
			// Visit each gate
			for i, g := range w.gates {
				g.indeg--
				j := w.indices[i]
				g.inputs[j] = w.value
				if g.indeg > 0 {
					continue
				}
				var res uint16
				switch g.typ {
				case opNot:
					res = ^w.value
				case opAnd:
					res = g.inputs[0] & g.inputs[1]
				case opOr:
					res = g.inputs[0] | g.inputs[1]
				case opLShift:
					res = g.inputs[0] << g.inputs[1]
				case opRShift:
					res = g.inputs[0] >> g.inputs[1]
				case opAssign:
					res = g.inputs[0]
				}
				w := getWire(g.out)
				w.value = res
				next = append(next, g.out)
			}
		}
		todos, next = next, todos
	}

	return fmt.Sprint(wires["a"].value)
}

func solve2(in *input) string {
	xs2 := make([]inputItem, len(in.xs))
	copy(xs2, in.xs)
	in.xs = xs2
	for i := range in.xs {
		if in.xs[i].op == opAssign && in.xs[i].out == "b" {
			in.xs[i].a = solve1(in)
		}
	}
	res := solve1(in)
	return res
}

type inputItem struct {
	s   string
	op  string
	a   any
	b   any
	out string
}

type input struct {
	n  int
	xs []inputItem
}

var notr = regexp.MustCompile(`NOT (\w+) -> (\w+)`)
var bir = regexp.MustCompile(`(\w+) (\w+) (\w+) -> (\w+)`)
var assignr = regexp.MustCompile(`(\w+) -> (\w+)`)

const opNot = "NOT"
const opLShift = "LSHIFT"
const opRShift = "RSHIFT"
const opAnd = "AND"
const opOr = "OR"
const opAssign = "ASSIGN"

func (p *input) parse(s string) {
	var x inputItem

	switch {
	case notr.MatchString(s):
		ss := notr.FindStringSubmatch(s)
		x.op = opNot
		x.a = ss[1]
		x.out = ss[2]
		x.s = fmt.Sprintf("NOT %v", x.a)
	case bir.MatchString(s):
		ss := bir.FindStringSubmatch(s)
		x.op = ss[2]
		x.a = any(ss[1])
		x.b = any(ss[3])
		x.out = ss[4]
		x.s = fmt.Sprintf("%v %v %v", x.a, x.op, x.b)
	case assignr.MatchString(s):
		ss := assignr.FindStringSubmatch(s)
		x.op = opAssign
		x.a = any(ss[1])
		x.out = ss[2]
	default:
		panic(s)
	}

	x.s = s
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	rows := ax.ReadLines(os.Stdin)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
