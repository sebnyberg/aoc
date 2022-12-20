package main

import (
	"fmt"
	"regexp"

	"github.com/sebnyberg/aoc/ax"
)

type robotCosts struct {
	oreOre     int
	clayOre    int
	obsOre     int
	obsClay    int
	geodeOre   int
	geodeObs   int
	maxOreCost int
}

var pat = regexp.MustCompile(
	`Blueprint \d+: ` +
		`Each ore robot costs (\d+) ore. ` +
		`Each clay robot costs (\d+) ore. ` +
		`Each obsidian robot costs (\d+) ore and (\d+) clay. ` +
		`Each geode robot costs (\d+) ore and (\d+) obsidian.`,
)

// const (
// 	maxT         = 25
// 	maxOre       = 25
// 	maxClay      = 55
// 	maxObs       = 30
// 	maxOreRobot  = 5
// 	maxClayRobot = 20
// 	maxObsRobot  = 21
// )

// const (
// 	maxT         = 25
// 	maxOreRobot  = 5
// 	maxClayRobot = 20
// 	maxObsRobot  = 21
// 	maxOre       = 50
// 	maxClay      = 100
// 	maxObs       = 100
// )

func solve(inf string, t int, part2 bool) any {
	lines := ax.MustReadFileLines(inf)
	var res int
	if part2 {
		res = 1
	}
	for i, l := range lines {
		g := pat.FindStringSubmatch(l)
		var c robotCosts
		c.oreOre = ax.MustParseInt[int](g[1])
		c.clayOre = ax.MustParseInt[int](g[2])
		c.obsOre = ax.MustParseInt[int](g[3])
		c.obsClay = ax.MustParseInt[int](g[4])
		c.geodeOre = ax.MustParseInt[int](g[5])
		c.geodeObs = ax.MustParseInt[int](g[6])
		c.maxOreCost = ax.Max(c.obsOre, ax.Max(c.geodeOre, ax.Max(c.oreOre, c.clayOre)))

		// var dp [maxT][maxOreRobot][maxClayRobot][maxObsRobot][maxOre][maxClay][maxObs]int8
		// for i := range dp {
		// 	for j := range dp[i] {
		// 		for k := range dp[i][j] {
		// 			for ii := range dp[i][j][k] {
		// 				for jj := range dp[i][j][k][ii] {
		// 					for kk := range dp[i][j][k][ii][jj] {
		// 						for iii := range dp[i][j][k][ii][jj][kk] {
		// 							dp[i][j][k][ii][jj][kk][iii] = math.MinInt8
		// 						}
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		// var initialState state
		// initialState.oreRobots = 1
		mem := make(map[state]int)
		s := state{
			oreRobots: 1,
			t:         1,
		}
		a := dfs(mem, &c, s, int(t))
		fmt.Printf("Blueprint %v could open at most %v geodes in %v minutes\n",
			(i + 1), a, t,
		)
		if !part2 {
			res += (i + 1) * int(a)
		} else {
			res *= int(a)
		}
	}
	return res
}

var maxRecordedOre int = 0
var maxRecordedClay int = 0
var maxRecordedObs int = 0
var hits = 0

type state struct {
	oreRobots  int
	clayRobots int
	obsRobots  int
	t          int
	ore        int
	clay       int
	obs        int
}

func dfs(mem map[state]int, c *robotCosts, s state, n int) int {
	if s.t == n {
		return 0
	}
	// The difficulty lies in limiting the search space without missing out on
	// the optimal solution.
	//
	// What we can say for sure is that if the current amount of mineral +
	// incoming mineral due to existing robots exceeds the maximum amount that
	// could be used, then we can redure the static amount by the difference.
	//
	// For example, if we have 10 ores and 3 ore robots, then if there are 6
	// rounds remaining, we could never need more than 6 ores (even if we build
	// no more robots).
	//
	// if s.oreRobots == c.maxOreCost {
	// 	s.ore = int(s.oreRobots)
	// }
	// if s.clayRobots == c.obsClay {
	// 	s.clay = int(s.clayRobots)
	// }
	// if s.obsRobots == c.geodeObs {
	// 	s.obs = int(s.obsRobots)
	// }
	// tt := (n - s.t + 1)
	// // s.ore = ax.Min(s.ore, (tt)*(c.maxOreCost-int(s.oreRobots)))
	// if s.ore > maxRecordedOre {
	// 	fmt.Println("ore", s.ore)
	// 	maxRecordedOre = s.ore
	// }
	// // clay = ax.Min(clay, (t+1)*(c.obsClay-clayR))
	// if s.clay > maxRecordedClay {
	// 	fmt.Println("clay", s.clay)
	// 	maxRecordedClay = s.clay
	// }
	// // obs = ax.Min(obs, (t+1)*(c.geodeObs-obsR))
	// if s.obs > maxRecordedObs {
	// 	fmt.Println("obs", s.obs)
	// 	maxRecordedObs = s.obs
	// }
	// nextOre := ore + oreR
	// nextClay := clay + clayR
	// nextObs := obs + obsR
	if v, exists := mem[s]; exists {
		return v
	}
	// want := state{
	// 	t:         6, // according to the list, this should be the next T
	// 	ore:       1,
	// 	oreRobots: 2,
	// }
	// want := state{
	// 	t:          9 + 1, // according to the list, this should be the next T
	// 	ore:        2,
	// 	oreRobots:  2,
	// 	clay:       3,
	// 	clayRobots: 3,
	// }
	// want := state{
	// 	t:          12 + 1, // according to the list, this should be the next T
	// 	ore:        3,
	// 	oreRobots:  2,
	// 	clay:       15,
	// 	clayRobots: 6,
	// 	obs:        0,
	// 	obsRobots:  0,
	// }
	// want := state{
	// 	t:          13 + 1, // according to the list, this should be the next T
	// 	ore:        3,
	// 	oreRobots:  2,
	// 	clay:       21,
	// 	clayRobots: 7,
	// 	obs:        0,
	// 	obsRobots:  0,
	// }
	// want := state{
	// 	t:          14 + 1, // according to the list, this should be the next T
	// 	ore:        2,
	// 	oreRobots:  2,
	// 	clay:       14,
	// 	clayRobots: 7,
	// 	obs:        0,
	// 	obsRobots:  1,
	// }
	// want := state{
	// 	t:          16 + 1, // according to the list, this should be the next T
	// 	ore:        3,
	// 	oreRobots:  2,
	// 	clay:       14,
	// 	clayRobots: 7,
	// 	obs:        2,
	// 	obsRobots:  2,
	// }

	// Do nothing
	nextttt := s.tick()

	res := dfs(mem, c, nextttt, n)

	if s.ore >= int(c.oreOre) && s.oreRobots < c.maxOreCost {
		// Build an ore robot
		next := s.tick()
		next.ore -= int(c.oreOre)
		next.oreRobots++
		a := dfs(mem, c, next, n)
		res = ax.Max(res, a)
	}

	// if s == want {
	// 	fmt.Print("")
	// }
	if s.ore >= int(c.clayOre) && s.clayRobots < int(c.obsClay) {
		// Build an clay robot
		next := s.tick()
		next.ore -= int(c.clayOre)
		next.clayRobots++
		a := dfs(mem, c, next, n)
		res = ax.Max(res, a)
	}

	if s.ore >= int(c.obsOre) && s.clay >= int(c.obsClay) && s.obsRobots < c.geodeObs {
		// Build an obsidian robot
		next := s.tick()
		next.ore -= int(c.obsOre)
		next.clay -= int(c.obsClay)
		next.obsRobots++
		a := dfs(mem, c, next, n)
		res = ax.Max(res, a)
	}

	if s.ore >= int(c.geodeOre) && s.obs >= int(c.geodeObs) {
		// Build a geode robot
		next := s.tick()
		next.ore -= int(c.geodeOre)
		next.obs -= int(c.geodeObs)
		a := dfs(mem, c, next, n)
		res = ax.Max(res, (n-s.t)+a)
	}

	mem[s] = res
	return res
}

func (s state) tick() state {
	s.ore += int(s.oreRobots)
	s.clay += int(s.clayRobots)
	s.obs += int(s.obsRobots)
	s.t++
	return s
}

func main() {
	// f := "input"
	fmt.Printf("Result1:\n%v\n", solve("input", 24, false))
	fmt.Printf("Result2:\n%v\n", solve("input3", 32, true))
}
