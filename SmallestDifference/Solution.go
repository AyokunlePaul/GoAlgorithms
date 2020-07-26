package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	firstArray := []int{-1, 5, 10, 20, 28, 3, 25}
	secondArray := []int{26, 134, 135, 25, 17}
	fmt.Printf("Values with the smallest difference: %v", smallestArrayDifference(firstArray, secondArray))
}

func smallestArrayDifference(first, second []int) []int {
	value := make([]int, 2)
	sort.Ints(first)
	sort.Ints(second)

	currentDifference := math.MaxInt64
	smallestDifference := math.MaxInt64

	for firstIndex, secondIndex := 0, 0; firstIndex < len(first) && secondIndex < len(second); {
		firstValue := first[firstIndex]
		secondValue := second[secondIndex]

		if firstValue < secondValue {
			currentDifference = secondValue - firstValue
			firstIndex += 1
		} else if secondValue < firstValue {
			currentDifference = firstValue - secondValue
			secondIndex += 1
		} else {
			return []int{firstValue, secondValue}
		}

		if smallestDifference > currentDifference {
			smallestDifference = currentDifference
			value[0] = firstValue
			value[1] = secondValue
		}
	}
	return value
}

func absolute(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}
