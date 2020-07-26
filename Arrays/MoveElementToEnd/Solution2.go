package main

import "fmt"

func main() {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	secondArray := []int{4, 2, 4, 5, 6, 7, 2, 2, 4, 2, 7, 8, 2, 4, 5, 1, 4, 7, 8, 9, 2, 4}
	fmt.Printf("First array: %v\n", moveElementToEnd2(array, 2))
	fmt.Printf("Second array: %v\n", moveElementToEnd2(secondArray, 4))
}

func moveElementToEnd2(array []int, element int) []int {
	firstIndex := 0
	secondIndex := len(array) - 1
	for firstIndex < secondIndex {
		for firstIndex < secondIndex && array[secondIndex] == element {
			secondIndex -= 1
		}
		if array[firstIndex] == element {
			array[firstIndex], array[secondIndex] = array[secondIndex], array[firstIndex]
		}
		firstIndex += 1
	}
	return array
}
