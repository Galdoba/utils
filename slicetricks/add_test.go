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
