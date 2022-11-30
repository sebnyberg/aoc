package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

func Solve1(input []any) any {
	return nil
}

func Solve2(input []any) any {
	return nil
}

func Parse(s string) any {
	return s
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var p ax.Problem
	p.HeadN = 3
	p.TailN = 3
	for sc.Scan() {
		s := sc.Text()
		p.Input = append(p.Input, s)
		p.Parsed = append(p.Parsed, Parse(s))
	}
	p.Result1 = Solve1(p.Parsed)
	p.Result2 = Solve2(p.Parsed)
	fmt.Println(p)
}
