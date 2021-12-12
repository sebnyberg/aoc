package main

import (
	"aoc/aoc2021/day10part1"
	"aoc/aoc2021/day10part2"
	"aoc/aoc2021/day11part1"
	"aoc/aoc2021/day11part2"
	"aoc/aoc2021/day1part1"
	"aoc/aoc2021/day1part2"
	"aoc/aoc2021/day2part1"
	"aoc/aoc2021/day2part2"
	"aoc/aoc2021/day3part1"
	"aoc/aoc2021/day3part2"
	"aoc/aoc2021/day4part1"
	"aoc/aoc2021/day4part2"
	"aoc/aoc2021/day5part1"
	"aoc/aoc2021/day5part2"
	"aoc/aoc2021/day6part1"
	"aoc/aoc2021/day6part2"
	"aoc/aoc2021/day7part1"
	"aoc/aoc2021/day7part2"
	"aoc/aoc2021/day8part1"
	"aoc/aoc2021/day8part2"
	"aoc/aoc2021/day9part1"
	"aoc/aoc2021/day9part2"
	"aoc/ax"
	"fmt"
	"time"
)

type solver struct {
	f    func([]string) string
	day  int
	part int
	want string
}

var solvers = []solver{
	{day1part1.Run, 1, 1, "1292"},
	{day1part2.Run, 1, 2, "1262"},
	{day2part1.Run, 2, 1, "1524750"},
	{day2part2.Run, 2, 2, "1592426537"},
	{day3part1.Run, 3, 1, "775304"},
	{day3part2.Run, 3, 2, "1592426537"},
	{day4part1.Run, 4, 1, "65325"},
	{day4part2.Run, 4, 2, "4624"},
	{day5part1.Run, 5, 1, "5092"},
	{day5part2.Run, 5, 2, "20484"},
	{day6part1.Run, 6, 1, "395627"},
	{day6part2.Run, 6, 2, "1767323539209"},
	{day7part1.Run, 7, 1, "352997"},
	{day7part2.Run, 7, 2, "101571302"},
	{day8part1.Run, 8, 1, "375"},
	{day8part2.Run, 8, 2, "1019355"},
	{day9part1.Run, 9, 1, "425"},
	{day9part2.Run, 9, 2, "1135260"},
	{day10part1.Run, 10, 1, "366027"},
	{day10part2.Run, 10, 2, "1118645287"},
	{day11part1.Run, 11, 1, "1562"},
	{day11part2.Run, 11, 2, "268"},
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("total time: %v\n", time.Since(start))
	}(time.Now())
	for _, solver := range solvers {
		start := time.Now()
		input := fmt.Sprintf("aoc2021/day%vpart%v/input", solver.day, solver.part)
		solver.f(ax.MustReadFineLines(input))
		fmt.Printf("day: %v\tpart: %v\telapsed: %v\n", solver.day, solver.part, time.Since(start))
	}
}
