package ux

import "strconv"

// MustParseFloat parses the provided int as a base-10, 64-bit integer. If parsing
// is unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseFloat(s string) int {
	return MustParseFloatBase(s, 10)
}

// MustParseFloat parses the provided int as a 10, 64-bit integer with the
// provided base. If parsing is unsuccessful, the error will be logged followed
// by os.Exit(1).
func MustParseFloatBase(s string, base int) int {
	res, err := strconv.ParseFloat(s, 10)
	Check(err, "parse int")
	return int(res)
}
