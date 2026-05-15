package demo1

import (
	"fmt"

	"gonum.org/v1/gonum/graph/simple"
)

func create_simple_undirected_graph() *simple.UndirectedGraph {
	G := simple.NewUndirectedGraph()

	osobe := []Osoba{
		{id: 0, ime: "Ana", pol: "Z"},
		{id: 1, ime: "Milovan", pol: "M"},
		{id: 2, ime: "Pera", pol: "M"},
		{id: 3, ime: "Mika", pol: "M"},
		{id: 4, ime: "Stojan", pol: "M"},
	}

	veze := []Veza{
		{ko: osobe[0], koga: osobe[1], kako: "voli"},
		{ko: osobe[0], koga: osobe[2], kako: "ne-voli"},
		{ko: osobe[0], koga: osobe[4], kako: "voli"},
		{ko: osobe[2], koga: osobe[3], kako: "ne-voli"},
		{ko: osobe[1], koga: osobe[4], kako: "ne-voli"},
	}

	for _, osoba := range osobe {
		G.AddNode(osoba)
	}

	for _, veza := range veze {
		G.SetEdge(veza)
	}

	return G
}

func Demo1() {
	G := create_simple_undirected_graph()

	fmt.Println("Osobe:\n")

	osobe := G.Nodes()

	for osobe.Next() {
		osoba := osobe.Node().(Osoba)
		fmt.Println(osoba.ime, osoba.pol)
	}

	fmt.Println("\nVeze:\n")

	veze := G.Edges()

	for veze.Next() {
		veza := veze.Edge().(Veza)
		fmt.Println(veza.ko.ime, veza.kako, veza.koga.ime)
	}

	fmt.Println("\nSusedstva:")

	osobe = G.Nodes()

	for osobe.Next() {
		osoba := osobe.Node().(Osoba)

		print("\n")
		fmt.Println("Susedi osobe", osoba.ime)
		print("\n")
		veze = G.Edges()

		for veze.Next() {
			veza := veze.Edge().(Veza)

			if veza.ko == osoba {
				fmt.Println(veza.koga.ime, veza.kako)
			}
			if veza.koga == osoba {
				fmt.Println(veza.ko.ime, veza.kako)
			}
		}
	}
}
