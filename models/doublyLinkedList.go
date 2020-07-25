package models

type linkedListNode struct {
	right *linkedListNode
	left  *linkedListNode
	value int
}

type DoublyLinkedList struct {
	Head *linkedListNode
	Tail *linkedListNode
}

func (list *DoublyLinkedList) setHead(node *linkedListNode) {
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}
	list.insertBefore(list.Head, node)
}

func (list *DoublyLinkedList) setTail(node *linkedListNode) {
	if list.Tail == nil {
		list.setHead(node)
		return
	}
	list.insertAfter(list.Tail, node)
}

func (list *DoublyLinkedList) insertAt(position int, nodeToInsert *linkedListNode) {
	if position == 1 {
		list.setHead(nodeToInsert)
	}
	node := list.Head
	currentPosition := 1
	for node != nil && currentPosition != position {
		node = node.right
		currentPosition += 1
	}
	if node != nil {
		list.insertBefore(node, nodeToInsert)
	} else {
		list.setTail(nodeToInsert)
	}
}

func (list *DoublyLinkedList) insertBefore(node *linkedListNode, nodeToInsert *linkedListNode) {
	if nodeToInsert == list.Head && nodeToInsert == list.Tail {
		return
	}
	list.remove(nodeToInsert)
	previousNode := node.left
	nodeToInsert.left = previousNode
	nodeToInsert.right = node
	if node.left == nil {
		list.Head = nodeToInsert
	} else {
		previousNode.right = nodeToInsert
	}
	node.left = nodeToInsert
}

func (list *DoublyLinkedList) insertAfter(node *linkedListNode, nodeToInsert *linkedListNode) {
	if nodeToInsert == list.Head && nodeToInsert == list.Tail {
		return
	}
	list.remove(nodeToInsert)
	nextNode := node.right
	nodeToInsert.right = nextNode
	nodeToInsert.left = node
	if node.right == nil {
		list.Tail = nodeToInsert
	} else {
		nextNode.left = nodeToInsert
	}
	node.right = nodeToInsert
}

func (list *DoublyLinkedList) removeNodeWithValue(value int) {
	currentNode := list.Head
	for currentNode != nil {
		potentialNodeToRemove := currentNode
		currentNode = currentNode.right
		if potentialNodeToRemove.value == value {
			list.remove(potentialNodeToRemove)
		}
	}
}
func (list *DoublyLinkedList) remove(node *linkedListNode) {
	if node == list.Head {
		list.Head = list.Head.right
	} else if node == list.Tail {
		list.Tail = list.Tail.left
	}
	removeNodeBindings(node)
}

func (list *DoublyLinkedList) containsNodeWithValue(value int) bool {
	currentNode := list.Head
	for currentNode != nil && currentNode.value != value {
		currentNode = currentNode.right
	}
	return currentNode != nil
}

func removeNodeBindings(node *linkedListNode) {
	if node.left != nil {
		node.left.right = node.right
	}
	if node.right != nil {
		node.right.left = node.left
	}
	node.left = nil
	node.right = nil
}
