package demo4

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo2"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

func CreateEgolessEgoGraph(node graph.Node, graph *simple.UndirectedGraph) *simple.UndirectedGraph {
	ego_graph := simple.NewUndirectedGraph()

	edges := graph.Edges()

	for edges.Next() {
		edge := edges.Edge()

		if edge.From().ID() == node.ID() {
			ego_graph.AddNode(edge.To())
		}
		if edge.To().ID() == node.ID() {
			ego_graph.AddNode(edge.From())
		}
	}

	edges.Reset()

	for edges.Next() {
		edge := edges.Edge()

		node1 := ego_graph.Node(edge.From().ID())
		if node1 == nil {
			continue
		}
		node2 := ego_graph.Node(edge.To().ID())
		if node2 == nil {
			continue
		}
		ego_graph.SetEdge(edge)
	}
	return ego_graph
}

func GetClusterCoef(node graph.Node, graph *simple.UndirectedGraph) float32 {
	ego_graph := CreateEgolessEgoGraph(node, graph)

	edge_num := float32(ego_graph.Edges().Len())
	node_num := float32(ego_graph.Nodes().Len())

	if node_num == 1 {
		return 0
	}

	return edge_num / (node_num * (node_num - 1) / 2)
}

func Demo4() {
	G := demo2.LoadZachary()

	nodes := G.Nodes()
	for nodes.Next() {
		node := nodes.Node()

		cluster_coef := GetClusterCoef(node, G)

		fmt.Printf("Node: %d, ClusterCoef: %.4f\n", node.ID(), cluster_coef)
	}
}
