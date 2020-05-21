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
