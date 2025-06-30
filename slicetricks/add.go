package slicetricks

import "slices"

// Contains - return true if element is present in slice.
func Contains[T comparable](slice []T, element T) bool {
	return slices.Contains(slice, element)
}

func ContainsAllFunc[T any](slice []T, equalFunc func(e1, e2 T) bool, elements ...T) bool {
	for _, element := range elements {
		if !containsFunc(slice, equalFunc, element) {
			return false
		}
	}
	return true
}

func ContainsAnyFunc[T any](slice []T, equalFunc func(e1, e2 T) bool, elements ...T) bool {
	for _, element := range elements {
		if containsFunc(slice, equalFunc, element) {
			return true
		}
	}
	return false
}

func containsFunc[T any](slice []T, equalFunc func(e1, e2 T) bool, element T) bool {
	for _, inSlice := range slice {
		if equalFunc(inSlice, element) {
			return true
		}
	}
	return false
}

// ContainsAll - return true if ALL element is present in slice.
func ContainsAll[T comparable](slice []T, elements ...T) bool {
	containMap := make(map[T]bool)
	for i, e := range elements {
		if e == slice[i] {
			containMap[e] = true
		}
	}
	for _, val := range containMap {
		if !val {
			return false
		}
	}
	return true
}

// Append - its just append... nothing added.
func Append[T any](slice []T, elements ...T) []T {
	return append(slice, elements...)
}

// AppendUnique - append ONLY elements NOT contained in slice.
func AppendUnique[T comparable](slice []T, elements ...T) []T {
	for _, element := range elements {
		slice = appendUnique(slice, element)
	}
	return slice
}

func appendUnique[T comparable](slice []T, element T) []T {
	for _, value := range slice {
		if value == element {
			return slice
		}
	}
	return append(slice, element)
}

// Prepend - add elements
func Prepend[T any](slice []T, elements ...T) []T {
	for _, e := range elements {
		slice = append([]T{e}, slice...)
	}
	return slice
}

// Insert - insert elements starting from index.
func Insert[T any](slice []T, index int, elements ...T) []T {
	for i, element := range elements {
		slice = insert(slice, index+i, element)
	}
	return slice
}

func insert[T any](slice []T, index int, element T) []T {
	if index < 0 {
		index = 0
	}
	n := len(slice)
	if index > n {
		index = n
	}
	slice = append(slice, element)
	if index < n {
		copy(slice[index+1:], slice[index:])
		slice[index] = element
	}
	return slice
}

// FilterAny - filter slice to create new slice.
// new slice consists of elements if ANy condition is true.
func FilterAny[T any](slice []T, conditionConfirmFuncs ...func(T) bool) []T {
	newSlice := []T{}
elementLoop:
	for _, element := range slice {
		for _, condition := range conditionConfirmFuncs {
			if condition(element) {
				newSlice = append(newSlice, element)
				continue elementLoop
			}
		}
	}
	return newSlice
}

// FilterAny - filter slice to create new slice.
// new slice consists of elements if ALL condition is true.
func FilterAll[T any](slice []T, conditionConfirmFuncs ...func(T) bool) []T {
	newSlice := []T{}
elementLoop:
	for _, element := range slice {
		for _, condition := range conditionConfirmFuncs {
			if !condition(element) {
				continue elementLoop
			}
		}
		newSlice = append(newSlice, element)
	}
	return newSlice
}

// ExcludeDuplicates - return slice with all duplicated elebents excluded.
func ExcludeDuplicates[T comparable](slice []T) []T {
	elementsMet := make(map[T]int)
	for _, element := range slice {
		elementsMet[element]++
	}
	condencedSlice := []T{}
	for _, element := range slice {
		if elementsMet[element] != -1 {
			condencedSlice = append(condencedSlice, element)
			elementsMet[element] = -1
		}
	}
	return condencedSlice
}
