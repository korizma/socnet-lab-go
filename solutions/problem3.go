package solutions

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo2"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

func CalculateDistanceDistribution(g *simple.UndirectedGraph) map[int]int {
	paths := path.DijkstraAllPaths(g)

	distance_distribution := make(map[int]int)

	nodes := graph.NodesOf(g.Nodes())

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			_, dist, _ := paths.Between(nodes[i].ID(), nodes[j].ID())
			distance_distribution[int(dist)]++
		}
	}
	return distance_distribution
}

func normalizeDistanceDistribution(dist_dist map[int]int) (map[int]float64, int) {
	total := 0
	for _, count := range dist_dist {
		total += count
	}

	max_dist := 0
	normalized := make(map[int]float64)
	for dist, count := range dist_dist {
		normalized[dist] = float64(count) / float64(total)
		max_dist = max(max_dist, dist)

	}
	return normalized, max_dist
}

func Sol3() {
	// G, err := demo2.LoadZachary()
	// G, err := demo2.LoadFlorentine()
	// G, err := demo2.LoadWomen()
	G, err := demo2.LoadMiserables()
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	dist_dist := CalculateDistanceDistribution(G)
	normalized_dist_dist, max_dist := normalizeDistanceDistribution(dist_dist)

	fmt.Println("Distance Distribution:")
	for dist := 1; dist <= max_dist; dist++ {
		distribution, ok := normalized_dist_dist[dist]
		if ok {
			fmt.Printf("Distance %d: %.4f\n", dist, distribution)
		}
	}
}
