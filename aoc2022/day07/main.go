package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

var pat = regexp.MustCompile(``)

func isnum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func walkDirs(inf string) map[string]int {
	lines := ax.MustReadFileLines(inf)
	path := []string{}
	size := []int{}
	dirSize := make(map[string]int)
	var m int
	var i int
	for i < len(lines) {
		l := lines[i]
		fs := strings.Fields(l)
		if fs[0][0] == '$' {
			if fs[1] == "cd" {
				if fs[2] == ".." {
					d := strings.Join(path, "/")
					dirSize[d] += size[m-1]
					size[m-2] += size[m-1]
					path = path[:m-1]
					size = size[:m-1]
					m--
				} else {
					path = append(path, fs[2])
					size = append(size, 0)
					m++
				}
				i++
			} else if fs[1] == "ls" {
				i++
				for i < len(lines) && lines[i][0] != '$' {
					fs = strings.Fields(lines[i])
					if !isnum(fs[0]) {
						i++
						continue
					}
					x := ax.Atoi(fs[0])
					size[m-1] += x
					i++
				}
			} else {
				panic("not good")
			}
		}
	}
	// Pop remainder of stack
	for m > 1 {
		d := strings.Join(path, "/")
		dirSize[d] += size[m-1]
		size[m-2] += size[m-1]
		path = path[:m-1]
		size = size[:m-1]
		m--
	}
	dirSize["/"] = size[0]
	return dirSize
}
func solve1(inf string) any {
	dirSize := walkDirs(inf)
	var res int
	for _, sz := range dirSize {
		if sz <= 100000 {
			res += sz
		}
	}
	return res
}

func solve2(inf string) any {
	dirSize := walkDirs(inf)
	minSz := math.MaxInt32
	needAtLeast := 30000000 - (70000000 - dirSize["/"])
	for _, sz := range dirSize {
		if sz >= needAtLeast && sz <= minSz {
			minSz = sz
		}
	}
	return fmt.Sprint(minSz)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
