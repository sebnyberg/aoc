package ux

import "strconv"

// MustParseInt parses the provided int as a base-10, 64-bit integer. If parsing
// is unsuccessful, the error will be logged followed by os.Exit(1).
func MustParseInt(s string) int {
	return MustParseIntBase(s, 10)
}

// MustParseInt parses the provided int as a 10, 64-bit integer with the
// provided base. If parsing is unsuccessful, the error will be logged followed
// by os.Exit(1).
func MustParseIntBase(s string, base int) int {
	res, err := strconv.ParseInt(s, 10, 64)
	Check(err, "parse int")
	return int(res)
}
