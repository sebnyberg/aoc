package main

import (
	"fmt"
	"math"

	"github.com/sebnyberg/aoc/ax"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var numpad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{0, '0', 'A'},
}

var (
	btnUp     = [2]int{0, 1}
	btnDirAck = [2]int{0, 2}
	btnLeft   = [2]int{1, 0}
	btnDown   = [2]int{1, 1}
	btnRight  = [2]int{1, 2}
)

var (
	btn7       = [2]int{0, 0}
	btn8       = [2]int{0, 1}
	btn9       = [2]int{0, 2}
	btn4       = [2]int{1, 0}
	btn5       = [2]int{1, 1}
	btn6       = [2]int{1, 2}
	btn1       = [2]int{2, 0}
	btn2       = [2]int{2, 1}
	btn3       = [2]int{2, 2}
	btnInvalid = [2]int{3, 0}
	btn0       = [2]int{3, 1}
	btnNumAck  = [2]int{3, 2}
)

var charToButton = [][2]int{
	'0': btn0,
	'1': btn1,
	'2': btn2,
	'3': btn3,
	'4': btn4,
	'5': btn5,
	'6': btn6,
	'7': btn7,
	'8': btn8,
	'9': btn9,
	'A': btnNumAck,
}

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var dirpadPos = [4][2]int{btnRight, btnLeft, btnDown, btnUp}

func fill(arr [][][][]int, dims [4]int, val int) [][][][]int {
	arr = append(arr[:0], make([][][][]int, dims[0])...)
	for i := range arr {
		arr[i] = append(arr[i][:0], make([][][]int, dims[1])...)
		for j := range arr[i] {
			arr[i][j] = append(arr[i][j][:0], make([][]int, dims[2])...)
			for k := range arr[i][j] {
				arr[i][j][k] = append(arr[i][j][k][:0], make([]int, dims[3])...)
				for l := range arr[i][j][k] {
					arr[i][j][k][l] = val
				}
			}
		}
	}
	return arr
}

type pos struct {
	q [2]int // position in pad
	d [2]int // position in controlling dirpad
}

const bigNumber = math.MaxInt64 / 2

func calcPressCost(m, n int, dirpadPressCost [][][][]int, valid [][]bool) [][][][]int {
	dm := len(dirpadPressCost)
	dn := len(dirpadPressCost[0])

	// dirpadPressCost[pi][pj][qi][qj] is the cost of moving from (pi, pj) to (qi,
	// qj), pressing the button, returning to "Ack", and pressing the button once
	// more.

	// pressCost[pi][pj][qj][qj] is the cost of moving from (pi, pj) to (qi, qj)
	// and pressing the button using a direction pad.
	pressCost := fill([][][][]int{}, [4]int{m, n, m, n}, bigNumber)

	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && valid[i][j]
	}

	curr := []pos{}
	next := []pos{}
	var moveCost [][][][]int
	for pi := 0; pi < m; pi++ {
		for pj := 0; pj < n; pj++ {
			if !valid[pi][pj] {
				continue
			}
			p0 := [2]int{pi, pj}
			curr = append(curr[:0], pos{p0, btnDirAck})
			moveCost = fill([][][][]int{}, [4]int{m, n, dm, dn}, bigNumber)
			moveCost[pi][pj][btnDirAck[0]][btnDirAck[1]] = 0
			for len(curr) > 0 {
				next = next[:0]
				for _, x := range curr {
					for i, d := range dirs {
						qi, qj := x.q[0]+d[0], x.q[1]+d[1]
						if !ok(qi, qj) {
							continue
						}
						// cost of being in position q with direction d
						cost := moveCost[x.q[0]][x.q[1]][x.d[0]][x.d[1]]

						// calculate cost of moving on the dirpad
						nextDirpadPos := dirpadPos[i]
						dirpadCost := dirpadPressCost[x.d[0]][x.d[1]][nextDirpadPos[0]][nextDirpadPos[1]]
						cost += dirpadCost

						// if new cost is better, update moveCost and add to next round
						if moveCost[qi][qj][nextDirpadPos[0]][nextDirpadPos[1]] < cost {
							continue
						}
						moveCost[qi][qj][nextDirpadPos[0]][nextDirpadPos[1]] = cost
						next = append(next, pos{[2]int{qi, qj}, nextDirpadPos})
					}
				}
				curr, next = next, curr
			}

			// we now have the minimum cost of moving from (pi, pj) to (qi, qj) and
			// having a current direction of (di, dj).
			// The minimum presscost is the cost of moving from (pi, pj) to (qi, qj)
			// plus the cost of returning to press 'A' on the dirpad.
			for qi := range pressCost {
				for qj := range pressCost[qi] {
					for di := range dirpadPressCost {
						for dj := range dirpadPressCost[di] {
							currPressCost := pressCost[pi][pj][qi][qj]
							moveCost := moveCost[qi][qj][di][dj]
							// Add cost of moving the dirpad back to 'A' and pressing the
							// button
							moveBackCost := dirpadPressCost[di][dj][btnDirAck[0]][btnDirAck[1]]
							newPressCost := moveCost + moveBackCost
							if currPressCost < newPressCost {
								continue
							}
							pressCost[pi][pj][qi][qj] = newPressCost
						}
					}
				}
			}
		}
	}

	return pressCost
}

func solve(inf string, ndirpads int) any {
	// Initialise the cost of pressing buttons on the first dirpad
	dirpadPressCost := fill([][][][]int{}, [4]int{2, 3, 2, 3}, bigNumber)
	for i := range dirpadPressCost {
		for j := range dirpadPressCost[i] {
			if i == 0 && j == 0 {
				continue // invalid state
			}
			for k := range dirpadPressCost[i][j] {
				for l := range dirpadPressCost[i][j][k] {
					dirpadPressCost[i][j][k][l] = 1
				}
			}
		}
	}

	dirpadValid := [][]bool{
		{false, true, true},
		{true, true, true},
	}
	numpadValid := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
		{false, true, true},
	}

	for k := 0; k < ndirpads; k++ {
		dirpadPressCost = calcPressCost(2, 3, dirpadPressCost, dirpadValid)
	}
	numpadPressCost := calcPressCost(4, 3, dirpadPressCost, numpadValid)

	var res int
	for _, l := range ax.MustReadFileLines(inf) {
		var totalCost int
		l = "A" + l
		numericPart := 0
		for i := 1; i < len(l); i++ {
			prev := charToButton[l[i-1]]
			next := charToButton[l[i]]
			if l[i] >= '0' && l[i] <= '9' {
				numericPart *= 10
				numericPart += int(l[i] - '0')
			}
			cost := numpadPressCost[prev[0]][prev[1]][next[0]][next[1]]
			totalCost += cost
		}
		res += numericPart * totalCost
	}

	return res
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve("input", 2))
	fmt.Printf("Result2:\n%v\n", solve("input", 25))
}
