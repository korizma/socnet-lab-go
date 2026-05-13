package demo4

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/graph"
)

func GetClusterCoef(node graph.Node, graph *graph.Graph) float32 {
	ego_graph := graph.CreateEgoGraph(node)

	ego_graph.RemoveNode(node)

	edge_num := float32(len(ego_graph.GetEdges()))
	node_num := float32(len(ego_graph.GetNodes()))

	if node_num == 1 || node_num == 0 {
		return 0
	}

	return edge_num / (node_num * (node_num - 1) / 2)
}

func Demo4() {
	G := graph.LoadGraph("zachary.txt")

	nodes := G.GetNodes()
	for _, node := range nodes {
		cluster_coef := GetClusterCoef(node, G)

		fmt.Printf("Node: %d, ClusterCoef: %.4f\n", node.GetID(), cluster_coef)
	}
}
