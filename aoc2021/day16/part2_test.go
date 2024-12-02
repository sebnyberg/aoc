package day16

import (
	"encoding/hex"
	"testing"

	"github.com/sebnyberg/aoc/ax"

	"github.com/stretchr/testify/assert"
)

var day16part2res int

func BenchmarkDay16Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day16part2res = Part2(ax.MustReadFileLines("input")[0])
	}
}

func TestDay16Part2(t *testing.T) {
	assert.Equal(t, 3, Part2("C200B40A82"))
	assert.Equal(t, 54, Part2("04005AC33890"))
	assert.Equal(t, 7, Part2("880086C3E88112"))
	assert.Equal(t, 9, Part2("CE00C43D881120"))
	assert.Equal(t, 1, Part2("D8005AC2A8F0"))
	assert.Equal(t, 0, Part2("F600BC2D8F"))
	assert.Equal(t, 0, Part2("9C005AC2F8F0"))
	assert.Equal(t, 1, Part2("9C0141080250320F1802104A08"))
	assert.Equal(t, 180616437720, Part2(ax.MustReadFileLines("input")[0]))
}

func Part2(row string) int {
	hexBytes, err := hex.DecodeString(row)
	ax.Check(err, "failed to decode string")
	bs := hexAsBinary(hexBytes)
	p := parser{
		bs: bs,
	}
	res := p.parsePacket()
	return res
}
