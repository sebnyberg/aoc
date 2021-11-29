package ax

import "log"

func Check(err error, msg string) {
	if err != nil {
		log.Fatalf("%v, %v\n", msg, err)
	}
}
