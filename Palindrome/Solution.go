package main

import "fmt"

func main() {
	value := "madam"
	fmt.Printf("Is Palindrome: %t", isPalindrome(value))
}

func isPalindrome(value string) bool {
	if len(value) <= 1 {
		return true
	}
	first := value[0]
	last := value[len(value)-1]
	return first == last && isPalindrome(value[1:len(value)-1])
}
