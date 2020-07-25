package main

import (
	"Algorithms/models"
	"fmt"
)

func main() {
	tree := models.CreateBinaryTree()
	fmt.Printf("%d", calculateNodeDepth(tree, 0))
}

func calculateNodeDepth(tree *models.BinaryTree, nodeLevel int) (currentSum int) {
	if tree == nil {
		currentSum = 0
		return
	}
	currentSum += nodeLevel + calculateNodeDepth(tree.Left, nodeLevel+1) + calculateNodeDepth(tree.Right, nodeLevel+1)

	return
}
