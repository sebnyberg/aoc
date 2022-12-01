package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(in *input) string {
	var res int
	xs := in.xs
	var sum int
	n := len(in.xs)
	nums := make([]int, n)
	for i, x := range xs {
		nums[i] = x.a
		sum += x.a
	}

	want := sum / 3

	mem := make(map[[2]int]bool)
	sort.Ints(nums)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	var f finder
	f.minK = math.MaxInt32
	f.minProd = math.MaxInt32

	f.dfs(mem, nums, 0, 0, n, 0, 1, want, 0, 3)
	fmt.Println(f)
	return fmt.Sprint(res)
}

func solve2(in *input) string {
	var res int
	xs := in.xs
	var sum int
	n := len(in.xs)
	nums := make([]int, n)
	for i, x := range xs {
		nums[i] = x.a
		sum += x.a
	}

	want := sum / 4

	mem := make(map[[2]int]bool)
	sort.Ints(nums)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	var f finder
	f.minK = math.MaxInt32
	f.minProd = math.MaxInt32

	f.dfs(mem, nums, 0, 0, n, 0, 1, want, 0, 3)
	fmt.Println(f)
	return fmt.Sprint(res)
}

type finder struct {
	minK    int
	minProd int
}

func (f *finder) dfs(mem map[[2]int]bool, nums []int, bm, i, n, k, prod, want, sum, nparts int) {
	if sum > want {
		return
	}
	if k > f.minK {
		return
	}
	if k == f.minK && prod >= f.minProd {
		return
	}
	if sum == want {
		if canSplitIntoParts(mem, nums, bm, 0, n, want, want, nparts-2) {
			f.minK = k
			f.minProd = prod
		}
		return
	}
	if i == n {
		return
	}
	// Try to add and not add the current package to the center
	f.dfs(mem, nums, bm|(1<<i), i+1, n, k+1, prod*nums[i], want, sum+nums[i], nparts)
	f.dfs(mem, nums, bm, i+1, n, k, prod, want, sum, nparts)
}

func canSplitIntoParts(mem map[[2]int]bool, nums []int, bm, i, n, sum, want, k int) bool {
	key := [2]int{bm, k}
	if v, ok := mem[key]; ok {
		return v
	}
	if sum < 0 {
		return false
	}
	if sum == 0 {
		if k == 1 {
			return true
		}
		return canSplitIntoParts(mem, nums, bm, 0, n, want, want, 1)
	}
	if i == n {
		return false
	}
	// Don't pick
	if canSplitIntoParts(mem, nums, bm, i+1, n, sum, want, k) {
		mem[key] = true
		return true
	}

	// Pick
	if bm&(1<<i) == 0 {
		if canSplitIntoParts(mem, nums, bm|(1<<i), i+1, n, sum-nums[i], want, k) {
			mem[key] = true
			return true
		}
	}

	mem[key] = false
	return false
}

type inputItem struct {
	a int
}

type input struct {
	n  int
	xs []inputItem
}

var pat = regexp.MustCompile(``)

func (p *input) parse(s string) {
	var x inputItem
	x.a = ax.Atoi(s)
	p.xs = append(p.xs, x)
	p.n++
}

func main() {
	in := new(input)
	f, _ := os.Open("input")
	rows := ax.ReadLines(f)
	for _, s := range rows {
		in.parse(s)
	}
	fmt.Printf("Result1:\n%v\n", solve1(in))
	fmt.Printf("Result2:\n%v\n\n", solve2(in))
	fmt.Printf("Input:\n%v\n", ax.Debug(rows, 1))
	fmt.Printf("Parsed:\n%v\n", ax.Debug(in.xs, 1))
}
