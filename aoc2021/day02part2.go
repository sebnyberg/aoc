package aoc2021

import (
	"aoc/ax"
	"regexp"
	"strconv"
)

func Day02Part2(lines []string) string {
	pat := regexp.MustCompile(`^(\w+) (\d+)`)
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
