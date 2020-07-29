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

	//别的时区转成本地时区
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	//time.Local, time.UTC
	dd = Date("Y-m-d H:i:s", time.Date(2019, time.November, 10, 10, 0, 0, 0, time.UTC))
	t.Log("utc convert location, Y-m-d H:i:s format:", dd)

	tt1 := Time{t: time.Now()}
	dd = tt1.Date("Y-m-d H:i:s")
	t.Log(" Y-m-d H:i:s format:", dd)

	tt2 := Time{t: time.Now().AddDate(0, 0, -2)}
	dd = tt2.Date("Y-m-d H:i:s")
	t.Log(" Y-m-d H:i:s format:", dd)
}

func TestNowDate(t *testing.T) {
	dd := NowDate("Y-m-d H:M:S")
	t.Log("Y-m-d H:M:S format:", dd)
	dd = NowDate("F T")
	t.Log("F T format:", dd)
}
