package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

const (
	rock     = 1
	paper    = 2
	scissors = 3
)

const (
	loss = 0
	draw = 3
	win  = 6
)

func solve1() string {
	scores := map[string]int{
		// Opponent chooses Rock
		"A X": rock + draw,
		"A Y": paper + win,
		"A Z": scissors + loss,

		// Opponent chooses Paper
		"B X": rock + loss,
		"B Y": paper + draw,
		"B Z": scissors + win,

		// Opponent chooses Scissor
		"C X": rock + win,
		"C Y": paper + loss,
		"C Z": scissors + draw,
	}
	var score int
	for _, row := range ax.MustReadFileLines("input") {
		score += scores[row]
	}
	return fmt.Sprint(score)
}

func solve2() string {
	scores := map[string]int{
		// Opponent chooses Rock
		"A X": loss + scissors,
		"A Y": draw + rock,
		"A Z": win + paper,

		// Opponent chooses Paper
		"B X": loss + rock,
		"B Y": draw + paper,
		"B Z": win + scissors,

		// Opponent chooses Scissor
		"C X": loss + paper,
		"C Y": draw + scissors,
		"C Z": win + rock,
	}
	var score int
	for _, row := range ax.MustReadFileLines("input") {
		score += scores[row]
	}
	return fmt.Sprint(score)
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve1())
	fmt.Printf("Result2:\n%v\n\n", solve2())
}
