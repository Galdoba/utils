package slicetricks

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	slice = Insert(slice, 2, []int{9, 8, 7}...)
	fmt.Println(slice)
}

func isEven(i int) bool {
	return i%2 == 0
}

func moreThan10(i int) bool {
	return i > 10
}

func TestFilterANY(t *testing.T) {
	slice := []int{1, 4, 7, 10, 13, 16, 2, -11, 0, -2}
	newSlice := FilterAll(slice, moreThan10, isEven)
	fmt.Println(slice)
	fmt.Println(newSlice)
}
