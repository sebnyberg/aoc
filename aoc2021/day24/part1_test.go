package day23

import (
	"aoc/ax"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day23part1 int

func BenchmarkDay23Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day23part1 = Part1(ax.MustReadFineLines("input"))
	}
}

func TestDay23Part1(t *testing.T) {
	// assert.Equal(t, 12521, Part1(ax.MustReadFineLines("small")))
	assert.Equal(t, 16244, Part1(ax.MustReadFineLines("input")))
}

type evalFn func(program) program

type program struct {
	pc       int
	vals     []int
	input    string
	inputPos int
}

func Part1(rows []string) int {
	instructions := make([]evalFn, 0)
	pat := regexp.MustCompile(`^(\w+)\s([-a-z0-9]+)(\s([-a-z0-9]+))?$`)
	for _, row := range rows {
		parts := pat.FindStringSubmatch(row)
		switch parts[1] {
		case "inp":
			instructions = append(instructions, func(p program) program {
				first := parts[2][0] - 'w'
				p.vals[first] = int(p.input[p.inputPos] - '0')
				p.inputPos++
				p.pc++
				return p
			})
		case "mul":
			instructions = append(instructions, func(p program) program {
				first := parts[2][0] - 'w'
				if strings.ContainsRune("wxyz", rune(parts[4][0])) {
					second := parts[4][0] - 'w'
					p.vals[first] *= p.vals[second]
				} else { // numerical
					p.vals[first] *= ax.MustParseInt[int](parts[4])
				}
				p.pc++
				return p
			})
		case "add":
			instructions = append(instructions, func(p program) program {
				first := parts[2][0] - 'w'
				if strings.ContainsRune("wxyz", rune(parts[4][0])) {
					second := parts[4][0] - 'w'
					p.vals[first] += p.vals[second]
				} else { // numerical
					p.vals[first] += ax.MustParseInt[int](parts[4])
				}
				p.pc++
				return p
			})
		case "mod":
			instructions = append(instructions, func(p program) program {
				first := parts[2][0] - 'w'
				if strings.ContainsRune("wxyz", rune(parts[4][0])) {
					second := parts[4][0] - 'w'
					p.vals[first] %= p.vals[second]
				} else { // numerical
					p.vals[first] %= ax.MustParseInt[int](parts[4])
				}
				p.pc++
				return p
			})
		case "div":
			instructions = append(instructions, func(p program) program {
				first := parts[2][0] - 'w'
				if strings.ContainsRune("wxyz", rune(parts[4][0])) {
					second := parts[4][0] - 'w'
					p.vals[first] /= p.vals[second]
				} else { // numerical
					p.vals[first] /= ax.MustParseInt[int](parts[4])
				}
				p.pc++
				return p
			})
		case "eql":
			instructions = append(instructions, func(p program) program {
				first := parts[2][0] - 'w'
				if strings.ContainsRune("wxyz", rune(parts[4][0])) {
					second := parts[4][0] - 'w'
					if p.vals[first] == p.vals[second] {
						p.vals[first] = 1
					} else {
						p.vals[first] = 0
					}
				} else { // numerical
					if p.vals[first] == ax.MustParseInt[int](parts[4]) {
						p.vals[first] = 1
					} else {
						p.vals[first] = 0
					}
				}
				p.pc++
				return p
			})
		}
	}

	for v := 11111111111111; v <= 99999999999999; v++ {
		p := program{
			input: strconv.Itoa(v),
			vals:  make([]int, 4),
		}
		for i := range instructions {
			p = instructions[i](p)
		}
		if p.vals[3] == 0 {
			return v
		}
	}
	return 0
}
