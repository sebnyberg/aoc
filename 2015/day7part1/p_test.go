package p_test

import (
	"log"
	"regexp"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestPart(t *testing.T) {
	lines := ax.MustReadFineLinesChan("input")
	res := run(lines)
	require.Equal(t, 16076, res)
}

var load = regexp.MustCompile(`^(\d+) -> (\w+)$`)
var loadVar = regexp.MustCompile(`^(\w+) -> (\w+)$`)
var not = regexp.MustCompile(`^NOT (\w+) -> (\w+)`)
var binop = regexp.MustCompile(`^(\w+) (\w+) (\w+) -> (\w+)`)

func run(lines chan string) int {
	// The input is a directed acyclic graph.
	// The easiest way to represent this graph is to lazily create a set of
	// nodes and functions to evaluate the nodes.
	eval := make(map[string]func() int)
	evalCache := make(map[string]int)
	cached := func(k string, f func() int) func() int {
		return func() int {
			if _, exists := evalCache[k]; !exists {
				evalCache[k] = f()
			}
			return evalCache[k]
		}
	}

	// Figuring out the value of some node in the graph is a matter of executing
	// the its evaluation statement, which in turn evaluates dependencies' eval
	// statements, and so on.
	for line := range lines {
		switch {
		case load.MatchString(line):
			parts := load.FindStringSubmatch(line)
			eval[parts[2]] = cached(parts[2], func() int {
				return ax.MustParseIntBase(parts[1], 10)
			})
		case loadVar.MatchString(line):
			parts := loadVar.FindStringSubmatch(line)
			eval[parts[2]] = cached(parts[2], func() int {
				return eval[parts[1]]()
			})
		case not.MatchString(line):
			parts := not.FindStringSubmatch(line)
			eval[parts[2]] = cached(parts[2], func() int {
				return ^eval[parts[1]]()
			})
		case binop.MatchString(line):
			parts := binop.FindStringSubmatch(line)
			eval[parts[4]] = cached(parts[4], func() int {
				xStr, yStr := parts[1], parts[3]
				var xVal, yVal int
				if onlyLetters(xStr) {
					xVal = eval[xStr]()
				} else {
					xVal = ax.MustParseIntBase(xStr, 10)
				}
				if onlyLetters(yStr) {
					yVal = eval[yStr]()
				} else {
					yVal = ax.MustParseIntBase(yStr, 10)
				}
				var res int
				switch parts[2] {
				case "AND":
					res = xVal & yVal
				case "OR":
					res = xVal | yVal
				case "LSHIFT":
					res = xVal << yVal
				case "RSHIFT":
					res = xVal >> yVal
				default:
					log.Fatalln("invalid op ", parts)
				}
				evalCache[parts[4]] = res
				return res
			})
		}
	}
	return eval["a"]()
}

func onlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
