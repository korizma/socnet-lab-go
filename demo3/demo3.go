package demo3

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/lab"
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
	g, err := lab.LoadGraph("zachary.txt")
	if err != nil {
		fmt.Println("error:", err)
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
