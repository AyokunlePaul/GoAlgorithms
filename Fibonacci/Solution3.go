package main

import "fmt"

func main() {
	fmt.Printf("The 6th fibonacci is %d\n", calculateFibonacci(6))
	fmt.Printf("The 9th fibonacci is %d\n", calculateFibonacci(9))
	fmt.Printf("The 10th fibonacci is %d\n", calculateFibonacci(10))
	fmt.Printf("The 20th fibonacci is %d\n", calculateFibonacci(20))
}

func calculateFibonacci(n int) (value int) {
	first := 0
	second := 1
	for counter := 2; counter <= n; counter++ {
		value = first + second
		first, second = second, value
		//fmt.Printf("First: %d, Second: %d, Value: %d\n", first, second, value)
	}
	return
}
