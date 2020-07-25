package models

type Node struct {
	Name     string
	Children []*Node
}

func (node *Node) addChildren(nodes ...*Node) {
	for _, childNode := range nodes {
		node.Children = append(node.Children, childNode)
	}
}

func (node *Node) DepthFirstSearch(array *[]string) {
	*array = append(*array, node.Name)
	for _, children := range node.Children {
		children.DepthFirstSearch(array)
	}
}

func CreateGraph() (node *Node) {
	node = &Node{Name: "A"}
	nodeB := &Node{Name: "B"}
	nodeC := &Node{Name: "C"}
	nodeD := &Node{Name: "D"}
	nodeE := &Node{Name: "E"}
	nodeF := &Node{Name: "F"}
	nodeG := &Node{Name: "G"}
	nodeH := &Node{Name: "H"}
	nodeI := &Node{Name: "I"}
	nodeJ := &Node{Name: "J"}
	nodeK := &Node{Name: "K"}
	nodeF.addChildren(nodeI, nodeJ)
	nodeG.addChildren(nodeK)
	nodeB.addChildren(nodeE, nodeF)
	nodeD.addChildren(nodeG, nodeH)
	node.addChildren(nodeB, nodeC, nodeD)

	return
}
