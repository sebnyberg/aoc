package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(`p=(-?\d+)\,(-?\d+) v=(-?\d+),(-?\d+)`)

func move(x, y, dx, dy, k, w, h int) (int, int) {
	return (x + k*dx + 1e6*w) % w, (y + k*dy + 1e6*h) % h
}

func solve1(inf string, k, w, h int) any {
	lines := ax.MustReadFileLines(inf)
	count := make([][]int, h)
	for i := range count {
		count[i] = make([]int, w)
	}
	for _, l := range lines {
		ss := pat.FindStringSubmatch(l)
		x, y := ax.Atoi(ss[1]), ax.Atoi(ss[2])
		dx, dy := ax.Atoi(ss[3]), ax.Atoi(ss[4])
		finalX, finalY := move(x, y, dx, dy, k, w, h)
		count[finalY][finalX]++
	}
	var quadrantCount [4]int
	for i := range count {
		for j, v := range count[i] {
			if h%2 == 1 && i == h/2 || w%2 == 1 && j == w/2 {
				// middle
				continue
			}
			quadrant := 2*((2*i)/h) + (2*j)/w
			quadrantCount[quadrant] += v
		}
	}
	res := 1
	for _, v := range quadrantCount {
		res *= v
	}
	return res
}

func solve2(inf string, k, w, h int) any {
	lines := ax.MustReadFileLines(inf)
	type robot struct {
		dx, dy int
	}
	state := make([][][]robot, h)
	for i := range state {
		state[i] = make([][]robot, w)
	}
	for _, l := range lines {
		ss := pat.FindStringSubmatch(l)
		x, y := ax.Atoi(ss[1]), ax.Atoi(ss[2])
		dx, dy := ax.Atoi(ss[3]), ax.Atoi(ss[4])
		state[y][x] = append(state[y][x], robot{dx, dy})
	}
	diagonalScore := func(state [][][]robot) int {
		// The idea is that a tree will have many values along diagonals.
		// Therefore, we count the number of robots in each diagonal and multiply
		// their scores. Hopefully this will be enough to detect a tree.
		var maxScore int
		for j := range state[0] {
			var count int
			for k := 0; j+k < w && k < h; k++ {
				jj := j + k
				ii := k
				if len(state[ii][jj]) > 0 {
					count++
				}
			}
			maxScore = max(maxScore, count)
		}
		return maxScore
	}

	print := func(f io.Writer, state [][][]robot) {
		for i := range state {
			for _, v := range state[i] {
				if len(v) > 0 {
					fmt.Fprintf(f, "X")
				} else {
					fmt.Fprintf(f, "%v", len(v))
				}
			}
			fmt.Fprint(f, "\n")
		}
		fmt.Fprint(f, "\n")
	}

	curr := state
	next := make([][][]robot, h)
	for i := range next {
		next[i] = make([][]robot, w)
	}

	// open output file
	outf, err := os.OpenFile("out.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := outf.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	print(outf, curr)
	var maxScore int
	var res int
	for k := 0; k < 10000; k++ {
		score := diagonalScore(curr)
		if score >= maxScore {
			maxScore = score
			outf.Seek(0, 0)
			fmt.Fprintf(outf, "k: %v\n", k)
			res = k
			print(outf, curr)
		}

		// Reset next
		for i := range next {
			for j := range next[i] {
				next[i][j] = next[i][j][:0]
			}
		}
		for i := range curr {
			for j := range curr[i] {
				for _, robot := range curr[i][j] {
					xx, yy := move(j, i, robot.dx, robot.dy, 1, w, h)
					next[yy][xx] = append(next[yy][xx], robot)
				}
			}
		}
		curr, next = next, curr
	}
	return res
}

func main() {
	// fmt.Printf("Result1:\n%v\n", solve1("testinput", 100, 11, 7))
	// fmt.Printf("Result1:\n%v\n", solve1("input", 100, 101, 103))
	fmt.Printf("Result2:\n%v\n", solve2("input", 1, 101, 103))
}
