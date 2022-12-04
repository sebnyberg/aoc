package main

import (
	"fmt"
)

func solve1() string {
	for x := 1; ; x++ {
		if (x+1)%5 != 3 {
			continue
		}
		if (x+2)%13 != 6 {
			continue
		}
		if (x+3)%17 != 7 {
			continue
		}
		if (x+4)%3 != 1 {
			continue
		}
		if (x+5)%19 != 10 {
			continue
		}
		if (x+6)%7 != 0 {
			continue
		}
		return fmt.Sprint(x)
	}
}

func solve2() string {
	for x := 1; ; x++ {
		if (x+1)%5 != 3 {
			continue
		}
		if (x+2)%13 != 6 {
			continue
		}
		if (x+3)%17 != 7 {
			continue
		}
		if (x+4)%3 != 1 {
			continue
		}
		if (x+5)%19 != 10 {
			continue
		}
		if (x+6)%7 != 0 {
			continue
		}
		if (x+7)%11 != 0 {
			continue
		}
		return fmt.Sprint(x)
	}
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve1())
	fmt.Printf("Result2:\n%v\n\n", solve2())
}
