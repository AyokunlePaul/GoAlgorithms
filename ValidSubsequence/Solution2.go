package main

import "fmt"

func main() {
	mainArray := []int{-4, 8, 11, 9, 5, 10, 3, 12, 6, -1, 20}
	firstSubArray := []int{-4, 5, 10, 6, 20}
	secondSubArray := []int{-4, 5, 10, 6, 20, 8}

	fmt.Printf("First array is subsequence, %t\n", isValidSubsequence2(firstSubArray, mainArray))
	fmt.Printf("Second array is subsequence, %t\n", isValidSubsequence2(secondSubArray, mainArray))
}

func isValidSubsequence2(subsequence, mainArray []int) (returnedValue bool) {
	subsequenceIndex := 0
	for _, value := range mainArray {
		if subsequence[subsequenceIndex] == value {
			subsequenceIndex += 1
		}
	}
	returnedValue = subsequenceIndex == len(subsequence)
	return
}
