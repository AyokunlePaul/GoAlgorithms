package models

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) IsLeafNode() bool {
	return tree.Left == nil && tree.Right == nil
}

func CreateBinaryTree() (tree *BinaryTree) {
	tree = &BinaryTree{
		Value: 1,
		Left: &BinaryTree{Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left:  &BinaryTree{Value: 8, Left: nil, Right: nil},
				Right: &BinaryTree{Value: 9, Left: nil, Right: nil},
			},
			Right: &BinaryTree{
				Value: 5, Left: nil, Right: nil,
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
