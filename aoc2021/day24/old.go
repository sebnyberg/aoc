package day24

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/sebnyberg/aoc/ax"

	"go.uber.org/zap"
)

type evalFn func(stack) stack

type stack struct {
	vals     [4]int
	input    string
	inputPos int
}

// Keeping this old code for reference.. It does not work.
func old(rows []string) int {
	log, _ := zap.NewDevelopment()
	evals := compile(rows)
	type token struct{}
	poolCh := make(chan token, 100)
	var count uint64
	var result string
	var resultOnce sync.Once
	// Using errgroup to signal completion through context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := 99999999999999; v >= 11111111111111; v-- {
		valStr := strconv.Itoa(v)
		if strings.ContainsAny(valStr, "0") { // skip with zeroes
			continue
		}
		select {
		case <-ctx.Done():
			return ax.MustParseInt[int](result)
		case poolCh <- token{}:
		}

		go func() error {
			atomic.AddUint64(&count, 1)
			if c := atomic.LoadUint64(&count); c%1e6 == 0 {
				log.Info("checkin", zap.String("val", valStr))
			}
			defer func() { <-poolCh }() // release slot
			// Run number through program
			p := stack{
				input: valStr,
			}
			for segIdx := range evals {
				for i := range evals[segIdx] {
					p = evals[segIdx][i](p)
				}
			}
			if p.vals[3] == 0 {
				resultOnce.Do(func() {
					result = valStr
					cancel() // Cancel other goroutines
				})
			}
			return nil
		}()
	}
	return 0
}

// compile returns eval functions and the indices where there is an input
func compile(rows []string) [][]evalFn {
	evals := make([][]evalFn, 0, 250)
	k := -1
	pat := regexp.MustCompile(`^(\w+)\s([-a-z0-9]+)(\s([-a-z0-9]+))?$`)
	for i, row := range rows {
		parts := pat.FindStringSubmatch(row)
		first := parts[2][0] - 'w'
		second := -1
		_ = i
		if len(parts[4]) > 0 {
			if strings.ContainsRune("wxyz", rune(parts[4][0])) {
				second = int(parts[4][0] - 'w')
			}
		}
		switch parts[1] {
		case "inp":
			k++
			evals = append(evals, make([]evalFn, 0))
			evals[k] = append(evals[k], func(p stack) stack {
				p.vals[first] = int(p.input[p.inputPos] - '0')
				p.inputPos++
				return p
			})
		case "mul":
			if second == -1 {
				b := ax.MustParseInt[int](parts[4])
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] *= b
					return p
				})
			} else {
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] *= p.vals[second]
					return p
				})
			}
		case "add":
			if second == -1 {
				b := ax.MustParseInt[int](parts[4])
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] += b
					return p
				})
			} else {
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] += p.vals[second]
					return p
				})
			}
		case "mod":
			if second == -1 {
				b := ax.MustParseInt[int](parts[4])
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] %= b
					return p
				})
			} else {
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] %= p.vals[second]
					return p
				})
			}
		case "div":
			if second == -1 {
				b := ax.MustParseInt[int](parts[4])
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] /= b
					return p
				})
			} else {
				evals[k] = append(evals[k], func(p stack) stack {
					p.vals[first] /= p.vals[second]
					return p
				})
			}
		case "eql":
			if second == -1 {
				b := ax.MustParseInt[int](parts[4])
				evals[k] = append(evals[k], func(p stack) stack {
					if p.vals[first] == b {
						p.vals[first] = 1
					} else {
						p.vals[first] = 0
					}
					return p
				})
			} else {
				evals[k] = append(evals[k], func(p stack) stack {
					if p.vals[first] == p.vals[second] {
						p.vals[first] = 1
					} else {
						p.vals[first] = 0
					}
					return p
				})
			}
		}
	}
	return evals
}
