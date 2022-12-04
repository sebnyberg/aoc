package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

const input = "veumntbg"

func solve() string {
	hash := func(s string) string {
		hh := md5.Sum([]byte(s))
		a := hex.EncodeToString(hh[:])
		return a
	}
	n := 4
	pos := func(s string) (int, int) {
		var i, j int
		for _, ch := range s {
			switch ch {
			case 'D':
				i++
			case 'U':
				i--
			case 'R':
				j++
			case 'L':
				j--
			}
		}
		return i, j
	}
	curr := []string{input}
	next := []string{}

	m := n
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	var maxkk int
	var shortestString string
	for kk := 1; len(curr) > 0; kk++ {
		next = next[:0]
		for _, x := range curr {
			h := hash(x)
			for i, ch := range []byte("UDLR") {
				s := x + string(ch)
				ii, jj := pos(s)
				if !ok(ii, jj) || h[i] <= 'a' {
					continue
				}
				if ii == 3 && jj == 3 {
					maxkk = ax.Max(maxkk, kk)
					if shortestString == "" {
						shortestString = s[len(input):]
					}
					continue
				}
				next = append(next, s)
			}
		}
		curr, next = next, curr
	}
	return fmt.Sprintf("shortest:\n%v\nmax:\n%v\n", shortestString, maxkk)
}

func main() {
	fmt.Printf(solve())
}
