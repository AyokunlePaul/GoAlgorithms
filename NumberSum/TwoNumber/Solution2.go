package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 5, -4, 8, 11, 1, -1, 6}
	expectedSum := 10
	sort.Ints(values)

	leftIndex := 0
	rightIndex := len(values) - 1
	for range values {
		valueSum := values[leftIndex] + values[rightIndex]
		if valueSum == expectedSum {
			fmt.Printf("%d, %d", values[leftIndex], values[rightIndex])
			return
		} else if valueSum > expectedSum {
			rightIndex = rightIndex - 1
		} else if valueSum < expectedSum {
			leftIndex = leftIndex + 1
		}
	}
}
