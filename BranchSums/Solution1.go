package main

import "fmt"

func main() {
	tree := createBinaryTree()
	answer := new([]int)
	currentBranchSum := 0
	calculateBranchSum(tree, answer, currentBranchSum)
	fmt.Printf("%v", *answer)
}

func calculateBranchSum(tree *BinaryTree, answer *[]int, currentBranchSum int) {
	if tree == nil {
		return
	}
	newSum := currentBranchSum + tree.Value
	if tree.isLeafNode() {
		*answer = append(*answer, newSum)
		return
	} else {
		calculateBranchSum(tree.Left, answer, newSum)
		calculateBranchSum(tree.Right, answer, newSum)
	}
}

func createBinaryTree() (tree *BinaryTree) {
	tree = &BinaryTree{
		Value: 1,
		Left: &BinaryTree{Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left:  &BinaryTree{Value: 8, Left: nil, Right: nil},
				Right: &BinaryTree{Value: 9, Left: nil, Right: nil},
			},
			Right: &BinaryTree{
				Value: 5, Left: nil, Right: &BinaryTree{Value: 10, Left: nil, Right: nil},
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left:  &BinaryTree{Value: 6, Right: nil, Left: nil},
			Right: &BinaryTree{Value: 7, Right: nil, Left: nil},
		},
	}
	return
}
