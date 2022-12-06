package main

import (
	"fmt"
	"strconv"
)

func isnum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func run(a int) bool {
	// I couldn't figure out how to run this properly so I did it by hand.
	// We start with the value b = 182*14+a
	// Then, while b >= 0, we take res := b&1 and b /= 2
	// If an a can be chosen such tnat the total number of iterations is even
	// and res alternates each round, then we've found an answer.
	//
	b := 182*14 + a
	res := b & 1
	n := 1
	for b > 0 {
		b /= 2
		if b&1 != 1-res {
			return false
		}
		res = b & 1
	}
	return n%2 == 0
}

func solve() string {
	for a := 0; ; a++ {
		if run(a) {
			return fmt.Sprint(a)
		}
	}
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve())
}
