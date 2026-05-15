package demo7

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/lab"
	"gonum.org/v1/gonum/graph/network"
	"gonum.org/v1/gonum/graph/path"
)

func Demo7() {
	g, err := lab.LoadGraph("les_miserables.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// Betweenness centrality
	betweenness := network.Betweenness(g)

	fmt.Println("Betweeness centrality:")
	printTop10Centrality(betweenness, g)

	fmt.Println()

	// Closeness centrality
	closeness := network.Closeness(g, path.DijkstraAllPaths(g))

	fmt.Println("Closeness centrality:")
	printTop10Centrality(closeness, g)

	fmt.Println()

	// Eigenvector centrality (stub)
	eigenvector := EigenvectorCentrality(g)
	fmt.Println("Eigenvector centrality:")
	printTop10Centrality(eigenvector, g)
}
