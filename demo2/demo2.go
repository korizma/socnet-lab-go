package demo2

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/lab"
	"gonum.org/v1/gonum/graph/simple"
)

func LoadZachary() (*simple.UndirectedGraph, error) {
	return lab.LoadGraph("zachary.txt")
}

func LoadMiserables() (*simple.UndirectedGraph, error) {
	return lab.LoadGraph("les_miserables.txt")
}

func LoadWomen() (*simple.UndirectedGraph, error) {
	return lab.LoadGraph("southern_women.txt")
}

func LoadFlorentine() (*simple.UndirectedGraph, error) {
	return lab.LoadGraph("florentine.txt")
}

func Demo2() {
	g, err := LoadZachary()
	// g, err := LoadMiserables()
	// g, err := LoadWomen()
	// g, err := LoadFlorentine()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Nodes:", g.Nodes().Len())
	fmt.Println("Edges:", g.Edges().Len())
}
