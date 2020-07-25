package main

import "fmt"

func main() {
	value := "a"
	fmt.Printf("Is Palindrome: %t", isPalindrome2(value))
}

func isPalindrome2(value string) bool {
	for counter := 0; counter < (len(value) / 2); counter++ {
		if value[counter] != value[(len(value)-(counter+1))] {
			return false
		}
	}
	return true
}
