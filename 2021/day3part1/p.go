package day3part1

const (
	Problem = 3
	Part    = 1
)

func Run(lines []string) int {
	n := len(lines[0])
	m := len(lines)
	oneCount := make([]int, n)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				oneCount[i]++
			}
		}
	}
	var gamma, eps int
	for _, count := range oneCount {
		gamma <<= 1
		eps <<= 1
		if count*2 > m {
			gamma += 1
		} else {
			eps += 1
		}
	}
	return gamma * eps
}
