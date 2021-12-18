package day18

import (
	"aoc/ax"
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Part1(t *testing.T) {
	input := ax.MustReadFineLines("input")
	cur := parse(input[0])
	for cur.reduce() {
	}
	for i := 1; i < len(input); i++ {
		other := parse(input[i])
		for other.reduce() {
		}
		combStr := fmt.Sprintf("[%v,%v]", cur, other)
		comb := parse(combStr)
		for comb.reduce() {
			comb = parse(comb.String())
		}
		cur = comb
	}
	require.Equal(t, 4140, cur.magnitude())
}

func Test_Part2(t *testing.T) {
	inputs := ax.MustReadFineLines("input")
	var maxMagnitude int
	for i := 0; i < len(inputs)-1; i++ {
		for j := i + 1; j < len(inputs); j++ {
			comb1 := parse(fmt.Sprintf("[%v,%v]", inputs[i], inputs[j]))
			comb2 := parse(fmt.Sprintf("[%v,%v]", inputs[j], inputs[i]))
			for comb1.reduce() {
			}
			for comb2.reduce() {
			}
			maxMagnitude = ax.Max(maxMagnitude, comb1.magnitude())
			maxMagnitude = ax.Max(maxMagnitude, comb2.magnitude())
		}
	}
	require.Equal(t, 3993, maxMagnitude)
}

func Test_magnitude(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			res := parse(tc.input).magnitude()
			require.Equal(t, tc.want, res)
		})
	}
}

func Test_parseNode(t *testing.T) {
	for _, tc := range []struct {
		input string
	}{
		{"[1,2]"},
		{"[[1,2],3]"},
		{"[9,[8,7]]"},
		{"[[1,9],[8,5]]"},
		{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"},
		{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"},
		{"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			res := parse(tc.input)
			require.Equal(t, tc.input, res.String())
		})
	}
}

func Test_reduce(t *testing.T) {
	for _, tc := range []struct {
		input, want string
	}{
		// Split
		{"[10,1]", "[[5,5],1]"},
		// Explode
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.input), func(t *testing.T) {
			p := parse(tc.input)
			other := parse(tc.want)
			p.reduce()
			require.True(t, p.Equals(other))
		})
	}
}

func Test_add(t *testing.T) {
	for _, tc := range []struct {
		first, second string
		want          []string
	}{
		{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]",
			[]string{
				"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
				"[[[[0,7],4],[15,[0,13]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
				"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			},
		},
	} {
		first := parse(tc.first)
		second := parse(tc.second)
		comb := parse(fmt.Sprintf("[%v,%v]", first, second))
		for i, want := range tc.want {
			success := t.Run(fmt.Sprintf("%+v,%v", tc.first, i), func(t *testing.T) {
				other := parse(want)
				fmt.Println(comb)
				fmt.Println(want)
				comb.reduce()
				require.True(t, comb.Equals(other))
				require.Equal(t, want, comb.String())
			})
			require.True(t, success)
		}
	}
}

func reduceAll(n *snailNode) *snailNode {
	for n.reduce() {
	}
	return n
}

func (n *snailNode) magnitude() int {
	if n.left == nil {
		return n.val
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func (n *snailNode) reduce() bool {
	return n.explodeOne() || n.splitOne()
}

func (n *snailNode) Equals(other *snailNode) bool {
	if n == nil || other == nil {
		return n == nil && other == nil
	}
	if n.level != other.level || n.val != other.val {
		return false
	}
	if !n.left.Equals(other.left) {
		return false
	}
	if !n.right.Equals(other.right) {
		return false
	}
	if n.prev.GetValue() != other.prev.GetValue() ||
		n.next.GetValue() != other.next.GetValue() ||
		n.prev.GetLevel() != other.prev.GetLevel() ||
		n.next.GetLevel() != other.next.GetLevel() {
		fmt.Println("not equal!!")
		fmt.Println(n)
		fmt.Println(other)
		return false
	}
	return true
}

type snailNode struct {
	val   int // only has a value if left and right is nil
	level int
	prev  *snailNode
	next  *snailNode
	left  *snailNode
	right *snailNode
}

func (n *snailNode) IsPair() bool {
	return n != nil && n.left != nil
}

func (n *snailNode) GetValue() int {
	if n == nil {
		return math.MaxInt32
	}
	return n.val
}

func (n *snailNode) GetLevel() int {
	if n == nil {
		return -1
	}
	return n.level
}

func (n *snailNode) String() string {
	if n == nil {
		return ""
	}
	if !n.IsPair() {
		return strconv.Itoa(n.val)
	}
	return fmt.Sprintf("[%v,%v]", n.left, n.right)
}

func (n *snailNode) explodeOne() bool {
	if n == nil {
		return false
	}
	if n.left.explodeOne() {
		return true
	}
	if !n.IsPair() || n.level != 4 {
		return n.right.explodeOne()
	}
	// explode
	prev, p1, p2, next := n.left.prev, n.left, n.right, n.right.next
	if prev != nil {
		prev.val += p1.val
		prev.next = n
	}
	n.prev = prev
	if next != nil {
		next.val += p2.val
		n.next = next
	}
	if next != nil {
		next.prev = n
	}
	n.left = nil
	n.right = nil
	return true
}

func (n *snailNode) splitOne() bool {
	if n == nil {
		return false
	}
	if n.left.splitOne() {
		return true
	}
	if n.val < 10 {
		return n.right.splitOne()
	}
	prev, next := n.prev, n.next
	left := &snailNode{
		val:   n.val / 2,
		level: n.level + 1,
	}
	right := &snailNode{
		val:   n.val/2 + n.val%2,
		level: n.level + 1,
	}
	if prev != nil {
		prev.next = left
	}
	right.next = next
	right.prev = left
	left.next = right
	if next != nil {
		next.prev = right
	}
	left.prev = prev
	n.val = 0
	n.left = left
	n.right = right
	n.prev = nil
	n.next = nil
	return true
}

type parser struct {
	row   string
	pos   int
	nodes []*snailNode
}

func parse(s string) *snailNode {
	p := parser{
		row:   s,
		pos:   0,
		nodes: make([]*snailNode, 0),
	}
	root := p.parseNode(0)
	for i := range p.nodes {
		if i > 0 {
			p.nodes[i].prev = p.nodes[i-1]
		}
		if i < len(p.nodes)-1 {
			p.nodes[i].next = p.nodes[i+1]
		}
	}
	return root
}

func (p *parser) parseNode(level int) *snailNode {
	var n snailNode
	n.level = level
	ch := p.row[p.pos]
	if ch >= '0' && ch <= '9' {
		// this is only needed for split tests
		for p.row[p.pos] >= '0' && p.row[p.pos] <= '9' {
			n.val *= 10
			n.val += int(p.row[p.pos] - '0')
			p.pos++
		}
		p.nodes = append(p.nodes, &n)
		return &n
	}
	// must be '['
	p.pos++
	n.left = p.parseNode(level + 1)
	p.pos++ // ','
	n.right = p.parseNode(level + 1)
	p.pos++ // ']'
	return &n
}
