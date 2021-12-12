package day2part1

import (
	"aoc/ax"
	"regexp"
	"strconv"
)

const (
	Problem = 2
	Part    = 1
)

var pat = regexp.MustCompile(`^(\w+) (\d+)`)

func Run(lines []string) string {
	var horz, depth int
	for _, line := range lines {
		parts := pat.FindStringSubmatch(line)
		dir := parts[1]
		val := ax.MustParseInt[int](parts[2])
		switch dir {
		case "forward":
			horz += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return strconv.Itoa(depth * horz)
}
