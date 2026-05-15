package solutions

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo2"
	"github.com/korizma/socnet-lab-go/demo9"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/community"
	"gonum.org/v1/gonum/graph/simple"
)

func IsCommunityStrong(G *simple.UndirectedGraph, community [][]graph.Node, indx int) bool {
	strong := true

	node_map := make(map[int64]bool)

	for _, node := range community[indx] {
		node_map[node.ID()] = true
	}

	for _, node := range community[indx] {
		neighbours := graph.NodesOf(G.From(node.ID()))

		good_neigbours := 0
		for _, neighbour := range neighbours {
			_, ok := node_map[neighbour.ID()]
			if ok {
				good_neigbours++
			}
		}
		if good_neigbours < len(neighbours)/2 {
			strong = false
			break
		}
	}
	return strong
}

func Louvain(g graph.Undirected) {
	reduced := community.Modularize(g, 1, nil)
	comms := reduced.Communities()
	examinePartition(g, comms)
}

func GMO(g graph.Undirected) {
	comms := demo9.GreedyModularityCommunities(g)
	examinePartition(g, comms)
}

func LP(g graph.Undirected) {
	comms := demo9.LabelPropagation(g)
	examinePartition(g, comms)
}

func examinePartition(g graph.Undirected, partition [][]graph.Node) {
	strong := 0

	for i := 0; i < len(partition); i++ {
		if IsCommunityStrong(g.(*simple.UndirectedGraph), partition, i) {
			strong++
			fmt.Printf("Zajednica %d je jaka\n", i)
		} else {
			fmt.Printf("Zajednica %d nije jaka\n", i)
		}
	}
	fmt.Printf("Broj jakih zajednica: %d\n", strong)

	fmt.Printf("Modularnost zajednice je: %.4f\n", community.Q(g, partition, 1))
}

func Sol4() {
	// G, err := demo2.LoadZachary()
	// G, err := demo2.LoadFlorentine()
	// G, err := demo2.LoadWomen()
	G, err := demo2.LoadMiserables()
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}
	LP(G)
	fmt.Println()

	Louvain(G)
	fmt.Println()

	GMO(G)
}
