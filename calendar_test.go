package calendar

import (
	"testing"
	"time"
)

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

func TestDate(t *testing.T) {
	dd := Date("Y-m-d H:i:s", time.Now())
	t.Log(" Y-m-d H:i:s format:", dd)
}

func TestNowDate(t *testing.T) {
	dd := NowDate("Y-m-d H:M:S")
	t.Log("Y-m-d H:M:S format:", dd)
	dd = NowDate("F T")
	t.Log("F T format:", dd)
}
