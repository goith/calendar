package php

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	t.Log(Date("Y-m-d", time.Now()))
}
