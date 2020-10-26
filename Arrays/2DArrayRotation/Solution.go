package main

import "fmt"

func main() {
	array := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	fmt.Printf("Rotated array: %v", rotate(array))
}

func rotate(value [][]int) [][]int {
	for xIndex, _ := range value {
		for yIndex := xIndex; yIndex < len(value[xIndex]); yIndex++ {
			value[xIndex][yIndex], value[yIndex][xIndex] = value[yIndex][xIndex], value[xIndex][yIndex]
		}
	}

	for _, xValue := range value {
		reverse(xValue)
	}

	return value
}

func reverse(array []int) []int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
