package main

import (
	"fmt"
	"math"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	input := ax.MustReadFileLines("input")
	// input = []string{
	// 	"eedadn",
	// 	"drvtee",
	// 	"eandsr",
	// 	"raavrd",
	// 	"atevrs",
	// 	"tsrnev",
	// 	"sdttsa",
	// 	"rasrtv",
	// 	"nssdts",
	// 	"ntnada",
	// 	"svetve",
	// 	"tesnvt",
	// 	"vntsnd",
	// 	"vrdear",
	// 	"dvrsen",
	// 	"enarar",
	// }
	freq := make([][]int, 8)
	for i := range freq {
		freq[i] = make([]int, 26)
	}
	for _, l := range input {
		for i := range l {
			freq[i][l[i]-'a']++
		}
	}
	var res []rune
	for i := range freq {
		var mostFreq int
		var ch rune
		for j, v := range freq[i] {
			if v > mostFreq {
				mostFreq = v
				ch = rune(j + 'a')
			}
		}
		res = append(res, ch)
	}
	return fmt.Sprint(string(res))
}

func solve2(inf string) string {
	input := ax.MustReadFileLines("input")
	// input = []string{
	// 	"eedadn",
	// 	"drvtee",
	// 	"eandsr",
	// 	"raavrd",
	// 	"atevrs",
	// 	"tsrnev",
	// 	"sdttsa",
	// 	"rasrtv",
	// 	"nssdts",
	// 	"ntnada",
	// 	"svetve",
	// 	"tesnvt",
	// 	"vntsnd",
	// 	"vrdear",
	// 	"dvrsen",
	// 	"enarar",
	// }
	freq := make([][]int, 8)
	for i := range freq {
		freq[i] = make([]int, 26)
	}
	for _, l := range input {
		for i := range l {
			freq[i][l[i]-'a']++
		}
	}
	var res []rune
	for i := range freq {
		leastFreq := math.MaxInt32
		var ch rune
		for j, v := range freq[i] {
			if v > 0 && v < leastFreq {
				leastFreq = v
				ch = rune(j + 'a')
			}
		}
		res = append(res, ch)
	}
	return fmt.Sprint(string(res))
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
