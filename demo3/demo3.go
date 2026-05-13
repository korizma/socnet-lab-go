package demo3

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/graph"
)

func FormDegreeDistribution(graph *graph.Graph) []float32 {
	degrees := []int{}
	max_degree := 0

	nodes := graph.GetNodes()

	for _, node := range nodes {

		node_degree := graph.GetDegree(node)
		degrees = append(degrees, node_degree)

		max_degree = max(max_degree, node_degree)
	}

	degree_dist := make([]float32, max_degree+1)

	for _, degree := range degrees {
		degree_dist[degree] += 1
	}

	for i, _ := range degree_dist {
		degree_dist[i] /= float32(len(nodes))
	}

	return degree_dist
}

func Demo3() {
	G := graph.LoadGraph("zachary.txt")

	degree_dist := FormDegreeDistribution(G)

	for degree, dist := range degree_dist {
		fmt.Printf("%d:\t%.2f\n", degree, dist*100)
	}
}
