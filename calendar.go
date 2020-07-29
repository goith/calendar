package calendar

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//2006-01-02 15:04:05
var formapMapper = map[string]string{
	"d": "02",
	"F": "2006-01-02",
	"H": "15",
	"i": "04",
	"m": "01",
	"M": "04",
	"R": "15:04",
	"s": "05",
	"S": "05",
	"T": "15:04:05",
	"x": "01/02/2006",
	"y": "06",
	"Y": "2006",
}

var replacer *strings.Replacer

func init() {
	var pairs []string
	for o, n := range formapMapper {
		pairs = append(pairs, o, n)
	}

	replacer = strings.NewReplacer(pairs...)

}

type Format int8

const (
	PHP Format = iota
	Linux
)

type Time struct {
	t time.Time
	f Format
}

func Now() Time {
	return Time{t: time.Now()}
}

func (t *Time) Date(format string) string {
	goDateFmt := replacer.Replace(format)
	return t.t.Format(goDateFmt)
}

func NowDate(format string) string {
	goDateFmt := replacer.Replace(format)
	return time.Now().Format(goDateFmt)
}

func Date(format string, t time.Time) string {
	goDateFmt := replacer.Replace(format)
	return time.Unix(t.Unix(), 0).Format(goDateFmt)
}

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
