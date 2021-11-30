package ax

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func MustReadFineLinesChan(path string) chan string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	Check(err, fmt.Sprintf("open file at %q", path))
	return ReadLinesChan(f)
}

func ReadLinesChan(f io.ReadCloser) chan string {
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

func MustReadFineLines(path string) []string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	Check(err, fmt.Sprintf("open file at %q", path))
	return ReadLines(f)
}

func ReadLines(f io.ReadCloser) []string {
	res := make([]string, 0)
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res
}
