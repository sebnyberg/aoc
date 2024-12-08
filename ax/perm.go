package ax

// IntPerm generates all permutations of n integers in the range [0, x) and
// calls func(xs []int) for each permutation. To stop iteration, func must
// return false.
func IntPerm(x, n int, fn func(xs []int) (ok bool)) {
	xs := make([]int, n)
	permDFS(xs, 0, x, n, fn)
}

func permDFS(xs []int, i, x, n int, fn func(xs []int) (ok bool)) bool {
	if i == n {
		return fn(xs)
	}
	for j := 0; j < x; j++ {
		xs[i] = j
		if !permDFS(xs, i+1, x, n, fn) {
			return false
		}
	}
	return true
}
