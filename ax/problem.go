package ax

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"text/tabwriter"
)

// Problem contains metadata from an AoC problem. It is used primarily to
// continuously watch results and debug parsing logic from the terminal:
//
// $ watch -n 0.1 "go run main.go < input"
type Problem[T1, T2, T3 comparable] struct {
	Input    []string // Raw input lines
	Parsed   []T1     // Parsed results
	HeadN    int      // Print N lines from head of parsed file
	TailN    int      // Print N lines from tail of parsed file
	PrintIdx []int    // Print these input indices (prio over head/tail)
	Result1  T2       // Result to first part of the problem
	Result2  T3       // Result to second part of the problem
}

func (p Problem[T1, T2, T3]) String() string {
	var sb strings.Builder

	// Print result 1
	res1 := fmt.Sprintf("%v (default value)", *new(T2))
	if p.Result1 != *new(T2) {
		res1 = fmt.Sprint(p.Result1)
	}
	sb.WriteString(fmt.Sprintf("Result1:\n%v\n\n", res1))

	// Print result 2
	res2 := fmt.Sprintf("%v (default value)", *new(T3))
	if p.Result2 != *new(T3) {
		res2 = fmt.Sprint(p.Result2)
	}
	sb.WriteString(fmt.Sprintf("Result2:\n%v\n\n", res2))
	if len(p.Parsed) == 0 {
		sb.WriteString("No parsed lines\n")
		return sb.String()
	}

	// Check if all lines were parsed
	if len(p.Parsed) != len(p.Input) {
		d := len(p.Input) - len(p.Parsed)
		s := fmt.Sprintf("Note! Not all lines were parsed, mismatch: %d\n", d)
		sb.WriteString(s)
	}

	// Adjust head/tail in the case of -1
	headn := p.HeadN
	if headn < 0 {
		headn = math.MaxInt32
	}
	tailn := p.TailN
	if tailn < 0 {
		tailn = math.MaxInt32
	}

	tw := tabwriter.NewWriter(&sb, 0, 2, 2, ' ', 0)
	sb.WriteString("")
	fmt.Fprint(tw, "i\tinput\tparsed\n")
	fmt.Fprint(tw, "--------\t-------\t--------\n")

	var j int // position in PrintIdx
	sort.Ints(p.PrintIdx)

	n := len(p.Input)
	for i := range p.Input {
		var skip bool

		// Skip lines depending on tail/head
		if i > headn && n-tailn > i {
			skip = true
		}

		// Always print printIndices
		if j < len(p.PrintIdx) &&
			(p.PrintIdx[j] < 0 || p.PrintIdx[j] >= n) {
			// Unless out of bounds
			fmt.Fprintf(tw, "%v\terr:invalid idx\t\n", p.PrintIdx[j])
			break
		}
		if j < len(p.PrintIdx) && i == p.PrintIdx[j] {
			skip = false
		}

		if skip {
			continue
		}

		in := []byte(p.Input[i])
		if len(in) > 36 {
			// Truncate if too long
			in = append(in[:36], "..."...)
		}
		out := []byte("<nil>")
		if i < len(p.Parsed) {
			out = []byte(fmt.Sprint(p.Parsed[i]))
			// Truncate if too long
			if len(out) > 36 {
				out = append(out[:36], "..."...)
			}
		}
		fmt.Fprintf(tw, "%v\t%s\t%s\n", i, in, out)
	}
	tw.Flush()
	return sb.String()
}
