package day01

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	floor, basementIdx := solve()
	require.Equal(t, 232, floor)
	require.Equal(t, 1783, basementIdx)
}

func solve() (int, int) {
	var floor int
	path := "input"
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("open file %v failed, %v\n", path, err)
	}
	sc := bufio.NewScanner(f)
	basementIdx := -1
	for sc.Scan() {
		line := sc.Text()
		for i, ch := range line {
			switch ch {
			case '(':
				floor++
			case ')':
				if floor == 0 && basementIdx == -1 {
					basementIdx = i + 1
				}
				floor--
			}
		}
	}
	return floor, basementIdx
}
