package utils

import (
	"fmt"
	"math/rand"

	"github.com/Galdoba/convert"
)

//CheckError -
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

//BoundInt - если х не вписывается в min или max, приравнивает х к ближайшему из них.
func BoundInt(x, min, max int) int {
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

//BoundFloat64 - если х не вписывается в min или max, приравнивает х к ближайшему из них.
func BoundFloat64(x, min, max float64) float64 {
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

//SelectionOptionsMult - позволяет выбрать несколько опций возвращая
//перечни опций и решений от Пользователя
func SelectionOptionsMult(descr string, opt ...string) ([]string, []bool) {
	//сбор данных
	var optSlice []string
	var optStatuses []bool
	optSlice = append(optSlice, descr)
	optStatuses = append(optStatuses, false)
	for i := range opt {
		optSlice = append(optSlice, opt[i])
		optStatuses = append(optStatuses, false)
	}
	optSlice = append(optSlice, "[DONE]")
	optStatuses = append(optStatuses, false)
	printAllOptions(optSlice, optStatuses)
	done := false
	for !done {
		//pick := InputInt()
		var pick int
		fmt.Scan(&pick)
		if InRange(pick, 1, len(optSlice)-1) {
			fmt.Println("\033[0A"+convert.ItoS(pick), "toggled   ") //, optSlice[pick]) //убрать текст опции
			optStatuses[pick] = !optStatuses[pick]
		}
		fmt.Print("\033[" + convert.ItoS(len(optStatuses)+1) + "A")
		printAllOptions(optSlice, optStatuses)
		if pick == len(optSlice)-1 {
			done = true
		}
	}
	//анализ и возврат
	fmt.Println("\n")
	var returnSlc []string
	var resultSlc []bool
	for i := range optSlice {
		if i == 0 || i == len(optSlice)-1 {
			continue
		}
		returnSlc = append(returnSlc, optSlice[i])
		resultSlc = append(resultSlc, optStatuses[i])
	}
	return returnSlc, resultSlc
}

func printOption(optName string, optStatus bool, optNum int) {
	status := " "
	if optStatus == true {
		status = "X"
	}
	num := convert.ItoS(optNum)
	if InRange(optNum, 0, 9) {
		num = " " + num
	}
	fmt.Println(num + " [" + status + "] -- " + optName)
}

func printAllOptions(optSlice []string, optStatuses []bool) {
	for i := range optSlice {
		if i == 0 {
			fmt.Println(optSlice[0])
		} else {
			printOption(optSlice[i], optStatuses[i], i)
		}
	}
}
