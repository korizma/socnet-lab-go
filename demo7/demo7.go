package demo7

import (
	"fmt"
	"slices"

	"github.com/korizma/socnet-lab-go/graph"
)

type centralityScore struct {
	nodeID int32
	score  float64
}

func printTop10(centrality map[int32]float64) {
	scores := make([]centralityScore, 0, len(centrality))

	for nodeID, score := range centrality {
		scores = append(scores, centralityScore{nodeID: nodeID, score: score})
	}

	slices.SortFunc(scores, func(a centralityScore, b centralityScore) int {
		if a.score > b.score {
			return -1
		}
		if a.score < b.score {
			return 1
		}
		return int(a.nodeID - b.nodeID)
	})

	limit := min(10, len(scores))
	for i := 0; i < limit; i++ {
		fmt.Printf("%d %.6f\n", scores[i].nodeID, scores[i].score)
	}
}

func Demo7() {
	G := graph.LoadGraph("les_miserables.txt")

	bc := G.CalculateBetweenessCentrality()
	cc := G.CalculateClosenessCentrality()
	ec := G.CalculateEigenvectorCentrality()

	fmt.Println("Betweenness centrality")
	printTop10(bc)

	fmt.Println()
	fmt.Println("Closeness centrality")
	printTop10(cc)

	fmt.Println()
	fmt.Println("Eigenvector centrality")
	printTop10(ec)
}
