package demo6

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo5"
	"github.com/korizma/socnet-lab-go/lab"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
)

func dist(shortest path.AllShortest, node1 graph.Node, node2 graph.Node) int {
	_, dist, _ := shortest.Between(node1.ID(), node2.ID())
	return int(dist)
}

func avgMaxDist(G graph.Graph) (float64, int64) {
	shortest := path.DijkstraAllPaths(G)

	sum := 0
	diameter := 0

	nodes := graph.NodesOf(G.Nodes())

	for i := 0; i < len(nodes); i++ {
		for j := i; j < len(nodes); j++ {
			d := dist(shortest, nodes[i], nodes[j])

			diameter = max(diameter, d)

			sum += d
		}
	}

	if len(nodes) < 3 {
		return 0, 0
	}

	return float64(sum) / float64(len(nodes)*(len(nodes)-1)/2), int64(diameter)
}

func Demo6() {
	g, err := lab.LoadGraph("zachary.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Zachary's karate club graph")
	avgDist, diameter := avgMaxDist(g)

	fmt.Printf("Number of nodes:%d\n", g.Nodes().Len())
	fmt.Printf("Average distance: %.2f\n", avgDist)
	fmt.Printf("Diameter: %d\n", diameter)

	fmt.Println()

	g = demo5.GenerateErdosRenyiGraph(100, 0.01)

	fmt.Println("Erdos-Renyi graph with 100 nodes and p=0.0")
	avgDist, diameter = avgMaxDist(g)

	fmt.Printf("Number of nodes:%d\n", g.Nodes().Len())
	fmt.Printf("Average distance: %.2f\n", avgDist)
	fmt.Printf("Diameter: %d\n", diameter)
}
