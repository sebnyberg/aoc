package main

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"regexp"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func findMinMax(in *input) (int, int) {
	p := player{
		hp: 100,
	}
	nw := len(in.s.weapons)
	na := len(in.s.armors)
	nr := len(in.s.rings)

	minCost := math.MaxInt32
	var maxCost int

	// Pick a weapon
	for w := 0; w <= (1<<nw)-1; w++ {
		if bits.OnesCount(uint(w)) != 1 {
			continue
		}
		var cost int
		p := p
		for i := 0; i < nw; i++ {
			if w&(1<<i) > 0 {
				cost += in.s.weapons[i].cost
				p = p.withItem(in.s.weapons[i])
				break
			}
		}

		// Pick 0-1 armor
		for a := 0; a <= (1<<na)-1; a++ {
			if bits.OnesCount(uint(a)) > 1 {
				continue
			}
			p := p
			cost := cost
			for i := 0; i < na; i++ {
				if a&(1<<i) > 0 {
					cost += in.s.armors[i].cost
					p = p.withItem(in.s.armors[i])
				}
			}

			// Pick 0-2 rings
			for r := 0; r <= (1<<nr)-1; r++ {
				if bits.OnesCount(uint(r)) > 2 {
					continue
				}
				p := p
				cost := cost
				for i := 0; i < nr; i++ {
					if r&(1<<i) > 0 {
						cost += in.s.rings[i].cost
						p = p.withItem(in.s.rings[i])
					}
				}

				// Finally ready to battle
				myDmg := ax.Max(1, p.dmg-in.boss.armor)
				bossDmg := ax.Max(1, in.boss.dmg-p.armor)
				myHP := p.hp
				bossHP := in.boss.hp
				var playerWon bool
				for i := 0; ; i++ {
					if i&1 == 0 {
						bossHP -= myDmg
						if bossHP <= 0 {
							playerWon = true
							break
						}
					} else {
						myHP -= bossDmg
						if myHP <= 0 {
							break
						}
					}
				}

				if playerWon && cost < minCost {
					minCost = cost
				}
				if !playerWon && cost > maxCost {
					maxCost = cost
				}
			}
		}

	}
	return minCost, maxCost
}

type player struct {
	dmg   int
	hp    int
	armor int
}

func (p player) withItem(x item) player {
	return player{
		dmg:   p.dmg + x.dmg,
		armor: p.armor + x.armor,
		hp:    p.hp,
	}
}

type shop struct {
	armors  []item
	weapons []item
	rings   []item
}

type item struct {
	dmg   int
	armor int
	cost  int
}

type inputItem struct {
	s       string
	t       string
	a, b, c int
	x, y    int
	x1, y1  int
	x2, y2  int
}

type input struct {
	boss player
	s    shop
}

var pat = regexp.MustCompile(`\d+`)
var pat2 = regexp.MustCompile(`\s\d+`)

func (p *input) parse(s string, i int) {
	if i == 0 {
		p.boss.hp = ax.Atoi(pat.FindString(s))
	} else if i == 1 {
		p.boss.dmg = ax.Atoi(pat.FindString(s))
	} else if i == 2 {
		p.boss.armor = ax.Atoi(pat.FindString(s))
	} else if s[0] == 'W' {
		ss := pat2.FindAllString(s, 100)
		p.s.weapons = append(p.s.weapons, item{
			cost:  ax.Atoi(strings.Trim(ss[0], " ")),
			dmg:   ax.Atoi(strings.Trim(ss[1], " ")),
			armor: ax.Atoi(strings.Trim(ss[2], " ")),
		})
	} else if s[0] == 'A' {
		ss := pat2.FindAllString(s, 100)
		p.s.armors = append(p.s.armors, item{
			cost:  ax.Atoi(strings.Trim(ss[0], " ")),
			dmg:   ax.Atoi(strings.Trim(ss[1], " ")),
			armor: ax.Atoi(strings.Trim(ss[2], " ")),
		})
	} else if s[0] == 'R' {
		ss := pat2.FindAllString(s, 100)
		p.s.rings = append(p.s.rings, item{
			cost:  ax.Atoi(strings.Trim(ss[0], " ")),
			dmg:   ax.Atoi(strings.Trim(ss[1], " ")),
			armor: ax.Atoi(strings.Trim(ss[2], " ")),
		})
	}
}

func solve1(in *input) string {
	minCost, _ := findMinMax(in)
	return fmt.Sprint(minCost)
}

func solve2(in *input) string {
	_, maxCost := findMinMax(in)
	return fmt.Sprint(maxCost)
}

func main() {
	in := new(input)
	f, _ := os.Open("input")
	rows := ax.ReadLines(f)
	for i, s := range rows {
		in.parse(s, i)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
}
