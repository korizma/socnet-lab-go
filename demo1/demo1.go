package demo1

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/graph"
)

func create_simple_undirected_graph() *graph.Graph {
	G := graph.NewGraph()

	ana := graph.NewNode(0, map[string]any{"ime": "ana", "pol": "Z"})
	milovan := graph.NewNode(1, map[string]any{"ime": "milovan", "pol": "Z"})
	pera := graph.NewNode(2, map[string]any{"ime": "pera", "pol": "Z"})
	mika := graph.NewNode(3, map[string]any{"ime": "mika", "pol": "Z"})
	stojan := graph.NewNode(4, map[string]any{"ime": "stojan", "pol": "Z"})

	G.AddEdge(graph.NewEdge(ana, milovan, map[string]any{"kako": "voli"}))
	G.AddEdge(graph.NewEdge(ana, pera, map[string]any{"kako": "ne-voli"}))
	G.AddEdge(graph.NewEdge(ana, stojan, map[string]any{"kako": "voli"}))
	G.AddEdge(graph.NewEdge(pera, mika, map[string]any{"kako": "ne-voli"}))
	G.AddEdge(graph.NewEdge(milovan, stojan, map[string]any{"kako": "ne-voli"}))

	return G
}

func Demo1() {
	G := create_simple_undirected_graph()

	fmt.Println("Cvorovi:")
	fmt.Println()

	nodes := G.GetNodes()

	for _, node := range nodes {
		fmt.Println(node.GetAttr("ime"))
	}

	fmt.Println()
	fmt.Println("Veze:")
	fmt.Println()

	edges := G.GetEdges()

	for _, edge := range edges {
		fmt.Println(edge.Node1().GetAttr("ime"), edge.GetAttr("kako"), edge.Node2().GetAttr("ime"))
	}

	fmt.Println()
	fmt.Println("Susedstva:")
	fmt.Println()

	for _, node := range nodes {
		fmt.Println("Susedi", node.GetAttr("ime"))

		neighbours := G.GetNeighbours(node)

		for _, neighbour := range neighbours {
			fmt.Println(neighbour.GetAttr("ime"))
		}
		fmt.Println()
	}

}
