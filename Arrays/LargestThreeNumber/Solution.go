package LargestThreeNumber

import "fmt"

func main() {
	array := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	largest3Sum := getLargest3Number(array)
	fmt.Printf("Largest sum: %v", largest3Sum)
}

func getLargest3Number(array []int) (value []int) {
	firstNumber := array[0]
	secondNumber := array[1]
	thirdNumber := array[2]
	for counter := 3; counter < len(array); counter++ {
		currentValue := array[counter]
		if currentValue > firstNumber {
			if firstNumber > secondNumber {
				if secondNumber > thirdNumber {
					thirdNumber = secondNumber
				}
				secondNumber = firstNumber
			} else if firstNumber > thirdNumber {
				thirdNumber = firstNumber
			}
			firstNumber = currentValue
		} else if currentValue > secondNumber {
			if secondNumber > thirdNumber {
				thirdNumber = secondNumber
			}
			secondNumber = currentValue
		} else if currentValue > thirdNumber {
			thirdNumber = currentValue
		}
	}
	value = append(value, firstNumber, secondNumber, thirdNumber)
	return
}
