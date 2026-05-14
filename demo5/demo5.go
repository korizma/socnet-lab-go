package demo5

import (
	"fmt"
	"slices"

	"github.com/korizma/socnet-lab-go/graph"
)

func CompareGraphSize(g1 *graph.Graph, g2 *graph.Graph) int {
	return len(g1.GetNodes()) - len(g2.GetNodes())
}

func DetermineConnectedComponents(graph *graph.Graph) []*graph.Graph {
	comps := graph.GetConnectedComponents()

	slices.SortFunc(comps, CompareGraphSize)

	return comps
}

func Demo5() {
	G := graph.LoadGraph("southern_women.txt")
	comps := DetermineConnectedComponents(G)

	fmt.Println("Broj komponenti:", len(comps))
	fmt.Println()
	for i, graph := range comps {
		fmt.Println("Graph", i, "ima", len(graph.GetNodes()), "cvorova")
	}

	nodes := 100
	p := 0.02
	randomGraph := graph.GenerateErdosRenyiGraph(nodes, p)

	comps = DetermineConnectedComponents(randomGraph)

	fmt.Println("Broj komponenti:", len(comps))
	fmt.Println()
	for i, graph := range comps {
		fmt.Println("Graph", i, "ima", len(graph.GetNodes()), "cvorova")
	}
}
