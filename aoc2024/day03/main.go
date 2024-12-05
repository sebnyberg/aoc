package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/sebnyberg/aoc/ax"
)

var mulpat = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func solve1(inf string) any {
	f, _ := os.Open(inf)
	defer f.Close()
	b, _ := io.ReadAll(f)
	program := string(b)
	var res int
	for _, s := range mulpat.FindAllStringSubmatch(program, -1) {
		a := ax.MustParseInt[int](s[1])
		b := ax.MustParseInt[int](s[2])
		res += a * b
	}
	return res
}

var pat = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

func solve2(inf string) any {
	f, _ := os.Open(inf)
	defer f.Close()
	b, _ := io.ReadAll(f)
	program := string(b)
	pat, _ = regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := pat.FindAllStringSubmatch(program, -1)
	enabled := true
	var res int
	for _, parts := range matches {
		switch parts[0][:3] {
		case "do(":
			enabled = true
		case "don":
			enabled = false
		default:
			if !enabled {
				continue
			}
			a, _ := strconv.Atoi(parts[1])
			b, _ := strconv.Atoi(parts[2])
			res += a * b
		}
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
