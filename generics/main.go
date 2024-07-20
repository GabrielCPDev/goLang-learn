package main

import "fmt"

func main() {
	sliceInt := []int{1, 2, 3, 4, 5}
	sliceString := []string{"a", "b", "c", "d", "e", "f"}

	sliceReverseInt := reverse(sliceInt)
	sliceReverseString := reverse(sliceString)

	fmt.Println(sliceReverseInt)
	fmt.Println(sliceReverseString)

}

func reverse[T int | string](slice []T) []T {
	newInt := make([]T, len(slice))

	sliceSize := len(slice) - 1

	for i := 0; i < len(slice); i++ {
		newInt[sliceSize] = slice[i]
		sliceSize--
	}
	return newInt
}
