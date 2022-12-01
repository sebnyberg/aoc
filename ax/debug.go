package ax

import (
	"fmt"
	"strings"
)

func Debug[T any](rows []T, n int) string {
	var sb strings.Builder
	m := len(rows)
	for i, row := range Head(rows, n) {
		if s, ok := any(row).(fmt.Stringer); ok {
			fmt.Fprintf(&sb, "%v:\t%s\n", i, s)
		} else {
			fmt.Fprintf(&sb, "%v:\t%+v\n", i, row)
		}
	}
	for i, row := range Tail(rows, n) {
		if s, ok := any(row).(fmt.Stringer); ok {
			fmt.Fprintf(&sb, "%v:\t%s\n", m-i-1, s)
		} else {
			fmt.Fprintf(&sb, "%v:\t%+v\n", m-i-1, row)
		}
	}
	return sb.String()
}
