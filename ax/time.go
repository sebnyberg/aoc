package ax

import (
	"testing"
	"time"
)

func PrintTime(t *testing.T, msg string, start time.Time) {
	t.Logf("%v\t%v\n", msg, time.Since(start))
}
