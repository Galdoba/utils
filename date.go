package utils

import (
	"fmt"
	"strconv"
	"time"
)

//DateStamp - return current date as a string in format: [YYYY-MM-DD]
//can take time.Duration arguments to move from Current Date
func DateStamp(durList ...time.Duration) string {
	currentTime := time.Now()
	for _, dur := range durList {
		currentTime = currentTime.Add(dur)
	}
	y, m, d := currentTime.Date()
	yy := strconv.Itoa(y)
	mm := strconv.Itoa(int(m))
	dd := strconv.Itoa(d)
	if int(m) < 10 {
		mm = "0" + mm
	}
	if d < 10 {
		dd = "0" + dd
	}
	return yy + "-" + mm + "-" + dd
}

//DurationStamp - return duration (float64 - seconds) as a string in format: [HH:MM:SS.ms]
func DurationStamp(dur float64) string {
	if dur < 0 {
		return "--NEGATIVE--"
	}
	stamp := ""
	hh := int(dur) / int(3600)
	hLeft := int(dur) % int(3600)
	mm := hLeft / 60
	ss := hLeft % 60
	sLeft := dur - (float64(hh*3600) + float64(mm*60) + float64(ss))
	ms := int(sLeft * 1000)
	////////
	hhs := strconv.Itoa(int(hh))
	for len(hhs) < 2 {
		hhs = "0" + hhs
	}
	mms := strconv.Itoa(int(mm))
	for len(mms) < 2 {
		mms = "0" + mms
	}
	sss := strconv.Itoa(int(ss))
	for len(sss) < 2 {
		sss = "0" + sss
	}
	mss := strconv.Itoa(int(ms))
	for len(mss) < 3 {
		mss = "0" + mss
	}
	stamp = hhs + ":" + mms + ":" + sss + "." + mss
	return stamp

}

//DateStampFrom - same as DateStamp() but requires input Time
func DateStampFrom(inputTime time.Time) string {
	m := ""
	switch inputTime.Month() {
	case time.January:
		m = "01"
	case time.February:
		m = "02"
	case time.March:
		m = "03"
	case time.April:
		m = "04"
	case time.May:
		m = "05"
	case time.June:
		m = "06"
	case time.July:
		m = "07"
	case time.August:
		m = "08"
	case time.September:
		m = "09"
	case time.October:
		m = "10"
	case time.November:
		m = "11"
	case time.December:
		m = "12"
	default:
		m = "Error"
	}
	d := strconv.Itoa(inputTime.Day())
	if len(d) < 2 {
		d = "0" + d
	}
	return fmt.Sprintf("%v-%v-%v", inputTime.Year(), m, d)
}
