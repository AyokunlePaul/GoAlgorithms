package main

import (
	"Algorithms/models"
	"fmt"
)

func main() {
	graph := models.CreateGraph()
	answer := new([]string)

	graph.DepthFirstSearch(answer)
	fmt.Printf("%v", *answer)
}
