package model

import (
	"fmt"
	"strconv"
	"time"
)

func GetDay() string {
	//2006-01-02
	return time.Now().Format("20060102")
}
func GetWeek() string {
	year, week := time.Now().ISOWeek()
	return fmt.Sprintf("%d%d", year, week)
}

//根据周算所在月
//如果这周周三属于哪个月，这周就属于哪个月
func GetMonthOfWeek(in string) string {

	y, wTmp := in[:4], in[4:]

	w, _ := strconv.Atoi(wTmp)

	days := w * 7

	d, _ := time.Parse("20060102", y+"0101")
	ww := int(d.Weekday()) + 1                       //这一周从0开始
	s := d.AddDate(0, 0, days-ww-3).Format("200601") //周三

	return s
}

func GetMondayINWeek(in string) string {

	y, wTmp := in[:4], in[4:]

	w, _ := strconv.Atoi(wTmp)

	days := w * 7

	d, _ := time.Parse("20060102", y+"0101")
	ww := int(d.Weekday()) + 1 //这一周从0开始
	s := d.AddDate(0, 0, days-ww-5).Format("2006-01-02")

	return s
}

func GetWeekDay(wd time.Weekday) time.Time {
	now := time.Now()

	offset := int(wd - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	return weekStart
}

func GetMonth() string {
	//2006-01-02
	return time.Now().Format("200601")
}

func GetWeekNumBetweenMonth(a, b string) int {
	d1, _ := time.Parse("20060102", a+"01")
	tmp, _ := time.Parse("20060102", b+"01")
	d2 := GetLastDateOfMonth(tmp)

	days := int(((d2.Sub(d1).Hours())/24 + 1))

	weeks := 0
	if d1.Weekday() <= time.Wednesday {
		weeks += 1
		days -= (6 - int(d1.Weekday()) + 1)
	}

	if d2.Weekday() >= time.Wednesday {
		weeks += 1
		days -= (int(d1.Weekday()) + 1)
	}
	weeks += (days / 7)

	return weeks
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
