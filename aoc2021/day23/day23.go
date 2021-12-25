package day23

import (
	"fmt"
	"strings"
)

type amphipod int8

const (
	ampAmber  = 1
	ampBronze = 2
	ampCopper = 3
	ampDesert = 4
)

func (a amphipod) String() string {
	switch a {
	case ampAmber:
		return "A"
	case ampBronze:
		return "B"
	case ampCopper:
		return "C"
	case ampDesert:
		return "D"
	}
	return "."
}

func (a amphipod) cost() int {
	switch a {
	case ampAmber:
		return 1
	case ampBronze:
		return 10
	case ampCopper:
		return 100
	case ampDesert:
		return 1000
	}
	return -1
}

// The state at a given point in time is the combination of locked rooms,
// room occupancy and the hallway.
type state struct {
	// The moment an amphipod moves into a room, the room becomes "locked"
	locked [4]bool
	// Four rooms with two amphipods, first slot is upper, second lower
	rooms [4][2]amphipod
	// The hallway has 11 slots for amphipods
	hallway [11]amphipod
}

func (s state) String() string {
	var sb strings.Builder
	sb.WriteString("#############\n")
	sb.WriteRune('#')
	for i := 0; i < 11; i++ {
		sb.WriteString(s.hallway[i].String())
	}
	sb.WriteString("#\n")
	sb.WriteString(fmt.Sprintf("###%v#%v#%v#%v###\n", s.rooms[0][0], s.rooms[1][0], s.rooms[2][0], s.rooms[3][0]))
	sb.WriteString(fmt.Sprintf("  #%v#%v#%v#%v#  \n", s.rooms[0][1], s.rooms[1][1], s.rooms[2][1], s.rooms[3][1]))
	sb.WriteString("  #########  \n")
	sb.WriteString("  #")
	writeLocked := func(locked bool) {
		if locked {
			sb.WriteString("L")
		} else {
			sb.WriteString("O")
		}
	}
	writeLocked(s.locked[0])
	sb.WriteString("#")
	writeLocked(s.locked[1])
	sb.WriteString("#")
	writeLocked(s.locked[2])
	sb.WriteString("#")
	writeLocked(s.locked[3])
	sb.WriteString("#  \n")
	return sb.String()
}
