package main

import (
	"fmt"
)

type floorState struct {
	generators int
	chips      int
}

type state struct {
	elevatorLevel int
	floors        [4]floorState
}

const (
	promethium = 1 << 0
	cobalt     = 1 << 1
	curium     = 1 << 2
	ruthenium  = 1 << 3
	plutonium  = 1 << 4
	elerium    = 1 << 5 // part 2
	dilithium  = 1 << 6 // part 2
	hydrogen   = 1 << 0
	lithium    = 1 << 1
)

func (s state) isValid() bool {
	for i := range s.floors {
		if !s.floors[i].isValid() {
			return false
		}
	}
	return true
}

func (s floorState) isValid() bool {
	if s.generators == 0 {
		return true
	}
	return s.chips&s.generators == s.chips
}

func solve1(inf string) string {
	// The elevator needs at least one generator or microchip to function.
	// The elevator can take at most two items in combination and must bring at
	// least one item at a time.
	//
	// A micro is not allowed to be on the same floor as a generator unless that
	// floor also has its own generator.
	//
	initialState := state{
		elevatorLevel: 0,
		floors: [4]floorState{
			{
				generators: promethium,
				chips:      promethium,
			},
			{
				generators: cobalt | curium | ruthenium | plutonium,
				chips:      0,
			},
			{
				generators: 0,
				chips:      cobalt | curium | ruthenium | plutonium,
			},
			{
				generators: 0,
				chips:      0,
			},
		},
	}
	nitems := 5
	want := state{
		elevatorLevel: 3,
		floors: [4]floorState{
			3: {
				generators: (1 << nitems) - 1,
				chips:      (1 << nitems) - 1,
			},
		},
	}
	curr := []state{initialState}
	next := []state{}

	seen := make(map[state]bool)
	for steps := 0; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			if x == want {
				return fmt.Sprint(steps)
			}
			nextLevels := []int{}
			if x.elevatorLevel >= 1 {
				nextLevels = append(nextLevels, x.elevatorLevel-1)
			}
			if x.elevatorLevel <= 2 {
				nextLevels = append(nextLevels, x.elevatorLevel+1)
			}
			for _, nextFloor := range nextLevels {
				for i := 0; i < nitems*2; i++ {
					nextState := x
					if i < nitems {
						if x.floors[x.elevatorLevel].generators&(1<<i) == 0 {
							continue
						}
						nextState.floors[nextFloor].generators |= 1 << i
						nextState.floors[x.elevatorLevel].generators &^= 1 << i
					} else {
						if x.floors[x.elevatorLevel].chips&(1<<(i-nitems)) == 0 {
							continue
						}
						nextState.floors[nextFloor].chips |= 1 << (i - nitems)
						nextState.floors[x.elevatorLevel].chips &^= 1 << (i - nitems)
					}
					nextState.elevatorLevel = nextFloor
					// Choose only this item
					if nextState.isValid() && !seen[nextState] {
						seen[nextState] = true
						next = append(next, nextState)
					}

					// Or add an additional item
					for j := i + 1; j < nitems*2; j++ {
						nextnext := nextState
						if j < nitems {
							if x.floors[x.elevatorLevel].generators&(1<<j) == 0 {
								continue
							}
							nextnext.floors[nextFloor].generators |= 1 << j
							nextnext.floors[x.elevatorLevel].generators &^= 1 << j
						} else {
							if x.floors[x.elevatorLevel].chips&(1<<(j-nitems)) == 0 {
								continue
							}
							nextnext.floors[nextFloor].chips |= 1 << (j - nitems)
							nextnext.floors[x.elevatorLevel].chips &^= 1 << (j - nitems)
						}
						if nextnext.isValid() && !seen[nextnext] {
							seen[nextnext] = true
							next = append(next, nextnext)
						}
					}
				}
			}
		}
		curr, next = next, curr
	}

	return fmt.Sprint("")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve2(inf string) string {
	// The elevator needs at least one generator or microchip to function.
	// The elevator can take at most two items in combination and must bring at
	// least one item at a time.
	//
	// A micro is not allowed to be on the same floor as a generator unless that
	// floor also has its own generator.
	//
	initialState := state{
		elevatorLevel: 0,
		floors: [4]floorState{
			{
				generators: promethium | elerium | dilithium,
				chips:      promethium | elerium | dilithium,
			},
			{
				generators: cobalt | curium | ruthenium | plutonium,
				chips:      0,
			},
			{
				generators: 0,
				chips:      cobalt | curium | ruthenium | plutonium,
			},
			{
				generators: 0,
				chips:      0,
			},
		},
	}
	nitems := 7
	want := state{
		elevatorLevel: 3,
		floors: [4]floorState{
			3: {
				generators: (1 << nitems) - 1,
				chips:      (1 << nitems) - 1,
			},
		},
	}
	curr := []state{initialState}
	next := []state{}

	seen := make(map[state]bool)
	for steps := 0; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			if x == want {
				return fmt.Sprint(steps)
			}
			nextLevels := []int{}
			if x.elevatorLevel >= 1 {
				nextLevels = append(nextLevels, x.elevatorLevel-1)
			}
			if x.elevatorLevel <= 2 {
				nextLevels = append(nextLevels, x.elevatorLevel+1)
			}
			for _, nextFloor := range nextLevels {
				for i := 0; i < nitems*2; i++ {
					nextState := x
					if i < nitems {
						if x.floors[x.elevatorLevel].generators&(1<<i) == 0 {
							continue
						}
						nextState.floors[nextFloor].generators |= 1 << i
						nextState.floors[x.elevatorLevel].generators &^= 1 << i
					} else {
						if x.floors[x.elevatorLevel].chips&(1<<(i-nitems)) == 0 {
							continue
						}
						nextState.floors[nextFloor].chips |= 1 << (i - nitems)
						nextState.floors[x.elevatorLevel].chips &^= 1 << (i - nitems)
					}
					nextState.elevatorLevel = nextFloor
					// Choose only this item
					if nextState.isValid() && !seen[nextState] {
						seen[nextState] = true
						next = append(next, nextState)
					}

					// Or add an additional item
					for j := i + 1; j < nitems*2; j++ {
						nextnext := nextState
						if j < nitems {
							if x.floors[x.elevatorLevel].generators&(1<<j) == 0 {
								continue
							}
							nextnext.floors[nextFloor].generators |= 1 << j
							nextnext.floors[x.elevatorLevel].generators &^= 1 << j
						} else {
							if x.floors[x.elevatorLevel].chips&(1<<(j-nitems)) == 0 {
								continue
							}
							nextnext.floors[nextFloor].chips |= 1 << (j - nitems)
							nextnext.floors[x.elevatorLevel].chips &^= 1 << (j - nitems)
						}
						if nextnext.isValid() && !seen[nextnext] {
							seen[nextnext] = true
							next = append(next, nextnext)
						}
					}
				}
			}
		}
		curr, next = next, curr
	}

	return fmt.Sprint("")
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
