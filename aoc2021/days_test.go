package aoc2021

import (
	"aoc/ax"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type solver struct {
	f     func([]string) string
	wants []want
}

type want struct {
	fname string
	val   string
}

var solvers = []solver{
	// {Day01Part1, []want{{"small", "7"}, {"input", "1292"}}},
	// {Day01Part2, []want{{"small", "5"}, {"input", "1262"}}},
	// {Day02Part1, []want{{"small", "150"}, {"input", "1524750"}}},
	// {Day02Part2, []want{{"small", "900"}, {"input", "1592426537"}}},
	// {Day03Part1, []want{{"small", "198"}, {"input", "775304"}}},
	// {Day03Part2, []want{{"small", "198"}, {"input", "775304"}}},
	// {Day04Part1, []want{{"small", "4512"}, {"input", "65325"}}},
	// {Day04Part2, []want{{"small", "1924"}, {"input", "4624"}}},
	// {Day05Part1, []want{{"small", "5"}, {"input", "5092"}}},
	// {Day05Part2, []want{{"small", "12"}, {"input", "20484"}}},
	// {Day06Part1, []want{{"small", "5934"}, {"input", "395627"}}},
	// {Day06Part2, []want{{"small", "26984457539"}, {"input", "1767323539209"}}},
	// {Day07Part1, []want{{"small", "37"}, {"input", "352997"}}},
	// {Day07Part2, []want{{"small", "168"}, {"input", "101571302"}}},
	// {Day08Part1, []want{{"small", "26"}, {"input", "375"}}},
	// {Day08Part2, []want{{"small", "61229"}, {"input", "1019355"}}},
	// {Day09Part1, []want{{"small", "15"}, {"input", "425"}}},
	// {Day09Part2, []want{{"small", "1134"}, {"input", "1135260"}}},
	// {Day10Part1, []want{{"small", "26397"}, {"input", "366027"}}},
	// {Day10Part2, []want{{"small", "288957"}, {"input", "1118645287"}}},
	// {Day11Part1, []want{{"small", "1656"}, {"input", "1562"}}},
	// {Day11Part2, []want{{"small", "195"}, {"input", "268"}}},
	// {Day12Part1, []want{{"small", "10"}, {"input", "268"}}},
	{Day12Part2, []want{{"small", "36"}, {"input", "122880"}}},
}

func TestAll(t *testing.T) {
	for _, solver := range solvers {
		name := path.Ext(funcName(solver.f))[1:]
		dayPart := strings.ToLower(name[:5])
		t.Run(name, func(t *testing.T) {
			for _, w := range solver.wants {
				path := fmt.Sprintf("testdata/%v/%v", dayPart, w.fname)
				input := ax.MustReadFineLines(path)
				require.Equal(t, w.val, solver.f(input))
			}
		})
	}
}

func BenchmarkParts(b *testing.B) {
	var res string
	for _, solver := range solvers {
		name := path.Ext(funcName(solver.f))[1:]
		dayPart := strings.ToLower(name[:5])
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				path := fmt.Sprintf("testdata/%v/input", dayPart)
				input := ax.MustReadFineLines(path)
				res = solver.f(input)
			}
		})
	}
	_ = res
}

func funcName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
