package day11part2

const (
	Problem = 11
	Part    = 2
)

var emptyboard [10][10]byte

func run(rows []string) int {
	type pos struct{ i, j int }
	ok := func(p pos) bool {
		return p.i >= 0 && p.i < 10 && p.j >= 0 && p.j < 10
	}
	var board [10][10]byte
	for i, row := range rows {
		for j := range row {
			board[i][j] = byte(row[j] - '0')
		}
	}
	willFlash := make([]pos, 0, 100)
	var sum int
	var steps int
	for {
		steps++
		willFlash = willFlash[:0]

		// Collect 9's
		var seen [10][10]bool
		for i := range board {
			for j := range board[i] {
				if board[i][j] != 9 {
					continue
				}
				willFlash = append(willFlash, pos{i, j})
				board[i][j]++
			}
		}
		// While there are places to visit
		for i := 0; i < len(willFlash); i++ {
			p := willFlash[i]
			seen[p.i][p.j] = true
			for _, q := range []pos{
				{p.i + 1, p.j}, {p.i - 1, p.j}, {p.i, p.j + 1}, {p.i, p.j - 1},
				{p.i - 1, p.j - 1}, {p.i - 1, p.j + 1}, {p.i + 1, p.j - 1}, {p.i + 1, p.j + 1},
			} {
				if !ok(q) {
					continue
				}
				board[q.i][q.j]++
				if board[q.i][q.j] == 9 {
					willFlash = append(willFlash, q)
				}
			}
		}
		sum += len(willFlash)
		for i := range board {
			for j := range board[i] {
				if board[i][j] >= 9 {
					board[i][j] = 0
				} else {
					board[i][j]++
				}
			}
		}
		if board == emptyboard {
			return steps
		}
	}
}
