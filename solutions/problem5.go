package solutions

import (
	"fmt"
	"slices"

	"github.com/korizma/socnet-lab-go/demo10"
	"github.com/korizma/socnet-lab-go/demo2"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/community"
)

func cmp(a, b demo10.LinkPredictionScore) int {
	return int(b.Score*1000) - int(a.Score*1000)
}

func getInterCommunityEdgePrediction(g graph.Graph) []demo10.LinkPredictionScore {
	all_edge_predictions := demo10.AdamicAdarIndex(g)
	comms := community.Modularize(g, 1, nil)

	node_to_community := make(map[int64]int)
	for i, comm := range comms.Communities() {
		for _, node := range comm {
			node_to_community[node.ID()] = i
		}
	}

	intercom_edge_predictions := []demo10.LinkPredictionScore{}
	for _, pred := range all_edge_predictions {
		if node_to_community[pred.FromID] != node_to_community[pred.ToID] {
			intercom_edge_predictions = append(intercom_edge_predictions, pred)
		}
	}

	slices.SortFunc(intercom_edge_predictions, cmp)

	return intercom_edge_predictions
}

func Sol5() {
	// G, err := demo2.LoadZachary()
	// G, err := demo2.LoadFlorentine()
	// G, err := demo2.LoadWomen()
	G, err := demo2.LoadMiserables()
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	intercom_edges := getInterCommunityEdgePrediction(G)

	fmt.Println("Top 10 inter-community edge predictions:")
	for i, pred := range intercom_edges {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. From: %d, To: %d, Score: %.4f\n", i+1, pred.FromID, pred.ToID, pred.Score)
	}
}
