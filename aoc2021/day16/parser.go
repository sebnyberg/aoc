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

func (p *parser) parsePacket() int {
	_, id := p.parseHeader()
	if id == idLit {
		return p.parseLiteral()
	}
	vals := p.parseOperator()
	return eval(id, vals)
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

func (p *parser) parseOperator() []int {
	var res []int
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
	res := make([]byte, 0, len(hexBytes)*4)
	for _, hexByte := range hexBytes {
		for i := 0; i < 8; i++ {
			res = append(res, hexByte&(1<<7)>>7+'0')
			hexByte <<= 1
		}
	}
	return res
}

func eval(id int, vals []int) int {
	// Collect values and perform op
	switch id {
	case idSum:
		for i := 1; i < len(vals); i++ {
			vals[0] += vals[i]
		}
	case idProduct:
		for i := 1; i < len(vals); i++ {
			vals[0] *= vals[i]
		}
	case idMin:
		for i := 1; i < len(vals); i++ {
			vals[0] = min(vals[0], vals[i])
		}
	case idMax:
		for i := 1; i < len(vals); i++ {
			vals[0] = max(vals[0], vals[i])
		}
	case idGt:
		if vals[0] > vals[1] {
			return 1
		}
		return 0
	case idLt:
		if vals[0] < vals[1] {
			return 1
		}
		return 0
	case idEq:
		if vals[0] == vals[1] {
			return 1
		}
		return 0
	}
	return vals[0]
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
