package day23

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

// The stateEnergy at a given point in time is the combination of locked rooms,
// room occupancy and the hallway.
type stateEnergy struct {
	energy int
	state
	parents []stateEnergy
}

func (s stateEnergy) copy() stateEnergy {
	cpy := s
	cpy.parents = make([]stateEnergy, len(s.parents))
	copy(cpy.parents, s.parents)
	return cpy
}

type state struct {
	// Four rooms with two amphipods, first slot is upper, second lower
	rooms [4]string

	// The hallway has 10 slots for amphipods
	hallway string
}

func IsExactly(s string, ch rune) bool {
	for _, c := range s {
		if c != ch {
			return false
		}
	}
	return true
}

func (s stateEnergy) valid() bool {
	return IsExactly(s.hallway, '.') &&
		IsExactly(s.rooms[0], 'A') &&
		IsExactly(s.rooms[1], 'B') &&
		IsExactly(s.rooms[2], 'C') &&
		IsExactly(s.rooms[3], 'D')
}

func (s stateEnergy) String() string {
	var sb strings.Builder
	sb.WriteString("#############\n")
	sb.WriteRune('#')
	sb.WriteString(s.hallway)
	sb.WriteString("#\n")
	sb.WriteString(fmt.Sprintf("###%v#%v#%v#%v###\n", s.rooms[0][0], s.rooms[1][0], s.rooms[2][0], s.rooms[3][0]))
	sb.WriteString(fmt.Sprintf("  #%v#%v#%v#%v#  \n", s.rooms[0][1], s.rooms[1][1], s.rooms[2][1], s.rooms[3][1]))
	sb.WriteString("  #########  \n")
	sb.WriteString("  #")
	return sb.String()
}

func cost(a byte) int {
	switch a {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}
	return -1000000
}

func getMovesFromHallway(s stateEnergy) []stateEnergy {
	// checkDestRoom checks if the element at hallPos can be moved to its
	// destination room. Returns true and the next state
	// or -1, -1 if moving the current hall pos was not possible.
	checkDestRoom := func(hallPos int) (next stateEnergy, ok bool) {
		a := s.hallway[hallPos]
		destRoom := int(a - 'A')

		// Check if the way is clear
		roomPos := 2 + (2 * destRoom)
		cpy := s.copy()
		cpy.parents = append(cpy.parents, s)
		cpy.energy += cost(a) * ax.Abs(hallPos-roomPos)
		l, r := ax.MinMax(roomPos, hallPos)
		for i := l; i <= r; i++ {
			if i != hallPos && cpy.hallway[i] != '.' {
				return cpy, false
			}
		}
		// Destination room may only contain the 'right' amphipod
		for i := 0; i < len(cpy.rooms[destRoom]); i++ {
			v := cpy.rooms[destRoom][i]
			if v != '.' && v != a {
				return cpy, false
			}
		}

		cpy.hallway = strWithByte(cpy.hallway, '.', hallPos)
		// Find first empty slot in the room
		for i := len(cpy.rooms[destRoom]) - 1; i >= 0; i-- {
			v := cpy.rooms[destRoom][i]
			if v == '.' {
				cpy.rooms[destRoom] = strWithByte(cpy.rooms[destRoom], a, i)
				cpy.energy += (i + 1) * cost(a)
				return cpy, true
			}
		}
		return cpy, false
	}

	// Any amphipod currently in the hallway is moved to any possible destination
	// room. A possible destination is any room which contains only similar
	// amphipods, or is empty, and to which the hallway is not blocked.
	var res []stateEnergy
	for hallPos := 0; hallPos < len(s.hallway); hallPos++ {
		if s.hallway[hallPos] == '.' {
			continue
		}

		if next, ok := checkDestRoom(hallPos); ok {
			res = append(res, next)
		}
	}
	return res
}

func getMovesFromRooms(s stateEnergy) []stateEnergy {
	getHallwayMoves := func(a byte, i int) []stateEnergy {
		var res []stateEnergy
		// Can stop in any position not immediately outside a room
		l := 2 + (2 * i) - 1
		for l >= 0 && s.hallway[l] == '.' {
			if l < 2 || l%2 == 1 {
				cpy := s.copy()
				cpy.parents = append(cpy.parents, s)
				cpy.hallway = strWithByte(cpy.hallway, a, l)
				cpy.energy += ax.Abs((2+2*i)-l) * cost(a)
				res = append(res, cpy)
			}
			l--
		}
		r := 2 + (2 * i) + 1
		for r <= 10 && s.hallway[r] == '.' {
			if r > 8 || r%2 == 1 {
				cpy := s.copy()
				cpy.parents = append(cpy.parents, s)
				cpy.hallway = strWithByte(cpy.hallway, a, r)
				cpy.energy += ax.Abs((2+2*i)-r) * cost(a)
				res = append(res, cpy)
			}
			r++
		}
		return res
	}

	var res []stateEnergy
	for room := range s.rooms {
		for i := range s.rooms[room] {
			amph := s.rooms[room][i]
			if amph == '.' {
				continue
			}
			moves := getHallwayMoves(amph, room)
			for j := range moves {
				moves[j].rooms[room] = strWithByte(moves[j].rooms[room], '.', i)
				moves[j].energy += (i + 1) * cost(amph)
			}
			res = append(res, moves...)
			break
		}
	}
	return res
}

func strWithByte(s string, a byte, i int) string {
	return s[:i] + string(a) + s[i+1:]
}

type StateHeap []stateEnergy

func (h StateHeap) Len() int { return len(h) }
func (h StateHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h StateHeap) Less(i, j int) bool {
	return h[i].energy < h[j].energy
}
func (h *StateHeap) Push(x interface{}) {
	*h = append(*h, x.(stateEnergy))
}
func (h *StateHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
