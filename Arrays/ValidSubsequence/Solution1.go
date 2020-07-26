package main

import "fmt"

func main() {
	mainArray := []int{-4, 8, 11, 9, 5, 10, 3, 12, 6, -1, 20}
	firstSubArray := []int{-4, 5, 10, 6, 20}
	secondSubArray := []int{-4, 5, 10, 6, 20, 8}

	fmt.Printf("First array is subsequence, %t\n", isValidSubsequence(firstSubArray, mainArray))
	fmt.Printf("Second array is subsequence, %t\n", isValidSubsequence(secondSubArray, mainArray))
}

func isValidSubsequence(subsequence, mainArray []int) (returnValue bool) {
	mainCounter := 0
	subsequenceCounter := 0
	for mainCounter < len(mainArray) && subsequenceCounter < len(subsequence) {
		if mainArray[mainCounter] == subsequence[subsequenceCounter] {
			subsequenceCounter += 1
		}
		mainCounter += 1
	}

	returnValue = subsequenceCounter == len(subsequence)
	return
}
