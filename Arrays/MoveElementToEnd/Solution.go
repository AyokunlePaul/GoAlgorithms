package main

import "fmt"

func main() {
	array := []int{4, 2, 4, 5, 6, 7, 2, 2, 4, 2, 7, 8, 2, 4, 5, 1, 4, 7, 8, 9, 2, 4}
	fmt.Printf("Final array: %v", moveElementToEnd(array, 2))
}

func moveElementToEnd(array []int, element int) []int {
	firstIndex := 0
	secondIndex := len(array) - 1
	for firstIndex < secondIndex {
		firstValue := array[firstIndex]
		secondValue := array[secondIndex]
		if firstValue == element {
			if secondValue != element {
				array[firstIndex], array[secondIndex] = secondValue, firstValue
				firstIndex += 1
			} else {
				secondIndex -= 1
			}
		} else if secondValue == element {
			secondIndex -= 1
		} else {
			firstIndex += 1
		}
	}
	return array
}
