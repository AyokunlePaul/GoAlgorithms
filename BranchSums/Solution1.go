package main

import (
	"Algorithms/models"
	"fmt"
)

func main() {
	tree := models.CreateBinaryTree()
	answer := new([]int)
	currentBranchSum := 0
	calculateBranchSum(tree, answer, currentBranchSum)
	fmt.Printf("%v", *answer)
}

func calculateBranchSum(tree *models.BinaryTree, answer *[]int, currentBranchSum int) {
	if tree == nil {
		return
	}
	newSum := currentBranchSum + tree.Value
	if tree.IsLeafNode() {
		*answer = append(*answer, newSum)
		return
	} else {
		calculateBranchSum(tree.Left, answer, newSum)
		calculateBranchSum(tree.Right, answer, newSum)
	}
}
