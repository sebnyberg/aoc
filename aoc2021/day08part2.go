package aoc2021

import (
	"math/bits"
	"strconv"
	"strings"
)

func Day08Part2(rows []string) string {

	chars := func(field string) int {
		var res int
		for i := range field {
			res |= 1 << (field[i] - 'a')
		}
		return res
	}

	rowSum := func(row string) int {
		parts := strings.Split(row, "|")
		cipherFields := strings.Fields(parts[0])

		var matchChars [10]int
		for _, field := range cipherFields {
			switch len(field) {
			case 2:
				matchChars[1] = chars(field)
			case 3:
				matchChars[7] = chars(field)
			case 4:
				matchChars[4] = chars(field)
			case 7:
				matchChars[8] = chars(field)
			}
		}

		var sum int
		outputFields := strings.Fields(parts[1])
		for _, field := range outputFields {
			contents := chars(field)
			sum *= 10
			switch len(field) {
			case 2:
				sum += 1
			case 3:
				sum += 7
			case 4:
				sum += 4
			case 5:
				switch {
				case contents&matchChars[1] == matchChars[1]:
					sum += 3
				case bits.OnesCount(uint(contents&matchChars[4])) == 3:
					sum += 5
				default:
					sum += 2
				}
			case 6:
				switch {
				case contents&matchChars[1] != matchChars[1]:
					sum += 6
				case contents&matchChars[4] == matchChars[4]:
					sum += 9
				default:
				}
			case 7:
				sum += 8
			}
		}
		return sum
	}
	var sum int
	for _, row := range rows {
		sum += rowSum(row)
	}
	return strconv.Itoa(sum)
}
