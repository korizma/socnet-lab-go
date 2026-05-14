package demo8

import (
	"fmt"
	"slices"

	"github.com/korizma/socnet-lab-go/graph"
)

func Demo8() {
	M := graph.LoadGraph("les_miserables.txt")

	shellIndex := M.CalculateShellIndex(true)
	nodeIDs := make([]int32, 0, len(shellIndex))
	for nodeID := range shellIndex {
		nodeIDs = append(nodeIDs, nodeID)
	}
	slices.Sort(nodeIDs)

	for _, nodeID := range nodeIDs {
		fmt.Println(nodeID, shellIndex[nodeID])
	}
}
