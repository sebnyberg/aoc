package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	m := make(map[int]map[string][]string)
	var maxk int
	for _, mapping := range in.mappings {
		k := len(mapping.from)
		if _, exists := m[k]; !exists {
			m[k] = make(map[string][]string)
		}
		m[k][mapping.from] = append(m[k][mapping.from], mapping.to)
		maxk = ax.Max(maxk, k)
	}
	pat := in.replacement
	seen := make(map[string]struct{})
	for k := 1; k <= maxk; k++ {
		if _, exists := m[k]; !exists {
			continue
		}
		for i := 0; i < len(pat)-k+1; i++ {
			s := pat[i : i+k]
			l := pat[:i]
			r := pat[i+k:]
			for _, ss := range m[k][s] {
				sss := l + ss + r
				seen[sss] = struct{}{}
			}
		}
	}
	return fmt.Sprint(len(seen))
}

func solve2(in *input) string {
	maps := in.mappings
	var maxLen int
	m := make(map[string]string)
	for i := range maps {
		if _, exists := m[maps[i].to]; exists {
			panic("equal mappings")
		}
		m[maps[i].to] = maps[i].from
		maxLen = ax.Max(maxLen, len(maps[i].to))
	}

	transform := []string{
		"Th", "D",
		"Al", "E",
		"Rn", "(",
		"Ar", ")",
		"Y", ",",
		"Mg", "J",
		"Ti", "K",
		"Si", "L",
		"Ca", "Q",
	}
	r := strings.NewReplacer(transform...)
	f, _ := os.OpenFile("input1", os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	for i := range maps {
		maps[i].from = r.Replace(maps[i].from)
		maps[i].to = r.Replace(maps[i].to)
	}
	repl := r.Replace(in.replacement)

	sort.Slice(maps, func(i, j int) bool {
		return maps[i].to < maps[j].to
	})
	for i := range maps {
		fmt.Fprintf(f, "%s => %s\n", maps[i].from, maps[i].to)
	}
	fmt.Fprintf(f, "\n%s\n", repl)
	f.Close()

	// After reducing the double-character symbols, the problem becomes clearer.
	//
	// There exists two categories of productions:
	//
	// 1: X => XX
	// 2: X => {X(X), X(X,X), X(X,X,X)}
	//
	// There is no way to apply a production of (1) or (2) such that the number
	// of commas at a certain level of parenthesis would be any different from
	// what was added by the wrapping (2).
	//
	// This tells us that a) there is only a single solution to the problem, and
	// b) commas incur a reduction of one extra for free, parenthesis are free.
	//
	pat1 := regexp.MustCompile(`[\(\)]`)
	pat2 := regexp.MustCompile(`,`)
	aa := pat1.FindAllString(repl, -1)
	bb := pat2.FindAllString(repl, -1)
	res := len(repl) - len(aa) - 2*len(bb)
	return fmt.Sprint(res)
}

type mapping struct {
	from string
	to   string
}

type input struct {
	n                int
	raw              string
	mappings         []mapping
	replacement      string
	nextIsRepacement bool
}

var pat = regexp.MustCompile(``)

func (p *input) parse(s string) {
	if s == "" {
		p.nextIsRepacement = true
		return
	}
	if p.nextIsRepacement {
		p.replacement = s
		return
	}
	pat := regexp.MustCompile(`(\w+) => (\w+)`)
	ss := pat.FindStringSubmatch(s)
	var x mapping
	x.from = ss[1]
	x.to = ss[2]
	p.mappings = append(p.mappings, x)
	p.n++
}

func main() {
	in := new(input)
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	rows := ax.ReadLines(f)
	in.raw = strings.Join(rows, "\n")
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n", solve2(in))
}
