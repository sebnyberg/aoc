package aoc2021

import (
	"strconv"
	"strings"
)

func Day08Part1(rows []string) string {
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
	return strconv.Itoa(count)
}
