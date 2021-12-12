package ax

import (
	"constraints"
	"strconv"
)

// MustParseInt parses the input string as an int with base-10.
// If unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseInt[T constraints.Integer](s string) T {
	return MustParseIntBase[T](s, 10)
}

// MustParseIntBase[int] parses the input string as an int with the provided base.
// If unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseIntBase[T constraints.Integer](s string, base int) T {
	res, err := strconv.ParseInt(s, 10, 64)
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
