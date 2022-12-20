package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve(f, 1, 1))
	fmt.Printf("Result2:\n%v\n", solve(f, 10, 811589153))
}

func solve(fname string, rounds, key int) any {
	type numPos struct {
		v, origPos int
	}

	var ordered []*numPos
	var nums []*numPos
	var first *numPos

	for i, l := range ax.MustReadFileLines(fname) {
		x := &numPos{
			v:       ax.Atoi(l) * key,
			origPos: i,
		}
		nums = append(nums, x)
		ordered = append(ordered, x)
		if x.v == 0 {
			first = x
		}
	}

	n := len(nums)

	for round := 0; round < rounds; round++ {
		for _, x := range ordered {
			var j int
			for nums[j] != x {
				j++
			}

			// Cut number from sequence
			copy(nums[j:], nums[j+1:])

			// Make room at target position
			target := mod(j+x.v, n-1)
			copy(nums[target+1:], nums[target:])

			// Insert
			nums[target] = x
		}
	}

	var j int
	for nums[j] != first {
		j++
	}
	var res int
	for k := 1000; k <= 3000; k += 1000 {
		res += nums[(j+k)%n].v
	}
	return res
}

func mod(x, m int) int {
	if x < 0 {
		mm := ((-x) / m) + 1
		x += mm * m
	}
	return x % m
}
