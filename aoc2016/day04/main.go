package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) string {
	var res int
	lines := ax.MustReadFileLines(inf)
	// lines = []string{
	// 	"aaaaa-bbb-z-y-x-123[abxyz]",
	// 	"a-b-c-d-e-f-g-h-987[abcde]",
	// 	"not-a-real-room-404[oarel]",
	// 	"totally-real-room-200[decoy]",
	// }
	for _, l := range lines {
		parts := strings.Split(l, "-")
		m := len(parts)
		letterCounts := make([]int, 26)
		for i := 0; i < m-1; i++ {
			for _, ch := range parts[i] {
				letterCounts[ch-'a']++
			}
		}
		idx := make([]int, 26)
		for i := range idx {
			idx[i] = i
		}
		sort.Slice(idx, func(i, j int) bool {
			if letterCounts[idx[i]] == letterCounts[idx[j]] {
				return idx[i] < idx[j]
			}
			return letterCounts[idx[i]] > letterCounts[idx[j]]
		})
		parts2 := strings.Split(parts[m-1], "[")
		checksum := parts2[1][:len(parts2[1])-1]
		sectorID := ax.Atoi(parts2[0])
		for i, ch := range checksum {
			if rune(idx[i]+'a') != ch {
				// fmt.Println("invalid checksum")
				goto invalid
			}
		}
		res += sectorID
	invalid:
	}
	return fmt.Sprint(res)
}

func solve2(inf string) string {
	var res int
	lines := ax.MustReadFileLines(inf)
	for _, l := range lines {
		parts := strings.Split(l, "-")
		m := len(parts)
		letterCounts := make([]int, 26)
		for i := 0; i < m-1; i++ {
			for _, ch := range parts[i] {
				letterCounts[ch-'a']++
			}
		}
		idx := make([]int, 26)
		for i := range idx {
			idx[i] = i
		}
		sort.Slice(idx, func(i, j int) bool {
			if letterCounts[idx[i]] == letterCounts[idx[j]] {
				return idx[i] < idx[j]
			}
			return letterCounts[idx[i]] > letterCounts[idx[j]]
		})
		lastParts := strings.Split(parts[m-1], "[")
		checksum := lastParts[1][:len(lastParts[1])-1]
		sectorID := ax.Atoi(lastParts[0])
		var ss []string
		for i, ch := range checksum {
			if rune(idx[i]+'a') != ch {
				goto invalid
			}
		}
		for _, part := range parts[:m-1] {
			v := []byte(part)
			for j := range part {
				v[j] = byte((((int(part[j]) - 'a') + sectorID) % 26) + 'a')
			}
			ss = append(ss, string(v))
		}
		if strings.Join(ss, " ") == "northpole object storage" {
			return fmt.Sprint(sectorID)
		}
	invalid:
	}
	return fmt.Sprint(res)
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n\n", solve2(f))
}
