package main

import (
	"fmt"
	"sort"
)

func main() {
	targetSum := 0
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	var finalArray []interface{}

	sort.Ints(array)

	for index := 0; index < len(array)-2; index++ {
		leftIndex := index + 1
		rightIndex := len(array) - 1
		for leftIndex < rightIndex {
			currentSum := array[index] + array[leftIndex] + array[rightIndex]
			if currentSum == targetSum {
				finalArray = append(finalArray, []int{array[index], array[leftIndex], array[rightIndex]})
				leftIndex += 1
				rightIndex -= 1
			} else if currentSum < targetSum {
				leftIndex += 1
			} else if currentSum > targetSum {
				rightIndex -= 1
			}
		}
	}

	fmt.Printf("Final value: %v", finalArray)
}
