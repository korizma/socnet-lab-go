package solutions

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo9"
	"github.com/korizma/socnet-lab-go/lab"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/community"
)

func Louvain_6(g graph.Undirected) {
	reduced := community.Modularize(g, 1, nil)
	comms := reduced.Communities()
	examinePartition_6(g, comms)
}

func GMO_6(g graph.Undirected) {
	comms := demo9.GreedyModularityCommunities(g)
	examinePartition_6(g, comms)
}

func LP_6(g graph.Undirected) {
	comms := demo9.LabelPropagation(g)
	examinePartition_6(g, comms)
}

func examinePartition_6(g graph.Undirected, partition [][]graph.Node) {

	fmt.Println("Broj zajednica", len(partition))
	for _, communityNodes := range partition {
		fmt.Println(nodeIDs(communityNodes))
	}

	q := community.Q(g, partition, 1)
	fmt.Println("Modularnost", q)
}

func nodeIDs(nodes []graph.Node) []int64 {
	ids := make([]int64, 0, len(nodes))
	for _, node := range nodes {
		ids = append(ids, node.ID())
	}
	return ids
}

func Sol6() {
	G := lab.PlantedPartitionGraph(5, 10, 0.5, 0.05)

	LP_6(G)
	fmt.Println()

	Louvain_6(G)
	fmt.Println()

	GMO_6(G)
}
