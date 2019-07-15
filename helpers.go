package utils

import (
	"fmt"
	"math/rand"
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

//PickFewUniqueFromList - выбирает из списка несколько случайных позиций
//если длинна списка меньше n возвращает весь список в случайном порядке
func PickFewUniqueFromList(list []string, n int) []string {
	var result []string
	if n < 1 {
		return result
	}
	if len(list) < n {
		for i := range list {
			j := rand.Intn(i + 1)
			list[i], list[j] = list[j], list[i]
		}
		return list
	}
	for len(result) < n {
		result = AppendUniqueStr(result, RandomFromList(list))
	}
	return result
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
