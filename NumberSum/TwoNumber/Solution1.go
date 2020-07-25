package main

import "fmt"

func main() {
	values := []int{3, 5, -4, 8, 11, 1, -1, 6}
	expectedSum := 10
	hashTable := map[int]bool{}

	for _, value := range values {
		subtractedValue := expectedSum - value
		if hashTable[subtractedValue] {
			fmt.Printf("%d, %d\n", value, subtractedValue)
			return
		} else {
			hashTable[value] = true
		}
	}
}
