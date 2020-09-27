package php

import (
	"strings"
	"time"
)

//2006-01-02 15:04:05
var mapper = map[string]string{
	"d": "02",
	"D": "Mon",
	"j": "2",
	"M": "Jan",
	"H": "15",
	"i": "04",
	"m": "01",
	"s": "05",
	"y": "06",
	"Y": "2006",
}

var replacer *strings.Replacer

func init() {
	var pairs []string
	for o, n := range mapper {
		pairs = append(pairs, o, n)
	}

	replacer = strings.NewReplacer(pairs...)

}

type Time struct {
	t time.Time
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
