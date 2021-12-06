package ax

import (
	"constraints"
	"strconv"
)

func MustParseInt[T constraints.Integer](s string) T {
	return MustParseIntBase[T](s, 10)
}

// MustParseInt parses the provided int as a 10, 64-bit integer with the
// provided base. If parsing is unsuccessful, the error will be logged followed
// by os.Exit(1).
func MustParseIntBase[T constraints.Integer](s string, base int) T {
	res, err := strconv.ParseInt(s, 10, 64)
	Check(err, "parse int")
	return T(res)
}

// MustParseFloat parses the provided int as a base-10, 64-bit integer. If parsing
// is unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseFloat[T constraints.Float](s string) T {
	return MustParseFloatBase[T](s, 10)
}

// MustParseFloat parses the provided int as a 10, 64-bit integer with the
// provided base. If parsing is unsuccessful, the error will be logged followed
// by os.Exit(1).
func MustParseFloatBase[T constraints.Float](s string, base int) T {
	res, err := strconv.ParseFloat(s, 10)
	Check(err, "parse int")
	return T(res)
}
