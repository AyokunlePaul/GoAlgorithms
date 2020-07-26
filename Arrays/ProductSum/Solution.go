package main

import "fmt"

func main() {
	array := []interface{}{
		5, 2,
		[]interface{}{7, -1}, 3,
		[]interface{}{6, []interface{}{-13, 8}, 4}, //14 //32 //-15
	}
	fmt.Printf("Product sum: %d", getTheProductSum(array, 1))
}

func getTheProductSum(array []interface{}, level int) (value int) {
	for _, arrayValue := range array {
		switch arrayValue.(type) {
		case []interface{}:
			value += level * getTheProductSum(arrayValue.([]interface{}), level+1)
		default:
			value += level * arrayValue.(int)
		}
	}
	return
}
