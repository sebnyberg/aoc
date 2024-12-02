package day18

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/require"
)

func Test_Part1(t *testing.T) {
	input := ax.MustReadFileLines("input")
	cur := parse(input[0])
	for cur.reduce() {
	}
	for i := 1; i < len(input); i++ {
		other := parse(input[i])
		for other.reduce() {
		}
		combStr := fmt.Sprintf("[%v,%v]", cur, other)
		comb := parse(combStr)
		for comb.reduce() {
			comb = parse(comb.String())
		}
		cur = comb
	}
	require.Equal(t, 4140, cur.magnitude())
}

func Test_Part2(t *testing.T) {
	inputs := ax.MustReadFileLines("input")
	var maxMagnitude int
	for i := 0; i < len(inputs)-1; i++ {
		for j := i + 1; j < len(inputs); j++ {
			comb1 := parse(fmt.Sprintf("[%v,%v]", inputs[i], inputs[j]))
			comb2 := parse(fmt.Sprintf("[%v,%v]", inputs[j], inputs[i]))
			for comb1.reduce() {
			}
			for comb2.reduce() {
			}
			maxMagnitude = ax.Max(maxMagnitude, comb1.magnitude())
			maxMagnitude = ax.Max(maxMagnitude, comb2.magnitude())
		}
	}
	require.Equal(t, 3993, maxMagnitude)
}

func Test_magnitude(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			res := parse(tc.input).magnitude()
			require.Equal(t, tc.want, res)
		})
	}
}

func Test_parseNode(t *testing.T) {
	for _, tc := range []struct {
		input string
	}{
		{"[1,2]"},
		{"[[1,2],3]"},
		{"[9,[8,7]]"},
		{"[[1,9],[8,5]]"},
		{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"},
		{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"},
		{"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			res := parse(tc.input)
			require.Equal(t, tc.input, res.String())
		})
	}
}

func Test_reduce(t *testing.T) {
	for _, tc := range []struct {
		input, want string
	}{
		// Split
		{"[10,1]", "[[5,5],1]"},
		// Explode
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			p := parse(tc.input)
			other := parse(tc.want)
			p.reduce()
			require.True(t, p.equals(other))
		})
	}
}

func Test_add(t *testing.T) {
	for _, tc := range []struct {
		first, second string
		want          []string
	}{
		{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]",
			[]string{
				"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
				"[[[[0,7],4],[15,[0,13]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			},
		},
	} {
		first := parse(tc.first)
		second := parse(tc.second)
		comb := parse(fmt.Sprintf("[%v,%v]", first, second))
		for i, want := range tc.want {
			success := t.Run(fmt.Sprintf("%+v,%v", tc.first, i), func(t *testing.T) {
				other := parse(want)
				fmt.Println(comb)
				fmt.Println(want)
				comb.reduce()
				require.True(t, comb.equals(other))
				require.Equal(t, want, comb.String())
			})
			require.True(t, success)
		}
	}
}
