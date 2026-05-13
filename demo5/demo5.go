package demo5

import (
	"slices"

	"github.com/korizma/socnet-lab-go/graph"
)

func CompareGraphSize(g1 *graph.Graph, g2 *graph.Graph) int {
	return len(g1.GetNodes()) - len(g2.GetNodes())
}

func determine_connected_components(graph graph.Graph) []*graph.Graph {
	comps := graph.GetConnectedComponents()

	slices.SortFunc(comps, CompareGraphSize)

	return comps
}

func demo5() {

}
