package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f, false, 2022))
	fmt.Printf("Result2:\n%v\n", solve2(f, false, 1000000000000))
}

func solve2(inf string, debug bool, nrocktotal int) any {
	rocks := []rock{
		newRock(`####`),
		newRock(`.#|
###
.#.`),
		newRock(`..#
..#
###`),
		newRock(`#
#
#
#`),
		newRock(`##
##`),
	}

	// Check whether the rock r would hit anything it it were
	// moved to row j in the state
	hit := func(state []byte, j int, r rock) bool {
		if j == -1 {
			return true
		}
		for k := 0; k < 4; k++ {
			if state[j+k]&r.enc[k] > 0 {
				return true
			}
		}
		return false
	}

	move := map[byte]func(r rock) rock{
		'>': func(r rock) rock {
			return r.shiftRight()
		},
		'<': func(r rock) rock {
			return r.shiftLeft()
		},
	}

	stateStr := func(state []byte, r rock, ri int) string {
		var sb strings.Builder
		for j := len(state) - 1; j >= 0; j-- {
			var row []byte
			for i := 0; i < 7; i++ {
				if state[j]&(1<<i) > 0 {
					row = append(row, '#')
				} else {
					row = append(row, '.')
				}
			}
			if j >= ri && j <= ri+3 {
				for i := 0; i < 7; i++ {
					if r.enc[j-ri]&(1<<i) > 0 {
						row[i] = '@'
					}
				}
			}
			sb.WriteString(string(row))
			sb.WriteByte('\n')
		}
		return sb.String()
	}
	p := func(state []byte, r rock, ri int) {
		if debug {
			fmt.Println(stateStr(state, r, ri))

		}
	}

	var state []byte

	var ri int // rock row
	var hi int // highest clean row in the room
	var pi int // push index

	pushes := ax.MustReadFileLines(inf)[0]

	type stateKey struct {
		s      string
		ri, pi int
	}
	type stateValue struct {
		nrock  int
		height int
		dn     int
		dh     int
	}
	seen := make(map[stateKey]stateValue)
	var skippedHeight int
	var skipped bool

	for nrock := 1; nrock <= nrocktotal; nrock++ {
		// Start two units from the left wall and three units
		// above the highest rock in the room
		j := hi + 3 // Two units above clean row

		// Make room for the new rock
		if len(state) < j+4 {
			state = append(state, make([]byte, j+4-len(state))...)
		}

		r := rocks[ri]
		ri = (ri + 1) % len(rocks)

		// Alternate between push / move
		for {
			p(state, r, j)
			if !hit(state, j, move[pushes[pi]](r)) {
				r = move[pushes[pi]](r)
			}
			pi = (pi + 1) % len(pushes)
			p(state, r, j)
			if hit(state, j-1, r) {
				break
			}
			j--
		}

		// Place rock in state
		for i := range r.enc {
			state[j+i] |= r.enc[i]
		}
		p(state, r, j)

		// Part 2:
		// We need to detect a recurrence in the state
		//
		// Let's try to work with the "way the things look" and the current rock
		// + push index

		// Find top row
		for state[hi] != 0 {
			hi++
		}

		top := stateStr(state[j:], r, math.MaxInt32)
		k := stateKey{
			s:  top,
			ri: ri,
			pi: pi,
		}
		dn := nrock - seen[k].nrock
		dh := hi - seen[k].height
		v := stateValue{
			nrock:  nrock,
			height: hi,
			dn:     dn,
			dh:     dh,
		}

		if !skipped && seen[k].dn == dn && seen[k].dh == dh {
			// Recursion detected.
			// We can skip *a lot* of rounds by taking nrock % nrockdelta and
			// adding the missing height to the result
			// rockDelta := v.nrock - seen[k].nrock
			// heightDelta := v.height - seen[k].height
			remains := nrocktotal - nrock
			skippedLoops := remains / dn
			nrocktotal -= skippedLoops * dn
			skippedHeight += skippedLoops * dh
			skipped = true
		}
		seen[k] = v

		// Find top row
		for state[hi] != 0 {
			hi++
		}
	}
	return hi + skippedHeight
}

func solve1(inf string, debug bool, nrocktotal int) any {
	rocks := []rock{
		newRock(`####`),
		newRock(`.#|
###
.#.`),
		newRock(`..#
..#
###`),
		newRock(`#
#
#
#`),
		newRock(`##
##`),
	}

	// Check whether the rock r would hit anything it it were
	// moved to row j in the state
	hit := func(state []byte, j int, r rock) bool {
		if j == -1 {
			return true
		}
		for k := 0; k < 4; k++ {
			if state[j+k]&r.enc[k] > 0 {
				return true
			}
		}
		return false
	}

	move := map[byte]func(r rock) rock{
		'>': func(r rock) rock {
			return r.shiftRight()
		},
		'<': func(r rock) rock {
			return r.shiftLeft()
		},
	}

	stateStr := func(state []byte, r rock, ri int) string {
		var sb strings.Builder
		for j := len(state) - 1; j >= 0; j-- {
			var row []byte
			for i := 0; i < 7; i++ {
				if state[j]&(1<<i) > 0 {
					row = append(row, '#')
				} else {
					row = append(row, '.')
				}
			}
			if j >= ri && j <= ri+3 {
				for i := 0; i < 7; i++ {
					if r.enc[j-ri]&(1<<i) > 0 {
						row[i] = '@'
					}
				}
			}
			sb.WriteString(string(row))
			sb.WriteByte('\n')
		}
		return sb.String()
	}
	p := func(state []byte, r rock, ri int) {
		if debug {
			fmt.Println(stateStr(state, r, ri))

		}
	}

	var state []byte

	var ri int // rock row
	var hi int // highest clean row in the room
	var pi int // push index

	pushes := ax.MustReadFileLines(inf)[0]

	for nrock := 1; nrock <= nrocktotal; nrock++ {
		// Start two units from the left wall and three units
		// above the highest rock in the room
		j := hi + 3 // Two units above clean row

		// Make room for the new rock
		if len(state) < j+4 {
			state = append(state, make([]byte, j+4-len(state))...)
		}

		r := rocks[ri]
		ri = (ri + 1) % len(rocks)

		// Alternate between push / move
		for {
			p(state, r, j)
			if !hit(state, j, move[pushes[pi]](r)) {
				r = move[pushes[pi]](r)
			}
			pi = (pi + 1) % len(pushes)
			p(state, r, j)
			if hit(state, j-1, r) {
				break
			}
			j--
		}
		// Place rock in state
		for i := range r.enc {
			state[j+i] |= r.enc[i]
		}
		p(state, r, j)
		// Find top row
		for state[hi] != 0 {
			hi++
		}
	}

	return hi
}

type rock struct {
	enc [4]byte
}

func newRock(s string) rock {
	var r rock
	lines := strings.Split(s, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		ss := lines[i]
		var x byte
		for i, ch := range ss {
			if ch == '#' {
				x |= (1 << i)
			}
		}
		r.enc[len(lines)-1-i] = x
	}
	for i := range r.enc {
		r.enc[i] <<= 2
	}
	return r
}

func (r rock) String() string {
	var sb strings.Builder
	for i := 3; i >= 0; i-- {
		for j := 0; j <= 6; j++ {
			if r.enc[i]&(1<<j) > 0 {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (r rock) shiftLeft() rock {
	for i := range r.enc {
		if r.enc[i]&1 > 0 {
			// Do nothing
			return r
		}
	}
	// Shift left
	for i := range r.enc {
		r.enc[i] >>= 1
	}
	return r
}

func (r rock) shiftRight() rock {
	for i := range r.enc {
		if r.enc[i]&(1<<6) > 0 {
			// Do nothing
			return r
		}
	}
	// Shift right
	for i := range r.enc {
		r.enc[i] <<= 1
	}
	return r
}
