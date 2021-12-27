package day16

import (
	"aoc/ax"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day16part1 int

func BenchmarkDay16Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day16part1 = Part1(ax.MustReadFileLines("input")[0])
	}
}

func TestDay16Part1(t *testing.T) {
	assert.Equal(t, 16, Part1("8A004A801A8002F478"))
	assert.Equal(t, 12, Part1("620080001611562C8802118E34"))
	assert.Equal(t, 23, Part1("C0015000016115A2E0802F182340"))
	assert.Equal(t, 31, Part1("A0016C880162017C3686B18A3D4780"))
	assert.Equal(t, 974, Part1(ax.MustReadFileLines("input")[0]))
}

func Part1(row string) int {
	hexBytes, err := hex.DecodeString(row)
	ax.Check(err, "failed to decode string")
	bs := hexAsBinary(hexBytes)
	p := parser{
		bs: bs,
	}
	p.parsePacket()
	return p.verSum
}
