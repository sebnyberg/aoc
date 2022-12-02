package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func solve1(inf string) string {
	input := "cxdnnyjw"
	var count int
	var res []byte
	h := md5.New()
	var buf []byte
	for x := 0; count < 8; x++ {
		in := input + strconv.Itoa(x)
		h.Reset()
		h.Write([]byte(in))
		buf = h.Sum(buf[:0])
		a := hex.EncodeToString(buf)
		if a[:5] == "00000" {
			res = append(res, a[5])
			count++
		}
	}
	return fmt.Sprint(string(res))
}

func solve2(inf string) string {
	input := "cxdnnyjw"
	var count int
	res := make([]byte, 8)
	h := md5.New()
	var buf []byte
	for x := 0; count < 8; x++ {
		in := input + strconv.Itoa(x)
		h.Reset()
		h.Write([]byte(in))
		buf = h.Sum(buf[:0])
		a := hex.EncodeToString(buf)
		if a[:5] == "00000" {
			idx := int(a[5] - '0')
			if idx < 0 || idx > 7 {
				continue
			}
			char := a[6]
			if res[idx] == 0 {
				res[idx] = char
				count++
			}
		}
	}
	return fmt.Sprint(string(res))
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
