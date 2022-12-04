package main

import (
	"fmt"
)

func solve(n int) string {
	rev := func(s string) string {
		bs := []byte(s)
		for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
			bs[l], bs[r] = bs[r], bs[l]
		}
		return string(bs)
	}
	inv := func(s string) string {
		bs := []byte(s)
		for i := range s {
			bs[i] = '0' + (1 - (bs[i] - '0'))
		}
		return string(bs)
	}
	dragon := func(s string) string {
		return s + "0" + inv(rev(s))
	}
	checksum := func(s string) string {
		var res []byte
		for i := 0; i < len(s); i += 2 {
			if s[i] == s[i+1] {
				res = append(res, '1')
			} else {
				res = append(res, '0')
			}
		}
		return string(res)
	}
	val := "01111010110010011"
	for len(val) < n {
		val = dragon(val)
	}
	val = checksum(val[:n])
	for len(val)%2 == 0 {
		val = checksum(val)
	}
	return fmt.Sprint(val)
}

func main() {
	fmt.Printf("Result1:\n%v\n", solve(272))
	fmt.Printf("Result2:\n%v\n\n", solve(35651584))
}
