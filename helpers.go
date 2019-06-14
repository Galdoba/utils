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

//AppendUniqueStr - проверяет есть ли новый элемент в слайсе.
//Если нет, то добавляет его в слайс. В противном случае возвращает слайс без изменений.
func AppendUniqueStr(slice []string, newElem string) []string {
	for i := range slice {
		if slice[i] == newElem {
			return slice
		}
	}
	slice = append(slice, newElem)
	return slice
}

//AppendUniqueInt - проверяет есть ли новый элемент в слайсе.
//Если нет, то добавляет его в слайс. В противном случае возвращает слайс без изменений.
func AppendUniqueInt(slice []int, newElem int) []int {
	for i := range slice {
		if slice[i] == newElem {
			return slice
		}
	}
	slice = append(slice, newElem)
	return slice
}
