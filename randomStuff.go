package utils

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func FloatToString(input_num float64, roundLimit int) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', roundLimit, 64)
}

//Дает случайное число float64
func RandFloat(min, max float64, precision int) float64 {
	res := min + rand.Float64()*(max-min)
	res = toFixed(res, precision)
	return res
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func romanNumberStr(i int) string {
	res := ""
	switch i {
	case 1:
		res = "I"
	case 2:
		res = "II"
	case 3:
		res = "III"
	case 4:
		res = "IV"
	case 5:
		res = "V"
	case 6:
		res = "VI"
	case 7:
		res = "VII"
	case 8:
		res = "VIII"
	case 9:
		res = "IX"
	case 10:
		res = "X"
	case 11:
		res = "XI"
	case 12:
		res = "XII"
	case 13:
		res = "XIII"
	case 14:
		res = "XIV"
	case 15:
		res = "XV"
	default:
	}
	return res
}

//RandomFromList - возвращает случайное значение из списка ([]string)
func RandomFromList(sl []string) string {
	l := len(sl)
	if l < 1 {
		return "Null"
	}
	return sl[randInt(0, l-1)]
}

func randomSeed() int64 {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return seed
}

func inRange(i, min, max int) bool {
	if i > min-1 && i < max+1 {
		return true
	}
	return false
}

func combineStrings(s, add string) string {
	return s + add
}

// function := map[string]func(int, int) int{
// 	"someFunction1": someFunction1,
// 	"someFunction2": someFunction2,
// }
// fmt.Println(someOtherFunction(3, 2, function["someFunction1"]))
// fmt.Println(someOtherFunction(3, 2, function["someFunction2"]))
// fmt.Println(someOtherFunction(3, 2, function["someFunction2"]))

// func someFunction1(a, b int) int {
// 	return a + b
// }

// func someFunction2(a, b int) int {
// 	return a - b
// }

// func someOtherFunction(a, b int, f func(int, int) int) int {
// 	return f(a, b)
// }

//TakeOptions - takes Q, slice of A and returns number of chosen A and string of that A
func TakeOptions(question string, options ...string) (int, string) {
	fmt.Println(question)
	for i := range options {
		prefix := "[" + strconv.Itoa(i+1) + "] - "
		fmt.Println(prefix + options[i])
	}
	answer := 0
	gotIt := false
	for !gotIt {
		answer = InputInt()
		if answer < len(options)+1 && answer > 0 {
			gotIt = true
		} else {
			fmt.Println("Answer is incorrect...")
			fmt.Println(question)
		}
	}
	return answer, options[answer-1]
}

func describe(descr []string) {
	if len(descr) > 0 {
		fmt.Print(descr[0])
	}
}

//InputInt - takes Integer from User
func InputInt(descr ...string) int {
	describe(descr)
	var dataVal int
	fmt.Scan(&dataVal)
	return dataVal
}

//InputFloat64 - takes Float64 from User
func InputFloat64(descr ...string) float64 {
	describe(descr)
	var dataVal float64
	fmt.Scan(&dataVal)
	return dataVal
}

//InputString - takes Float64 from User
func InputString(descr ...string) string {
	describe(descr)
	var dataVal string
	fmt.Scan(&dataVal)
	return dataVal
}

//Str2Float64 - convert String to Float64
func Str2Float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

//Str2Int -
func Str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

//Float64ToStr -
func Float64ToStr(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 0, 64)
}

//ClearScreen - clearing comand console (for Windows)
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func askYesNo(str string) bool {
	gotAnswer := false
	for !gotAnswer {
		fmt.Print(str + "(y/n) ")
		answer := InputString()
		switch answer {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Error: Answer is incorrect. (Type 'y' or 'n')")
		}
	}
	return false

}

func roll1dX(x int) int {
	if x < 1 {
		x = 1
	}
	return randInt(1, x)
}

func rollXdY(x int, y int) int {
	res := 0
	for i := 0; i < x; i++ {
		res = res + roll1dX(y)
	}
	return res
}

//RandomSeed - Дает случайный Seed
func RandomSeed() int64 {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return seed
}

func randInt(min int, max int) int {
	return min + rand.Intn(max)
}

//RollDice - возвращает результат броска нескольких дайсов по выражению '2d6' и добавляет N модификаторов к результату. Если X не указан, то равен 1 ('d6')
func RollDice(expression string, mods ...int) int {
	diceData := strings.Split(expression, "d")
	diceQty := 1
	diceType := 1
	switch len(diceData) {
	case 0:
		return -999
	case 1:
		diceType, _ = strconv.Atoi(diceData[0])
	default:
		diceQty, _ = strconv.Atoi(diceData[0])
		if diceData[0] == "" {
			diceQty = 1
		}
		diceType, _ = strconv.Atoi(diceData[1])
	}
	result := rollXdY(diceQty, diceType)
	for i := range mods {
		result = result + mods[i]
	}
	return result
}
