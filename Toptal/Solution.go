package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(splitString2("babaa"))
	fmt.Println(splitString2("ababa"))
	fmt.Println(splitString2("aba"))
	fmt.Println(splitString2("bbbbb"))
}

func splitString(S string) (returnValue int) {
	returnValue = 0
	count := strings.Count(S, "a")
	if count % 3 != 0 {
		return
	}
	moduloCount, aCount := count / 3, 0
	countMap := map[int]int{}

	for _, value := range S {
		if string(value) == "a" { aCount += 1 }
		if aCount == 2 * moduloCount {
			if valueAccess, ok := countMap[moduloCount]; ok {
				returnValue += valueAccess
			}
		}
		if valueAccess, ok := countMap[aCount]; !ok {
			countMap[aCount] = 0 + 1
		} else {
			countMap[aCount] = valueAccess + 1
		}
	}

	return
}

func splitString2(S string) (returnValue int) {
	returnValue = 0
	stringLength := len(S)
	aCount := make([]int, stringLength)
	count := 0

	for index, runeValue := range S {
		if string(runeValue) == "a" {
			count += 1
		}
		aCount[index] = count
	}

	for index := 0; index <= stringLength - 3; index++ {
		for jIndex := index + 1; jIndex <= stringLength - 2; jIndex++ {
			firstCount := aCount[index]
			secondCount := aCount[jIndex] - aCount[index]
			thirdCount := aCount[stringLength - 1] - aCount[jIndex]
			if firstCount == secondCount && secondCount == thirdCount {
				returnValue += 1
			}
		}
	}

	return
}