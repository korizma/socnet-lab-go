package graph

import "fmt"

type Edge struct {
	node1 Node
	node2 Node
	attr  map[string]any
	id    string
}

func NewEdge(node1 Node, node2 Node, attr ...map[string]any) Edge {
	edge_id := fmt.Sprintf("%d_%d", min(node1.id, node2.id), max(node1.id, node2.id))
	edgeAttr := make(map[string]any)
	if len(attr) > 0 {
		edgeAttr = CopyMap(attr[0])
	}
	return Edge{node1: node1, node2: node2, attr: edgeAttr, id: edge_id}
}

func (edge Edge) SetAttr(name string, value any) {
	edge.attr[name] = value
}

func (edge Edge) GetAttr(name string) any {
	return edge.attr[name]
}

func (edge Edge) Copy() Edge {
	return Edge{node1: edge.node1, node2: edge.node2, attr: CopyMap(edge.attr)}
}

func (edge Edge) ToString() string {
	str_representation := "NodeID: " + string(edge.id)
	str_representation += "\nNode 1: " + edge.node1.ToString()
	str_representation += "\nNode 2: " + edge.node2.ToString()
	for key, val := range edge.attr {
		str_representation += "\n" + key + ": " + fmt.Sprint(val)
	}
	return str_representation
}

func (edge Edge) Node1() Node {
	return edge.node1
}

func (edge Edge) Node2() Node {
	return edge.node2
}
