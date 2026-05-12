package demo3

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo2"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

func CalculateNodeDegree(node graph.Node, graph *simple.UndirectedGraph) int {
	edges := graph.Edges()
	degree := 0

	for edges.Next() {
		edge := edges.Edge()
		if node.ID() == edge.From().ID() || node.ID() == edge.To().ID() {
			degree += 1
		}
	}

	return degree
}

func FormDegreeDistribution(graph *simple.UndirectedGraph) []float32 {
	degrees := []int{}
	max_degree := 0

	nodes := graph.Nodes()

	for nodes.Next() {
		node := nodes.Node()

		node_degree := CalculateNodeDegree(node, graph)
		degrees = append(degrees, node_degree)

		max_degree = max(max_degree, node_degree)
	}

	degree_dist := make([]float32, max_degree+1)

	for _, degree := range degrees {
		degree_dist[degree] += 1
	}

	for i, _ := range degree_dist {
		degree_dist[i] /= float32(graph.Nodes().Len())
	}

	return degree_dist
}

func Demo3() {
	G := demo2.LoadZachary()

	degree_dist := FormDegreeDistribution(G)

	for degree, dist := range degree_dist {
		fmt.Printf("%d: %.2f\n", degree, dist)
	}
}
