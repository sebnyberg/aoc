package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

var vals = [3][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

func solve1(inf string) string {
	lines := ax.MustReadFileLines(inf)
	var code []byte
	i := 1
	j := 1
	for _, l := range lines {
		for _, ch := range l {
			switch ch {
			case 'R':
				j = ax.Min(j+1, 2)
			case 'L':
				j = ax.Max(j-1, 0)
			case 'D':
				i = ax.Min(i+1, 2)
			case 'U':
				i = ax.Max(i-1, 0)
			}
		}
		code = append(code, byte(vals[i][j]+'0'))
	}
	return fmt.Sprint(string(code))
}

var vals2 = [5][5]byte{
	{'0', '0', '1', '0', '0'},
	{'0', '2', '3', '4', '0'},
	{'5', '6', '7', '8', '9'},
	{'0', 'A', 'B', 'C', '0'},
	{'0', '0', 'D', '0', '0'},
}

func solve2(inf string) string {
	lines := ax.MustReadFileLines(inf)
	var code []byte
	i := 1
	j := 1
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < 5 && j < 5 && vals2[i][j] != '0'
	}
	for _, l := range lines {
		for _, ch := range l {
			ii := i
			jj := j
			switch ch {
			case 'R':
				jj++
			case 'L':
				jj--
			case 'D':
				ii++
			case 'U':
				ii--
			}
			if !ok(ii, jj) {
				continue
			}
			i = ii
			j = jj
		}
		code = append(code, vals2[i][j])
	}
	return fmt.Sprint(string(code))
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
