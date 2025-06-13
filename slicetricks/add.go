package slicetricks

func Contains[T comparable](slice []T, element T) bool {
	for _, s := range slice {
		if element == s {
			return true
		}
	}
	return false
}

// Append - its just append... nothing added.
func Append[T any](slice []T, elements ...T) []T {
	return append(slice, elements...)
}

// AppendUnique - append ONLY elements NOT contained in slice.
func AppendUnique[T comparable](slice []T, elements ...T) []T {
	for _, e := range slice {
		if !Contains(slice, e) {
			slice = append(slice, e)
		}
	}
	return slice
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
	if index < 0 || index >= len(slice) {
		return slice
	}
	slice = append(slice[:index], append(elements, slice[index:]...)...)
	return slice
}

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
