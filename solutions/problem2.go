package solutions

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/demo5"
	"gonum.org/v1/gonum/graph/topo"
)

func GenerateGraphAndReturnNumConnectedComponents(n int64, p float64) (int, int) {
	g := demo5.GenerateErdosRenyiGraph(n, p)

	max_size := 0
	comps := topo.ConnectedComponents(g)

	for _, c := range comps {
		max_size = max(max_size, len(c))
	}

	return len(comps), max_size
}

func Sol2() {
	n := int64(100)
	p := 0.0
	step := 0.001

	for {
		numComps, maxCompSize := GenerateGraphAndReturnNumConnectedComponents(n, p)
		fmt.Printf("p: %.3f, numComps: %d, maxCompSize: %d\n", p, numComps, maxCompSize)
		p += step

		if numComps == 1 {
			break
		}
	}
}
