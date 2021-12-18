package day18

import (
	"fmt"
	"math"
	"strconv"
)

type snailNode struct {
	val   int // only has a value if left and right is nil
	level int
	prev  *snailNode
	next  *snailNode
	left  *snailNode
	right *snailNode
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

func (n *snailNode) equals(other *snailNode) bool {
	if n == nil || other == nil {
		return n == nil && other == nil
	}
	if n.level != other.level || n.val != other.val {
		return false
	}
	if !n.left.equals(other.left) {
		return false
	}
	if !n.right.equals(other.right) {
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
