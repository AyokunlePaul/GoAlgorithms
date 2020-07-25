package models

type graphNode struct {
	Name     string
	Children []*graphNode
}

func (node *graphNode) addChildren(nodes ...*graphNode) {
	for _, childNode := range nodes {
		node.Children = append(node.Children, childNode)
	}
}

func (node *graphNode) DepthFirstSearch(array *[]string) {
	*array = append(*array, node.Name)
	for _, children := range node.Children {
		children.DepthFirstSearch(array)
	}
}

func CreateGraph() (node *graphNode) {
	node = &graphNode{Name: "A"}
	nodeB := &graphNode{Name: "B"}
	nodeC := &graphNode{Name: "C"}
	nodeD := &graphNode{Name: "D"}
	nodeE := &graphNode{Name: "E"}
	nodeF := &graphNode{Name: "F"}
	nodeG := &graphNode{Name: "G"}
	nodeH := &graphNode{Name: "H"}
	nodeI := &graphNode{Name: "I"}
	nodeJ := &graphNode{Name: "J"}
	nodeK := &graphNode{Name: "K"}
	nodeF.addChildren(nodeI, nodeJ)
	nodeG.addChildren(nodeK)
	nodeB.addChildren(nodeE, nodeF)
	nodeD.addChildren(nodeG, nodeH)
	node.addChildren(nodeB, nodeC, nodeD)

	return
}
