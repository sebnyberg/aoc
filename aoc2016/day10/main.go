package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	lines := ax.MustReadFileLines(inf)
	bots := make(map[int]*bot)
	outs := make(map[int]int)
	initBot := func(x int) {
		if _, exists := bots[x]; !exists {
			bots[x] = &bot{}
		}
	}
	for _, l := range lines {
		fs := strings.Fields(l)
		if fs[0] == "value" {
			botNbr := ax.Atoi(fs[len(fs)-1])
			val := ax.Atoi(fs[1])
			initBot(botNbr)
			bots[botNbr].values = append(bots[botNbr].values, val)
		} else {
			from := ax.Atoi(fs[1])
			lowBot := fs[5] == "bot"
			low := ax.Atoi(fs[6])
			high := ax.Atoi(fs[11])
			initBot(from)
			if lowBot {
				initBot(low)
				bots[from].lowDest = low
			} else {
				bots[from].lowDest = -(1 + low)
			}
			initBot(high)
			bots[from].highDest = high
		}
	}
	// Loop through all bots until finding a match
	for {
		for id, bot := range bots {
			if len(bot.values) <= 1 {
				continue
			}
			// Pass values according to scheme
			sort.Ints(bot.values)
			if bot.values[0] == 17 && bot.values[1] == 61 {
				return fmt.Sprint(id)
			}
			if bot.lowDest < 0 {
				actual := -(1 + bot.lowDest)
				outs[actual] = bot.values[0]
			} else {
				v := bot.values[0]
				bots[bot.lowDest].values = append(bots[bot.lowDest].values, v)
			}
			v := bot.values[1]
			bots[bot.highDest].values = append(bots[bot.highDest].values, v)
			bot.values = bot.values[:0]
		}
	}
}

type bot struct {
	values   []int
	lowDest  int
	highDest int
}

func solve2(inf string) string {
	lines := ax.MustReadFileLines(inf)
	bots := make(map[int]*bot)
	outs := make(map[int]int)
	initBot := func(x int) {
		if _, exists := bots[x]; !exists {
			bots[x] = &bot{}
		}
	}
	for _, l := range lines {
		fs := strings.Fields(l)
		if fs[0] == "value" {
			botNbr := ax.Atoi(fs[len(fs)-1])
			val := ax.Atoi(fs[1])
			initBot(botNbr)
			bots[botNbr].values = append(bots[botNbr].values, val)
		} else {
			from := ax.Atoi(fs[1])
			lowBot := fs[5] == "bot"
			low := ax.Atoi(fs[6])
			high := ax.Atoi(fs[11])
			initBot(from)
			if lowBot {
				initBot(low)
				bots[from].lowDest = low
			} else {
				bots[from].lowDest = -(1 + low)
			}
			initBot(high)
			bots[from].highDest = high
		}
	}
	// Loop through all bots until finding a match
	for {
		exists := func(x int) bool {
			_, ok := outs[x]
			return ok
		}
		if exists(0) && exists(1) && exists(2) {
			return fmt.Sprint(outs[0] * outs[1] * outs[2])
		}
		for _, bot := range bots {
			if len(bot.values) <= 1 {
				continue
			}
			// Pass values according to scheme
			sort.Ints(bot.values)
			if bot.lowDest < 0 {
				actual := -(1 + bot.lowDest)
				outs[actual] = bot.values[0]
			} else {
				v := bot.values[0]
				bots[bot.lowDest].values = append(bots[bot.lowDest].values, v)
			}
			v := bot.values[1]
			bots[bot.highDest].values = append(bots[bot.highDest].values, v)
			bot.values = bot.values[:0]
		}
	}
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
