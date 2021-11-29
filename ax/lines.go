package ax

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
