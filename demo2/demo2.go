package demo2

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
)

func LoadZachary() *simple.UndirectedGraph {
	G := simple.NewUndirectedGraph()

	file, err := os.ReadFile("zachary.txt")
	if err != nil {
		fmt.Println("error:", err)
		return simple.NewUndirectedGraph()
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		pair := strings.Split(line, " ")

		if len(pair) != 2 {
			return simple.NewUndirectedGraph()
		}

		node1_id, err := strconv.ParseInt(pair[0], 10, 32)
		if err != nil {
			return simple.NewUndirectedGraph()
		}

		node2_id, err := strconv.ParseInt(pair[1], 10, 32)
		if err != nil {
			return simple.NewUndirectedGraph()
		}

		node1, _ := G.NodeWithID(node1_id)
		node2, _ := G.NodeWithID(node2_id)

		G.SetEdge(G.NewEdge(node1, node2))
	}
	return G
}

func Demo2() {
	G := LoadZachary()
	fmt.Printf("%T\n", G)
	fmt.Println("Broj cvorova:", G.Nodes().Len())
	fmt.Println("Broj veza:", G.Edges().Len())
}
