package ax

import "strconv"

func MustParseInt(s string) int {
	return MustParseIntBase(s, 10)
}

func MustParseIntBase(s string, base int) int {
	res, err := strconv.ParseInt(s, 10, 64)
	Check(err, "parse int")
	return int(res)
}
