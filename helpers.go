package utils

import (
	"fmt"
	"go/types"
)

func CheckError(descr string, err error) string {
	message := ""
	if err != nil {
		message = descr + " failed: " + fmt.Sprintln(err)
		panic(message)
	}
	return message
}

//AppendUnique - проверяет есть ли новый элемент в слайсе.
//Если нет, то добавляет его в слайс. В противном случае возвращает слайс без изменений.
func AppendUnique(slice []types.Type, newElem types.Type) []types.Type {
	for i := range slice {
		if slice[i] == newElem {
			return slice
		}
	}
	slice = append(slice, newElem)
	return slice
}
