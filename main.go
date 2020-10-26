package main

import (
	"Algorithms/Arrays"
	"fmt"
)

func main() {
	value := "AAIAAHIHHAIIHIBC"
	target := 2
	fmt.Println(Arrays.MaxLengthSubstring(value, target))
}
