package graph

import "fmt"

type Node struct {
	id   int32
	attr map[string]any
}

func NewNode(id int32, attr ...map[string]any) Node {
	nodeAttr := make(map[string]any)
	if len(attr) > 0 {
		nodeAttr = CopyMap(attr[0])
	}
	return Node{id: id, attr: nodeAttr}
}

func (node Node) SetAttr(name string, value any) {
	node.attr[name] = value
}

func (node Node) GetAttr(name string) any {
	return node.attr[name]
}

func (node Node) ToString() string {
	str_representation := "NodeID: " + fmt.Sprint(node.id)
	for key, val := range node.attr {
		str_representation += "\n" + key + ": " + fmt.Sprint(val)
	}
	return str_representation
}

func (node Node) GetID() int32 {
	return node.id
}
