package main

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) isLeafNode() bool {
	return tree.Left == nil && tree.Right == nil
}
