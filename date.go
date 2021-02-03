package utils

import (
	"strconv"
	"time"
)

//DateStamp - return current date as a string in format: [YYYY-MM-DD]
func DateStamp() string {
	y, m, d := time.Now().Date()
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
