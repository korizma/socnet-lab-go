package solutions

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo2"
	"gonum.org/v1/gonum/graph/simple"
)

func CheckGraphIfNodeRemoved(g *simple.UndirectedGraph, node_id int64) {

}

// this function returns the key with the max value in the map
func GetKeyWithMaxValue(m map[int64]float64) int64 {
	max_id := int64(-1)

	for id, val := range m {
		if max_id == -1 {
			max_id = id
			continue
		}

		if m[max_id] < val {
			max_id = id
		}
	}

	return max_id
}

// this function returns the top 1 node for BC, CC, EC
func GetTopCentralityNodes(g *simple.UndirectedGraph) (int64, int64, int64) {
	return 0, 0, 0
}

func Sol1() {
	// G, err := demo2.LoadZachary()
	// G, err := demo2.LoadFlorentine()
	// G, err := demo2.LoadWomen()
	G, err := demo2.LoadMiserables()
	if err != nil {
		fmt.Println("Error loading graph:", err)
		return
	}

	bc, cc, ec := GetTopCentralityNodes(G)

	fmt.Println("By removing the top 1 BC node:")
	CheckGraphIfNodeRemoved(G, bc)

	fmt.Println("By removing the top 1 CC node:")
	CheckGraphIfNodeRemoved(G, cc)

	fmt.Println("By removing the top 1 EC node:")
	CheckGraphIfNodeRemoved(G, ec)
}
