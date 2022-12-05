package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	input := "abcdefgh"
	n := len(input)
	curr := []byte(input)
	var next []byte
	lines := ax.MustReadFileLines(inf)
	rotate := func(x int) {
		next = next[:0]
		for i := range curr {
			next = append(next, curr[(i-x+n)%n])
		}
		curr, next = next, curr
	}
	for _, l := range lines {
		next = next[:0]
		fs := strings.Fields(l)
		switch fs[0] {
		case "rotate":
			switch fs[1] {
			case "right":
				x := ax.Atoi(fs[2])
				rotate(x)
			case "left":
				x := ax.Atoi(fs[2])
				rotate(-x)
			case "based":
				letter := fs[6]
				idx := strings.Index(string(curr), string(letter))
				// Rotate once right
				rotate(1)
				// Rotate once per index
				rotate(idx)
				// Then an extra time if idx >= 4
				if idx >= 4 {
					rotate(1)
				}
			default:
				panic(fs)
			}
		case "swap":
			switch fs[1] {
			case "position":
				a := ax.Atoi(fs[2])
				b := ax.Atoi(fs[5])
				curr[a], curr[b] = curr[b], curr[a]
			case "letter":
				a := fs[2][0]
				b := fs[5][0]
				for i := range curr {
					if curr[i] == b {
						curr[i] = a
					} else if curr[i] == a {
						curr[i] = b
					}
				}
			default:
				panic(fs)
			}
		case "reverse":
			a := ax.Atoi(fs[2])
			b := ax.Atoi(fs[4])
			for l, r := a, b; l < r; l, r = l+1, r-1 {
				curr[l], curr[r] = curr[r], curr[l]
			}
		case "move":
			a := ax.Atoi(fs[2])
			b := ax.Atoi(fs[5])
			// Step 1, remove from position a
			x := curr[a]
			copy(curr[a:], curr[a+1:])
			// Step 2, make space in position b
			copy(curr[b+1:], curr[b:])
			curr[b] = x
		}
	}
	return fmt.Sprint(string(curr))
}

func solve2(inf string) string {
	input := "fbgdceah"
	n := len(input)
	curr := []byte(input)
	var next []byte
	lines := ax.MustReadFileLines(inf)

	for l, r := 0, len(lines)-1; l < r; l, r = l+1, r-1 {
		lines[l], lines[r] = lines[r], lines[l]
	}

	rotate := func(x int, curr []byte) []byte {
		var next []byte
		for i := range curr {
			next = append(next, curr[(i-x+n)%n])
		}
		return next
	}
	same := func(a, b []byte) bool {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	// idxRotation := func(idx int, curr []byte) []byte {
	// 	next := rotate(1, curr)
	// 	next = rotate(idx, next)
	// 	if idx >= 4 {
	// 		next = rotate(1, next)
	// 	}
	// 	return next
	// }
	for _, l := range lines {
		next = next[:0]
		fs := strings.Fields(l)
		switch fs[0] {
		case "rotate":
			switch fs[1] {
			case "right":
				want := curr
				x := ax.Atoi(fs[2])
				actual := rotate(x, rotate(-x, curr))
				if !same(want, actual) {
					panic("aa")
				}
				curr = rotate(-x, curr)
			case "left":
				want := curr
				x := ax.Atoi(fs[2])
				actual := rotate(-x, rotate(x, curr))
				if !same(want, actual) {
					panic("aa")
				}
				curr = rotate(x, curr)
			case "based":
				letter := fs[6]
				idx := strings.Index(string(curr), string(letter))
				// Try each initial position that could've resulted in the
				// current index.
				// Note that this operation is ambiguous. It turns out that AoC
				// authors used the highest possible index. Pretty terrible
				// design on AoC's part imo.
				for i := n - 1; i >= 0; i-- {
					j := i + i + 1
					if i >= 4 {
						j++
					}
					j %= n
					if j != idx {
						continue
					}
					// Undo rotation
					delta := j - i
					curr = rotate(-delta, curr)
					break
				}
			default:
				panic(fs)
			}
		case "swap":
			switch fs[1] {
			case "position":
				a := ax.Atoi(fs[2])
				b := ax.Atoi(fs[5])
				curr[a], curr[b] = curr[b], curr[a]
			case "letter":
				a := fs[2][0]
				b := fs[5][0]
				for i := range curr {
					if curr[i] == b {
						curr[i] = a
					} else if curr[i] == a {
						curr[i] = b
					}
				}
			default:
				panic(fs)
			}
		case "reverse":
			a := ax.Atoi(fs[2])
			b := ax.Atoi(fs[4])
			for l, r := a, b; l < r; l, r = l+1, r-1 {
				curr[l], curr[r] = curr[r], curr[l]
			}
		case "move":
			b := ax.Atoi(fs[2])
			a := ax.Atoi(fs[5])
			// Step 1, remove from position a
			x := curr[a]
			copy(curr[a:], curr[a+1:])
			// Step 2, make space in position b
			copy(curr[b+1:], curr[b:])
			curr[b] = x
		}
	}
	return fmt.Sprint(string(curr))
}

func main() {
	inf := "input"
	// fmt.Printf("Result1:\n%v\n", solve1(inf))
	fmt.Printf("Result1:\n%v\n", solve2(inf))
}
