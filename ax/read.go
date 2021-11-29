package ax

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func MustReadFileLines(path string) chan string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	Check(err, fmt.Sprintf("open file at %q", path))
	return ReadLines(f)
}

func ReadLines(f io.ReadCloser) chan string {
	res := make(chan string)
	go func() {
		defer f.Close()
		defer close(res)
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			res <- sc.Text()
		}
	}()
	return res
}
