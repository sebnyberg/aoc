package ux

import "log"

// Check checks whether the provided error is non-nil. If it is non-nil, the
// message is printed followed by the error to log.Fatalf, which does an
// os.Exit(1) under the hood.
func Check(err error, msg string) {
	if err != nil {
		log.Fatalf("%v, %v\n", msg, err)
	}
}
