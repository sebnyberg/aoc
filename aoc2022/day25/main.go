package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	forward := [256]byte{
		'=': '0',
		'-': '1',
		'0': '2',
		'1': '3',
		'2': '4',
	}
	backward := [256]byte{
		'0': '=',
		'1': '-',
		'2': '0',
		'3': '1',
		'4': '2',
	}
	parseSnafu := func(s string) int {
		// To parse a SNAFU number, forward-map, parse as base-5, then remove
		// 2*5^i per character
		b := []byte(s)
		for i := range b {
			b[i] = forward[b[i]]
		}
		x64, _ := strconv.ParseInt(string(b), 5, 64)
		x := int(x64)
		pow := 1
		for range s {
			x -= pow * 2
			pow *= 5
		}
		return x
	}
	rev := func(s string) string {
		bs := []byte(s)
		for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
			bs[l], bs[r] = bs[r], bs[l]
		}
		return string(bs)
	}

	toSnafu := func(x int) string {
		// To construct a snafu number, convert it to base-5, add 2*5^i to each
		// index, then do a backward mapping.
		s := strconv.FormatInt(int64(x), 5)
		var carry int
		var res []byte
		for i := len(s) - 1; i >= 0; i-- {
			x := int(s[i]-'0') + 2 + carry
			carry = x / 5
			res = append(res, byte(x%5+'0'))
		}
		if carry > 0 {
			res = append(res, byte(carry+2+'0'))
		}
		for i := range res {
			res[i] = backward[res[i]]
		}
		return rev(string(res))
	}
	var sum int
	for _, l := range lines {
		x := parseSnafu(l)
		sum += x
	}
	return toSnafu(sum)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
}
