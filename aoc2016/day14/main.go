package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func solve1(inf string) string {
	const salt = "qzyelonm"
	findTriplet := func(s string) byte {
		n := len(s)
		for i := 0; i < n-2; i++ {
			if s[i] == s[i+1] && s[i+1] == s[i+2] {
				return s[i]
			}
		}
		return 0
	}
	hash := func(s string) string {
		buf := md5.Sum([]byte(s))
		res := hex.EncodeToString(buf[:])
		return res
	}
	gen := func(x int) string {
		return hash(salt + fmt.Sprint(x))
	}
	var kk int
	for x := 1; ; x++ {
		h1 := gen(x)
		b := findTriplet(h1)
		if b == 0 {
			continue
		}
		want := strings.Repeat(string(b), 5)
		for k := 1; k <= 1000; k++ {
			h2 := gen(x + k)
			if strings.Contains(h2, want) {
				// fmt.Println(x)
				kk++
				if kk == 64 {
					return fmt.Sprint(x)
				}
				break
			}
		}
	}

	var res int
	return fmt.Sprint(res)
}

func solve2(inf string) string {
	const salt = "qzyelonm"
	findTriplet := func(s string) byte {
		n := len(s)
		for i := 0; i < n-2; i++ {
			if s[i] == s[i+1] && s[i+1] == s[i+2] {
				return s[i]
			}
		}
		return 0
	}
	mem := make(map[string]string)
	hash := func(s string) string {
		buf := md5.Sum([]byte(s))
		res := hex.EncodeToString(buf[:])
		return res
	}
	superhash := func(s string) string {
		if v, exists := mem[s]; exists {
			return v
		}
		orig := s
		for i := 0; i < 2017; i++ {
			s = hash(s)
		}
		mem[orig] = s
		return s
	}
	var kk int
	for x := 0; ; x++ {
		h1 := superhash(salt + fmt.Sprint(x))
		b := findTriplet(h1)
		if b == 0 {
			continue
		}
		want := strings.Repeat(string(b), 5)
		for k := 1; k <= 1000; k++ {
			h2 := superhash(salt + fmt.Sprint(x+k))
			if strings.Contains(h2, want) {
				kk++
				if kk == 64 {
					return fmt.Sprint(x)
				}
				break
			}
		}
	}

	return ""
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
