package day8part1

import (
	"strings"
)

const (
	Problem = 8
	Part    = 1
)

func Run(rows []string) int {
	var count int
	for _, row := range rows {
		parts := strings.Split(row, "|")
		outputFields := strings.Fields(parts[1])
		for _, field := range outputFields {
			if len(field) == 2 || len(field) == 4 || len(field) == 3 || len(field) == 7 {
				count++
			}
		}
	}
	return count
}
