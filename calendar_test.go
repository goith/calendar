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

func TestGetWeekNumBetweenMonth(t *testing.T) {
	a := "202004"
	b := "202004"
	w := GetWeekNumBetweenMonth(a, b)
	t.Logf("month %s to %s is %d weeks", a, b, w)

	a = "202004"
	b = "202005"
	w = GetWeekNumBetweenMonth(a, b)
	t.Logf("month %s to %s is %d weeks", a, b, w)
}
