package main

import (
	"fmt"
	"math"
	"os"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

type state struct {
	playerHP      int
	armorBuffTime int
	poisonTime    int
	mana          int
	rechargeTime  int
	bossHP        int
	player        int
}

func solve1(in *input) string {
	initialState := state{
		playerHP:      50,
		armorBuffTime: 0,
		poisonTime:    0,
		mana:          500,
		rechargeTime:  0,
		bossHP:        51,
	}
	mem := make(map[state]int)
	res := dfs(mem, initialState, 9, true)
	return fmt.Sprint(res)
}

const (
	playerMe   = 0
	playerBoss = 1
)

func dfs(mem map[state]int, now state, bossDmg int, hardmode bool) int {
	if v, exists := mem[now]; exists {
		return v
	}
	if hardmode {
		if now.playerHP == 1 {
			return math.MaxInt32
		}
		now.playerHP--
	}
	now.tick()

	// Boss round is easiest, handle it first
	if now.player == playerBoss {
		if now.bossHP <= 0 {
			return 0
		}
		now.hitPlayer(bossDmg)
		if now.playerHP <= 0 {
			return math.MaxInt32
		}
		now.player = playerMe
		return dfs(mem, now, bossDmg, hardmode)
	}

	// Then consider player moves
	res := math.MaxInt32
	if now.mana >= 53 {
		// Magic missiles
		next := now
		next.mana -= 53
		next.hitBoss(4, 0)
		next.player = playerBoss
		res = ax.Min(res, 53+dfs(mem, next, bossDmg, hardmode))
	}
	if now.mana >= 73 {
		// Drain
		next := now
		next.mana -= 73
		next.hitBoss(2, 2)
		next.player = playerBoss
		res = ax.Min(res, 73+dfs(mem, next, bossDmg, hardmode))
	}
	if now.mana >= 113 && now.armorBuffTime == 0 {
		// Shield
		next := now
		next.mana -= 113
		next.hitBoss(0, 0)
		next.armorBuffTime = 8
		next.player = playerBoss
		res = ax.Min(res, 113+dfs(mem, next, bossDmg, hardmode))
	}
	if now.mana >= 173 && now.poisonTime == 0 {
		// Poison
		next := now
		next.mana -= 173
		next.hitBoss(0, 0)
		next.poisonTime = 6
		next.player = playerBoss
		res = ax.Min(res, 173+dfs(mem, next, bossDmg, hardmode))
	}
	maxMana := 300
	if hardmode {
		maxMana = 4000
	}
	if now.mana >= 229 && now.rechargeTime == 0 && now.mana <= maxMana {
		// Recharge
		next := now
		next.mana -= 229
		next.hitBoss(0, 0)
		next.rechargeTime = 5
		next.player = playerBoss
		res = ax.Min(res, 229+dfs(mem, next, bossDmg, hardmode))
	}
	mem[now] = res
	return res
}

func (s *state) tick() {
	if s.poisonTime > 0 {
		s.bossHP -= 3
		s.poisonTime--
	}
	if s.rechargeTime > 0 {
		s.mana += 101
		s.rechargeTime--
	}
	if s.armorBuffTime > 0 {
		s.armorBuffTime--
	}
}

func (s *state) hitBoss(dmg int, heal int) {
	s.playerHP += heal
	s.bossHP -= dmg
}

func (s *state) hitPlayer(dmg int) {
	if s.armorBuffTime > 0 {
		dmg = ax.Max(1, dmg-7)
	}
	s.playerHP -= dmg
}

func solve2(in *input) string {
	initialState := state{
		playerHP:      50,
		armorBuffTime: 0,
		poisonTime:    0,
		mana:          500,
		rechargeTime:  0,
		bossHP:        51,
	}
	mem := make(map[state]int)
	res := dfs(mem, initialState, 9, false)
	return fmt.Sprint(res)
}

type input struct {
	bossHP  int
	bossDmg int
}

var pat = regexp.MustCompile(`\d+`)

func (p *input) parse(s string, i int) {
	if i == 0 {
		p.bossHP = ax.Atoi(pat.FindString(s))
	} else {
		p.bossDmg = ax.Atoi(pat.FindString(s))
	}
}

func main() {
	in := new(input)
	f, _ := os.Open("input")
	rows := ax.ReadLines(f)
	for i, s := range rows {
		in.parse(s, i)
	}
	fmt.Println(in)
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
}
