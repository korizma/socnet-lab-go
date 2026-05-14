package demo6

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/graph"
)

func dist(G *graph.Graph, node1 graph.Node, node2 graph.Node, print_path bool) int {
	path := G.GetShortestPath(node1, node2)
	if print_path {
		graph.PrintPath(path)
	}
	return len(path)
}

func avgMaxDist(G *graph.Graph) (float64, int) {
	nodes := G.GetNodes()

	connected_nodes := 0
	max_path_len := 0
	path_len_sum := 0

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			path_len := dist(G, nodes[i], nodes[j], false)
			if path_len > 0 {
				connected_nodes += 1
			}
			max_path_len = max(max_path_len, path_len)
			path_len_sum += path_len
		}
	}

	if connected_nodes == 0 {
		return 0, 0
	}

	return float64(path_len_sum) / float64(connected_nodes), max_path_len
}

func Demo6() {
	G := graph.LoadGraph("zachary.txt")

	nodes := G.GetNodes()
	fmt.Println("Zachary: ")
	fmt.Println("The path between the first node and the last node is:")
	dist(G, nodes[len(nodes)-1], nodes[0], true)
	fmt.Println()

	avgDist, maxDist := avgMaxDist(G)
	fmt.Println("Broj Cvorova:", len(nodes))
	fmt.Println("Prosecna udaljenost:", avgDist)
	fmt.Println("Dijametar:", maxDist)

	G = graph.GenerateErdosRenyiGraph(100, 0.2)

	nodes = G.GetNodes()
	fmt.Println()
	fmt.Println("Random graph: ")
	fmt.Println()

	avgDist, maxDist = avgMaxDist(G)
	fmt.Println("Broj Cvorova:", len(nodes))
	fmt.Println("Prosecna udaljenost:", avgDist)
	fmt.Println("Dijametar:", maxDist)
}
