package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	s := in.xs[0].s
	h := md5.New()
	res := make([]byte, 0, 5)
	for i := 1; i < 1000000; i++ {
		ss := s + fmt.Sprint(i)
		res = res[:0]
		h.Reset()
		h.Write([]byte(ss))
		res = h.Sum(res)
		hh := hex.EncodeToString(res)
		if hh[:5] == "00000" {
			return fmt.Sprint(i)
		}
	}
	panic("neverending loop")
}

func solve2(in *input) string {
	s := in.xs[0].s
	h := md5.New()
	var res []byte
	for i := 1; i < 10000000; i++ {
		ss := s + fmt.Sprint(i)
		res = res[:0]
		h.Reset()
		h.Write([]byte(ss))
		res = h.Sum(res)
		hh := hex.EncodeToString(res)
		if hh[:6] == "000000" {
			return fmt.Sprint(i)
		}
	}
	panic("neverending loop")
}

type inputItem struct {
	s string
}

type input struct {
	n  int
	xs []inputItem
}

func (p *input) parse(s string) {
	var x inputItem
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
