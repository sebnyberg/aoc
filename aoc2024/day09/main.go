package main

import (
	"fmt"

	"github.com/sebnyberg/aoc/ax"
)

func solve1(inf string) any {
	line := ax.MustReadFileLines(inf)[0]
	disk := make([]int, 0, len(line))
	for i := range line {
		status := -1
		if i%2 == 0 {
			status = i / 2
		}
		for k := 0; k < int(line[i]-'0'); k++ {
			disk = append(disk, status)
		}
	}
	l := int(line[0] - '0')
	r := len(disk) - 1
	var res int
	for l < len(disk) {
		if r > l && disk[l] == -1 {
			disk[l] = disk[r]
			disk[r] = -1
			for disk[r] == -1 {
				r--
			}
		}
		if disk[l] != -1 {
			res += l * disk[l]
		}
		l++
	}
	return res
}

func solve2(inf string) any {
	line := ax.MustReadFileLines(inf)[0]
	type file struct {
		id       int16
		sz       int8
		blockIdx int16
	}
	type block struct {
		free      int8
		size      int8
		filesIdxs []int16
	}
	var files []file
	var disk []block
	for i := range line {
		sz := int8(line[i]) - '0'
		if i%2 == 0 {
			files = append(files, file{id: int16(i / 2), sz: sz})
			disk = append(disk, block{
				free:      0,
				size:      sz,
				filesIdxs: []int16{int16(i / 2)},
			})
		} else {
			disk = append(disk, block{
				free:      sz,
				size:      sz,
				filesIdxs: []int16{},
			})
		}
	}
	for r := len(disk) - len(disk)&1; r > 0; r -= 2 {
		for l := 1; l < r; l += 2 {
			if disk[l].free >= disk[r].size {
				disk[l].free -= disk[r].size
				disk[l].filesIdxs = append(disk[l].filesIdxs, disk[r].filesIdxs...)
				disk[r].filesIdxs = disk[r].filesIdxs[:0]
				disk[r].free = disk[r].size
				break
			}
		}
	}
	var res int
	var pos int
	for _, block := range disk {
		for _, f := range block.filesIdxs {
			for k := 0; k < int(files[f].sz); k++ {
				res += int(files[f].id) * pos
				pos++
			}
		}
		pos += int(block.free)
	}
	return res
}

func main() {
	f := "input"
	fmt.Printf("Result1:\n%v\n", solve1(f))
	fmt.Printf("Result2:\n%v\n", solve2(f))
}
