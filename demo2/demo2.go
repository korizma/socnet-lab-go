package demo2

import (
	"fmt"

	"github.com/korizma/socnet-lab-go/graph"
)

func Demo2() {
	Z := graph.LoadGraph("zachary.txt")
	fmt.Printf("%T\n", Z)
	fmt.Println("Broj cvorova:", len(Z.GetNodes()))
	fmt.Println("Broj veza:", len(Z.GetEdges()))
}
