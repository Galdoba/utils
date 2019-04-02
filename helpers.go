package utils

import (
	"fmt"
)

func CheckError(descr string, err error) string {
	message := ""
	if err != nil {
		message = descr + " failed: " + fmt.Sprintln(err)
		panic(message)
	}
	return message
}
