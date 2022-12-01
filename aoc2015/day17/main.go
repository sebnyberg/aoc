package main

import (
	"fmt"
	"math"
	"os"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	rs := in.xs
	m := len(rs)
	mem := make(map[[2]int]int)
	var dfs func(i, rem int) int
	dfs = func(i, rem int) int {
		if rem == 0 {
			return 1
		}
		if rem < 0 {
			return 0
		}
		if i == m {
			return 0
		}
		k := [2]int{i, rem}
		if v, exists := mem[k]; exists {
			return v
		}
		res := dfs(i+1, rem) + dfs(i+1, rem-rs[i].x)
		mem[k] = res
		return res
	}
	res := dfs(0, 150)
	return fmt.Sprint(res)
}

func ssolve2(in *input) string {
	rs := in.xs
	m := len(rs)
	minContainers := math.MaxInt32
	minContainerCount := 0
	var dfs func(i, rem, n int) int
	dfs = func(i, rem, n int) int {
		if rem == 0 {
			if n < minContainers {
				minContainers = n
				minContainerCount = 1
			} else if n == minContainers {
				minContainerCount++
			}
			return 1
		}
		if rem < 0 {
			return 0
		}
		if i == m {
			return 0
		}
		res := dfs(i+1, rem, n) + dfs(i+1, rem-rs[i].x, n+1)
		return res
	}
	dfs(0, 150, 0)
	fmt.Println(minContainerCount)
	return fmt.Sprint(minContainerCount)
}

type inputItem struct {
	x int
}

type input struct {
	n  int
	xs []inputItem
}

func (p *input) parse(s string) {
	var x inputItem
	x.x = ax.Atoi(s)
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
