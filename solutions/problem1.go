package solutions

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo2"
	"github.com/korizma/socnet-lab-go/demo7"
	"github.com/korizma/socnet-lab-go/demo8"
	"gonum.org/v1/gonum/graph/network"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func CheckGraphIfNodeRemoved(g *simple.UndirectedGraph, node_id int64) {
	conns := topo.ConnectedComponents(g)

	if len(conns) != 1 {
		fmt.Println("Graph is not fully connected!")
		return
	}

	newG := demo8.CopyGraph(*g)

	newG.RemoveNode(node_id)

	conns = topo.ConnectedComponents(newG)

	if len(conns) != 1 {
		fmt.Println("Graph is not fully connected after removing node:", node_id)
		fmt.Println("Number of components, after removing node:", len(conns))
		return
	}
	fmt.Println("Graph is fully connected after removing node:", node_id)
}

func GetKeyWithMaxValue(m map[int64]float64) int64 {
	max_id := int64(-1)

	for id, val := range m {
		if max_id == -1 {
			max_id = id
			continue
		}

		if m[max_id] < val {
			max_id = id
		}
	}

	return max_id
}

func GetTopCentralityNodes(g *simple.UndirectedGraph) (int64, int64, int64) {

	betweenness := network.Betweenness(g)

	closeness := network.Closeness(g, path.DijkstraAllPaths(g))

	eigenvector := demo7.EigenvectorCentrality(g)

	return GetKeyWithMaxValue(betweenness), GetKeyWithMaxValue(closeness), GetKeyWithMaxValue(eigenvector)
}

func Sol1() {
	// G, err := demo2.LoadZachary()
	// G, err := demo2.LoadFlorentine()
	// G, err := demo2.LoadWomen()
	G, err := demo2.LoadMiserables()
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	bc, cc, ec := GetTopCentralityNodes(G)

	fmt.Println("By removing the top 1 BC node:")
	CheckGraphIfNodeRemoved(G, bc)

	fmt.Println("By removing the top 1 CC node:")
	CheckGraphIfNodeRemoved(G, cc)

	fmt.Println("By removing the top 1 EC node:")
	CheckGraphIfNodeRemoved(G, ec)
}
