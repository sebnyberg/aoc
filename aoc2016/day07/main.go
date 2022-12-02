package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	lines := ax.MustReadFileLines(inf)
	// lines = []string{
	// 	"abba[mnop]qrst",
	// 	"abcd[bddb]xyyx",
	// 	"aaaa[qwer]tyui",
	// 	"ioxxoj[asdfgh]zxcvbn",
	// }
	parse := func(l string) ([]string, []string) {
		var outerRes []string
		var innerRes []string
		var inner []byte
		var outer []byte
		var isInner bool
		for i := range l {
			if l[i] == '[' {
				isInner = true
				outerRes = append(outerRes, string(outer))
				outer = outer[:0]
			} else if l[i] == ']' {
				isInner = false
				innerRes = append(innerRes, string(inner))
				inner = inner[:0]
			} else {
				if isInner {
					inner = append(inner, l[i])
				} else {
					outer = append(outer, l[i])
				}
			}
		}
		outerRes = append(outerRes, string(outer))
		return outerRes, innerRes
	}
	hasAbba := func(s string) bool {
		for i := 0; i < len(s)-3; i++ {
			if s[i] == s[i+1] {
				continue
			}
			if s[i] == s[i+3] && s[i+1] == s[i+2] {
				return true
			}
		}
		return false
	}
	var count int
	for _, l := range lines {
		outerRes, innerRes := parse(l)
		var hasOuter bool
		for _, o := range outerRes {
			if hasAbba(o) {
				hasOuter = true
				break
			}
		}
		if !hasOuter {
			continue
		}
		var hasInner bool
		for _, inner := range innerRes {
			if hasAbba(inner) {
				hasInner = true
				break
			}
		}
		if !hasInner {
			count++
		}
	}
	return fmt.Sprint(count)
}

func solve2(inf string) string {
	lines := ax.MustReadFileLines(inf)
	// lines = []string{
	// 	"abba[mnop]qrst",
	// 	"abcd[bddb]xyyx",
	// 	"aaaa[qwer]tyui",
	// 	"ioxxoj[asdfgh]zxcvbn",
	// }
	parse := func(l string) ([]string, []string) {
		var outerRes []string
		var innerRes []string
		var inner []byte
		var outer []byte
		var isInner bool
		for i := range l {
			if l[i] == '[' {
				isInner = true
				outerRes = append(outerRes, string(outer))
				outer = outer[:0]
			} else if l[i] == ']' {
				isInner = false
				innerRes = append(innerRes, string(inner))
				inner = inner[:0]
			} else {
				if isInner {
					inner = append(inner, l[i])
				} else {
					outer = append(outer, l[i])
				}
			}
		}
		outerRes = append(outerRes, string(outer))
		return outerRes, innerRes
	}
	findAbas := func(s string) []string {
		var res []string
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+1] {
				continue
			}
			if s[i] == s[i+2] {
				res = append(res, s[i:i+3])
			}
		}
		return res
	}
	var count int
	for _, l := range lines {
		outerRes, innerRes := parse(l)
		var abas []string
		for _, o := range outerRes {
			abas = append(abas, findAbas(o)...)
		}
		babs := make([]string, len(abas))
		for i := range abas {
			babs[i] = abas[i][1:2] + abas[i][:1] + abas[i][1:2]
		}
		var hasBab bool
		for _, inner := range innerRes {
			for _, bab := range babs {
				if strings.Contains(inner, bab) {
					hasBab = true
					break
				}
			}
		}
		if hasBab {
			count++
		}
	}
	return fmt.Sprint(count)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
