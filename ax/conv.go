package ax

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// Atoi aliases MustParseInt[int].
func Atoi(s string) int {
	return MustParseInt[int](s)
}

// Atou aliases MustParseInt[uint].
func Atou(s string) uint {
	return MustParseInt[uint](s)
}

// Atou16 aliases MustParseInt[uint].
func Atou16(s string) uint16 {
	return MustParseInt[uint16](s)
}

// Atof aliases MustParseFloat[float64].
func Atof(s string) float64 {
	return MustParseFloat[float64](s)
}

// MustParseInt parses the input string as an int with base-10.
// If unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseInt[T constraints.Integer](s string) T {
	return MustParseIntBase[T](s, 10)
}

// MustParseIntBase[int] parses the input string as an int with the provided base.
// If unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseIntBase[T constraints.Integer](s string, base int) T {
	res, err := strconv.ParseInt(s, base, 64)
	Check(err, "parse int")
	return T(res)
}

// MustParseFloat parses the input string as a float.
// If unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseFloat[T constraints.Float](s string) T {
	res, err := strconv.ParseFloat(s, 64)
	Check(err, "parse int")
	return T(res)
}
