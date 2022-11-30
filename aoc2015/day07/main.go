package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var absf = ax.Abs[float64]
var absi = ax.Abs[int]
var minf = ax.Min[float64]
var mini = ax.Min[int]
var minu = ax.Min[uint16]
var maxf = ax.Max[float64]
var maxi = ax.Max[int]
var maxu = ax.Max[uint16]
var print = fmt.Print
var printf = fmt.Printf
var println = fmt.Println
var sprint = fmt.Sprint
var sprintf = fmt.Sprintf
var sprintln = fmt.Sprintln
var tof = ax.MustParseFloat[float64]
var toi = ax.MustParseInt[int]
var tou = ax.MustParseInt[uint16]

func pprint(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmt.Printf(fmtStr, a...)
}
func pprintln(a ...any) {
	fmtStr := "%+v"
	for i := 1; i < len(a); i++ {
		fmtStr += ",%+v"
	}
	fmtStr += "\n"
	fmt.Printf(fmtStr, a...)
}

var intr = regexp.MustCompile(`[1-9][0-9]*|0`)

func isnum(s string) bool {
	return intr.MatchString(s)
}

func Solve1(rs []parsedRow) string {
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
			if isnum(a) {
				w := getWire(r.out)
				w.value = ^tou(a)
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
			if isnum(a) {
				g.indeg--
				g.inputs[0] = tou(a)
			} else {
				w := getWire(a)
				w.gates = append(w.gates, g)
				w.indices = append(w.indices, 0)
			}
			b := r.b.(string)
			if isnum(b) {
				g.indeg--
				g.inputs[1] = tou(b)
			} else {
				w := getWire(b)
				w.gates = append(w.gates, g)
				w.indices = append(w.indices, 1)
			}
			if isnum(a) && isnum(b) {
				panic("hehe")
			}
		case opAssign:
			a := r.a.(string)
			if isnum(a) {
				w := getWire(r.out)
				w.value = tou(a)
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

	return sprint(wires["a"].value)
}

func Solve2(rs []parsedRow) string {
	rs2 := make([]parsedRow, len(rs))
	copy(rs2, rs)
	rs = rs2
	for i := range rs {
		if rs[i].op == opAssign && rs[i].out == "b" {
			rs[i].a = Solve1(rs)
		}
	}
	res := Solve1(rs)
	return res
}

type parsedRow struct {
	s   string
	op  string
	a   any
	b   any
	out string
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

func Parse(s string) parsedRow {
	var r parsedRow
	switch {
	case notr.MatchString(s):
		ss := notr.FindStringSubmatch(s)
		r.op = opNot
		r.a = ss[1]
		r.out = ss[2]
		r.s = sprintf("NOT %v", r.a)
	case bir.MatchString(s):
		ss := bir.FindStringSubmatch(s)
		r.op = ss[2]
		r.a = any(ss[1])
		r.b = any(ss[3])
		r.out = ss[4]
		r.s = sprintf("%v %v %v", r.a, r.op, r.b)
	case assignr.MatchString(s):
		ss := assignr.FindStringSubmatch(s)
		r.op = opAssign
		r.a = any(ss[1])
		r.out = ss[2]
	default:
		panic(s)
	}
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var p ax.Problem[parsedRow]
	p.HeadN = 3
	p.TailN = 3
	for sc.Scan() {
		s := sc.Text()
		p.Input = append(p.Input, s)
		p.Parsed = append(p.Parsed, Parse(s))
	}
	p.Result1 = Solve1(p.Parsed)
	p.Result2 = Solve2(p.Parsed)
	fmt.Fprint(os.Stdout, p)
}
