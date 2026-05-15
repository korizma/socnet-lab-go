package demo1

import (
	"gonum.org/v1/gonum/graph"
)

type Osoba struct {
	id  int64
	ime string
	pol string
}

func (o Osoba) ID() int64 {
	return o.id
}

type Veza struct {
	ko   Osoba
	koga Osoba
	kako string
}

func (v Veza) From() graph.Node { return v.ko }

func (v Veza) To() graph.Node { return v.koga }

func (v Veza) ReversedEdge() graph.Edge { return Veza{ko: v.koga, koga: v.ko, kako: v.kako} }
