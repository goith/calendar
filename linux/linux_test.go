package linux

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	t.Log(Date("x T", time.Now()))
}
