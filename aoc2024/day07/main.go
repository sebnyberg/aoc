package main

import (
	"fmt"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var result int
	for _, l := range lines {
		parts := strings.Fields(l)
		want := ax.Atoi(parts[0][:len(parts[0])-1])
		var nums []int
		for _, x := range parts[1:] {
			nums = append(nums, ax.Atoi(x))
		}
		numOps := len(nums) - 1
		for permBits := 0; permBits < (1 << numOps); permBits++ {
			got := nums[0]
			for bit := 0; bit < numOps; bit++ {
				if permBits&(1<<bit) == 0 { // add
					got += nums[bit+1]
				} else { // mul
					got *= nums[bit+1]
				}
			}
			if got == want {
				result += want
				break
			}
		}
	}
	return result
}

func intpow(x, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}

func solve2(inf string) any {
	lines := ax.MustReadFileLines(inf)
	var result int
	for _, l := range lines {
		parts := strings.Fields(l)
		want := ax.Atoi(parts[0][:len(parts[0])-1])
		var nums []int
		for _, x := range parts[1:] {
			nums = append(nums, ax.Atoi(x))
		}
		numOps := len(nums) - 1
		for perm := 0; perm < intpow(3, numOps); perm++ {
			cpy := perm
			got := nums[0]
			for i := 0; i < numOps; i++ {
				switch cpy % 3 {
				case 0:
					got += nums[i+1]
				case 1:
					got *= nums[i+1]
				case 2:
					got = ax.Atoi(fmt.Sprintf("%v%v", got, nums[i+1]))
				}
				if got > want {
					break
				}
				cpy /= 3
			}
			if got == want {
				result += want
				break
			}
		}
	}
	return result
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
