package model

import "testing"

func TestGetMonthOfWeek(t *testing.T) {
	w := "202014"
	m := GetMonthOfWeek(w)
	t.Logf("%s week is in %s month", w, m)
}

func TestGetMondayINWeek(t *testing.T) {
	w := "202014"
	day := GetMondayINWeek(w)
	t.Logf("the date of %s week's Monday is %s", w, day)
}
