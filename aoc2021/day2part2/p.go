package day2part2

import (
	"aoc/ax"
	"regexp"
	"strconv"
)

const (
	Problem = 2
	Part    = 2
)

var pat = regexp.MustCompile(`^(\w+) (\d+)`)

func Run(lines []string) string {
	var horz, depth, aim int
	for _, line := range lines {
		parts := pat.FindStringSubmatch(line)
		dir := parts[1]
		val := ax.MustParseInt[int](parts[2])
		switch dir {
		case "forward":
			horz += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return strconv.Itoa(depth * horz)
}
