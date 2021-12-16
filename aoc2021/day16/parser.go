package day16

const (
	idSum     = 0
	idProduct = 1
	idMin     = 2
	idMax     = 3
	idLit     = 4
	idGt      = 5
	idLt      = 6
	idEq      = 7
)

type bitPacket struct {
	ver, id int
	sub     []bitPacket // if operator
	val     int         // if literal
}

type parser struct {
	bs     []byte
	pos    int
	verSum int
}

func (p *parser) parsePacket() bitPacket {
	var packet bitPacket
	packet.ver, packet.id = p.parseHeader()
	if packet.id == idLit {
		packet.val = p.parseLiteral()
	} else {
		packet.sub = p.parseOperator()
	}
	return packet
}

func (p *parser) readInt(n int) int {
	var res int
	for i := 0; i < n; i++ {
		res <<= 1
		res += int(p.bs[p.pos+i] % 2)
	}
	p.pos += n
	return res
}

func (p *parser) readBits(n int) []byte {
	res := p.bs[p.pos : p.pos+n]
	p.pos += n
	return res
}

func (p *parser) parseHeader() (int, int) {
	ver := p.readInt(3)
	p.verSum += ver
	id := p.readInt(3)
	return ver, id
}

func (p *parser) parseLiteral() int {
	var res int
	for {
		res <<= 4
		hasMore := p.readInt(1) == 1
		res += p.readInt(4)
		if !hasMore {
			return res
		}
	}
}

func (p *parser) parseOperator() []bitPacket {
	var res []bitPacket
	if p.readInt(1) == 0 {
		width := p.readInt(15)
		start := p.pos
		for p.pos-start < width {
			res = append(res, p.parsePacket())
		}
	} else {
		n := p.readInt(11)
		for i := 0; i < n; i++ {
			res = append(res, p.parsePacket())
		}
	}
	return res
}

func hexAsBinary(hexBytes []byte) []byte {
	var res []byte
	for _, hexByte := range hexBytes {
		for i := 0; i < 8; i++ {
			res = append(res, hexByte&(1<<7)>>7+'0')
			hexByte <<= 1
		}
	}
	return res
}

func (p bitPacket) eval() int {
	if p.id == idLit {
		return p.val
	}

	// Collect values and perform op
	first := p.sub[0].eval()
	res := p.sub[0].eval()
	for _, sub := range p.sub[1:] {
		switch p.id {
		case idSum:
			res += sub.eval()
		case idProduct:
			res *= sub.eval()
		case idMin:
			res = min(res, sub.eval())
		case idMax:
			res = max(res, sub.eval())
		case idGt:
			if first > sub.eval() {
				res = 1
			} else {
				res = 0
			}
		case idLt:
			if first < sub.eval() {
				res = 1
			} else {
				res = 0
			}
		case idEq:
			if first == sub.eval() {
				res = 1
			} else {
				res = 0
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
