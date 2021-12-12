package p_test

import (
	"aoc/ax"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	for i, tc := range []struct {
		fname string
		want  int
	}{
		{"input", 601},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			lines := ax.MustReadFineLines(tc.fname)
			require.Equal(t, tc.want, run(lines))
		})
	}
}

var pat = regexp.MustCompile(`^(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).$`)

func run(lines []string) int {

	personIdx := make(map[string]int)
	persons := make([]string, 0)
	preferences := make([][]int, 0)
	for _, line := range lines {
		parts := pat.FindStringSubmatch(line)
		a, b := parts[1], parts[4]
		if _, exists := personIdx[a]; !exists {
			personIdx[a] = len(persons)
			persons = append(persons, a)
		}
		if _, exists := personIdx[b]; !exists {
			personIdx[b] = len(persons)
			persons = append(persons, b)
		}
		sign := 1
		if parts[2] == "lose" {
			sign = -1
		}
		v := ax.MustParseInt[int](parts[3])
		i, j := personIdx[a], personIdx[b]
		preferences = append(preferences, []int{i, j, sign * v})
	}

	// There are 8! permutations, which is reasonable to do brute-force
	// Simply try all options using dfs
	persons = append(persons, "me")
	happiness := make([][]int, len(persons))
	for i := range happiness {
		happiness[i] = make([]int, len(persons))
	}
	for _, pref := range preferences {
		a, b, val := pref[0], pref[1], pref[2]
		happiness[a][b] = val
	}

	var f maxHappinessFinder
	f.happiness = happiness
	seating := make([]int, len(persons))
	f.findMaxHappiness(seating, 0, 0, len(persons))
	return f.maxHappiness
}

type maxHappinessFinder struct {
	happiness    [][]int
	maxHappiness int
}

func (f *maxHappinessFinder) findMaxHappiness(seating []int, pos, visited, n int) {
	if pos == n {
		result := calculateHappiness(f.happiness, seating)
		f.maxHappiness = ax.Max(f.maxHappiness, result)
	}
	for i := 0; i < n; i++ {
		if visited&(1<<i) > 0 {
			continue
		}
		seating[pos] = i
		f.findMaxHappiness(seating, pos+1, visited|(1<<i), n)
	}
}

func calculateHappiness(happiness [][]int, seating []int) int {
	var res int
	n := len(seating)
	for i := n; i < n*2; i++ {
		left, current, right := seating[(i-1)%n], seating[i%n], seating[(i+1)%n]
		res += happiness[current][left]
		res += happiness[current][right]
	}
	return res
}
