package main

import (
	"fmt"
	"regexp"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) string {
	wantRow := 3010 - 1
	wantCol := 3019 - 1

	var i, j int
	code := 20151125
	mul := 252533
	mod := 33554393
	nextI := 1
	for i != wantRow || j != wantCol {
		code = (code * mul) % mod
		if i == 0 {
			i = nextI
			nextI++
			j = 0
		} else {
			i--
			j++
		}
	}

	return fmt.Sprint(code)
}

func solve2(inf string) string {
	var res int
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
