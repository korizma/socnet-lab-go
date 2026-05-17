package solutions

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo10"
	"github.com/korizma/socnet-lab-go/demo2"
	"gonum.org/v1/gonum/graph"
)

func cmp(a, b demo10.LinkPredictionScore) int {
	return int(b.Score*1000) - int(a.Score*1000)
}

func getInterCommunityEdgePrediction(g graph.Graph) []demo10.LinkPredictionScore {

	// funkcija koja sortira slice LinkPredictionScore
	// slices.SortFunc(intercom_edge_predictions, cmp)

	return []demo10.LinkPredictionScore{}
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
