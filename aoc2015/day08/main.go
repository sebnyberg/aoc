package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	var code int
	var n int
	for i := range in.xs {
		code += len(in.xs[i].s)
		n += len(in.xs[i].inmem)
	}
	return fmt.Sprint(code - n)
}

func solve2(in *input) string {
	var code int
	var n int
	for i := range in.xs {
		code += len(in.xs[i].s)
		n += len(in.xs[i].encoded) + 2
	}
	return fmt.Sprint(n - code)
}

type inputItem struct {
	s       string
	inmem   string
	encoded string
}

type input struct {
	n  int
	xs []inputItem
}

var pat = regexp.MustCompile(``)

func (p *input) parse(s string) {
	var x inputItem
	s = s[1 : len(s)-1]
	var i int
	var res []byte
	for i < len(s) {
		if s[i] != '\\' {
			res = append(res, s[i])
			i++
			continue
		}
		if s[i+1] == 'x' {
			chs, err := hex.DecodeString(s[i+2 : i+4])
			ax.Check(err, "invalid hex "+s[i:i+4])
			ch := chs[0]
			res = append(res, ch)
			i += 4
			continue
		}
		res = append(res, s[i+1])
		i += 2
	}
	x.inmem = string(res)
	// Encode by adding \ in front of any " or \
	res = []byte{}
	i = 0
	for i := range x.s {
		if x.s[i] == '"' || x.s[i] == '\\' {
			res = append(res, '\\')
		}
		res = append(res, x.s[i])
		i++
	}
	x.encoded = string(res)

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
