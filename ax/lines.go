package ax

// LinesChan reads from a slice of strings into the returned channel.
func LinesChan(ss []string) chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		for _, s := range ss {
			lines <- s
		}
	}()
	return lines
}

// LinesChan reads all strings from a channel into a slice of strings.
func LinesFromChan(lines chan string) []string {
	res := make([]string, 0, 10)
	for line := range lines {
		res = append(res, line)
	}
	return res
}
