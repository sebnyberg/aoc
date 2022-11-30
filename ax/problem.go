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
type Problem struct {
	Result1  any      // Result to first part of the problem
	Result2  any      // Result to second part of the problem
	Input    []string // Raw input lines
	Parsed   []any    // Parsed results
	HeadN    int      // Print N lines from head of parsed file
	TailN    int      // Print N lines from tail of parsed file
	PrintIdx []int    // Print these input indices (prio over head/tail)
}

func (p Problem) String() string {
	var sb strings.Builder
	if p.Result1 != nil {
		sb.WriteString(fmt.Sprintf("Result1:\n%v\n", p.Result1))
	}
	if p.Result2 != nil {
		sb.WriteString(fmt.Sprintf("Result2:\n%v\n", p.Result1))
	}
	if len(p.Parsed) == 0 {
		sb.WriteString("No parsed lines\n")
		return sb.String()
	}
	if len(p.Parsed) != len(p.Input) {
		d := len(p.Input) - len(p.Parsed)
		s := fmt.Sprintf("Note! Not all lines were parsed, mismatch: %d\n", d)
		sb.WriteString(s)
	}
	headn := p.HeadN
	if headn < 0 {
		headn = math.MaxInt32
	}
	tailn := p.TailN
	if tailn < 0 {
		tailn = math.MaxInt32
	}
	n := len(p.Input)

	tw := tabwriter.NewWriter(&sb, 0, 2, 2, ' ', 0)
	sb.WriteString("")
	fmt.Fprint(tw, "i\tinput\tparsed\n")
	fmt.Fprint(tw, "--------\t-------\t--------\n")

	var j int // position in PrintIdx
	sort.Ints(p.PrintIdx)

	for i := range p.Input {
		doPrint := true
		if i > headn && n-tailn > i {
			// Skip lines depending on tail/head
			doPrint = false
		}
		if j < len(p.PrintIdx) && i == p.PrintIdx[j] {
			// PrintIdx takes priority
			doPrint = true
		}
		if !doPrint {
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
