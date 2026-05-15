package demo3

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo2"
	"gonum.org/v1/gonum/graph"
)

func DegreeDistribution(g graph.Undirected) map[int]int {
	degree_dist := make(map[int]int)

	nodes := g.Nodes()
	for nodes.Next() {
		node := nodes.Node()
		degree := g.From(node.ID()).Len()
		degree_dist[degree]++
	}

	return degree_dist
}

func Demo3() {
	g, err := demo2.LoadZachary()
	// g, err := demo2.LoadFlorentine()
	// g, err := demo2.LoadMiserables()
	// g, err := demo2.LoadMiserables()
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	dist := DegreeDistribution(g)
	if len(dist) == 0 {
		fmt.Println("degree distribution not yet implemented")
		return
	}

	for degree, count := range dist {
		fmt.Printf("Degree %d: %d nodes\n", degree, count)
	}
}
