package demo10

import (
	"fmt"
	"slices"

	"gonum.org/v1/gonum/graph"
)

func calculateJaccardCoefficient(g graph.Graph, node1 graph.Node, node2 graph.Node) {

}

func JaccardCoefficient(g graph.Graph) []LinkPredictionScore {
	// TODO: Gonum v0.17.0 does not expose NetworkX-style Jaccard link prediction.
	return []LinkPredictionScore{}
}

func AdamicAdarIndex(g graph.Graph) []LinkPredictionScore {
	// TODO: Gonum v0.17.0 does not expose NetworkX-style Adamic-Adar link prediction.
	return []LinkPredictionScore{}
}

func printTop10(scores []LinkPredictionScore) {
	slices.SortFunc(scores, func(a LinkPredictionScore, b LinkPredictionScore) int {
		if a.Score > b.Score {
			return -1
		}
		if a.Score < b.Score {
			return 1
		}
		if a.FromID != b.FromID {
			return int(a.FromID - b.FromID)
		}
		return int(a.ToID - b.ToID)
	})

	limit := min(10, len(scores))
	for i := 0; i < limit; i++ {
		fmt.Printf("%d %d %.6f\n", scores[i].FromID, scores[i].ToID, scores[i].Score)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
