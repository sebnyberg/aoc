package ax

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// MustReadFileLinesChan reads lines from the provided file and puts them into
// the returned channel. If there is an error, it is logged followed by
// os.Exit(1).
func MustReadFineLinesChan(path string) chan string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	Check(err, fmt.Sprintf("open file at %q", path))
	return ReadLinesChan(f)
}

// ReadLinesChan reads lines from the provided io.ReadCloser and puts them into
// the returned channel.
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

// ReadFileLines reads lines from the provided file and puts them into a slice
// of strings. If there is an error, it is logged followed by os.Exit(1).
func MustReadFileLines(path string) []string {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	Check(err, fmt.Sprintf("open file at %q", path))
	return ReadLines(f)
}

// ReadLines reads lines from the provided io.ReadCloser, putting them into a
// slice of strings.
func ReadLines(f io.ReadCloser) []string {
	res := make([]string, 0)
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res
}
