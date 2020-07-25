package main

import (
	"fmt"
)

var fibCache = map[int]int{
	0: 0,
	1: 1,
}

func main() {
	//fmt.Printf("The 6th fibonacci is: %d\n", getTheNthFibonacci(6))
	//fmt.Printf("The 9th fibonacci is: %d\n", getTheNthFibonacci(9))
	//fmt.Printf("The 10th fibonacci is: %d\n", getTheNthFibonacci(10))
	//fmt.Printf("The 20th fibonacci is: %d\n", getTheNthFibonacci(20))

	fmt.Printf("The 6th fibonacci is: %d\n", getTheNthFibonacci2(6))
	fmt.Printf("The 9th fibonacci is: %d\n", getTheNthFibonacci2(9))
	fmt.Printf("The 10th fibonacci is: %d\n", getTheNthFibonacci2(10))
	fmt.Printf("The 20th fibonacci is: %d\n", getTheNthFibonacci2(20))
}

func getTheNthFibonacci(n int) (value int) {
	if n < 2 {
		value = n
		return
	}
	value += getTheNthFibonacci(n-1) + getTheNthFibonacci(n-2)
	return
}

func getTheNthFibonacci2(n int) (value int) {
	if fibValue, ok := fibCache[n]; !ok {
		fibCache[n] = getTheNthFibonacci2(n-1) + getTheNthFibonacci2(n-2)
		value = fibCache[n]
	} else {
		value = fibValue
	}
	return
}
