package graph

import "fmt"

type Node struct {
	id   int32
	attr map[string]any
}

type Edge struct {
	node1 Node
	node2 Node
	attr  map[string]any
	id    string
}

type Graph struct {
	adjList map[int32][]Node
	nodes   []Node
	edges   []Edge
}

func NewNode(id int32) Node {
	return Node{id: id, attr: make(map[string]any)}
}

func (node Node) AddAttr(name string, value any) {
	node.attr[name] = value
}

func (node Node) GetAttr(name string) any {
	return node.attr[name]
}

func NewEdge(node1 Node, node2 Node) Edge {
	edge_id := fmt.Sprintf("%d_%d", min(node1.id, node2.id), max(node1.id, node2.id))
	return Edge{node1: node1, node2: node2, attr: make(map[string]any), id: edge_id}
}

func (edge Edge) AddAttr(name string, value any) {
	edge.attr[name] = value
}

func (edge Edge) GetAttr(name string) any {
	return edge.attr[name]
}

func (edge Edge) Copy() Edge {
	return Edge{node1: edge.node1, node2: edge.node2, attr: CopyMap(edge.attr)}
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int32][]Node)}
}

func (graph *Graph) AddNode(node Node) {
	_, err := FindInSlice(graph.nodes, node, CompareNodes)
	if err != nil {
		graph.nodes = append(graph.nodes, node)
	}
}

func (graph *Graph) AddEdge(edge Edge) {
	_, err := FindInSlice(graph.edges, edge, CompareEdges)
	if err == nil {
		return
	}

	graph.AddNode(edge.node2)
	graph.AddNode(edge.node1)

	graph.edges = append(graph.edges, edge)

	graph.adjList[edge.node2.id] = append(graph.adjList[edge.node2.id], edge.node1)
	graph.adjList[edge.node1.id] = append(graph.adjList[edge.node1.id], edge.node2)
}

func (graph *Graph) RemoveNode(node Node) {
	graph.nodes = RemoveFromSlice(graph.nodes, node, CompareNodes)

	for _, neighbour := range graph.adjList[node.id] {

		graph.adjList[neighbour.id] = RemoveFromSlice(graph.adjList[neighbour.id], node, CompareNodes)
	}
	delete(graph.adjList, node.id)
}

func (graph *Graph) RemoveEdge(edge Edge) {
	graph.edges = RemoveFromSlice(graph.edges, edge, CompareEdges)

	graph.adjList[edge.node1.id] = RemoveFromSlice(graph.adjList[edge.node1.id], edge.node2, CompareNodes)

	graph.adjList[edge.node2.id] = RemoveFromSlice(graph.adjList[edge.node2.id], edge.node1, CompareNodes)

}

func (graph *Graph) GetNeighbours(node Node) []Node {
	return CopySlice(graph.adjList[node.id])
}

func (graph *Graph) CreateEgoGraph(ego Node) *Graph {
	ego_graph := NewGraph()
	ego_graph.AddNode(ego)

	nodes := graph.adjList[ego.id]

	for _, edge := range graph.edges {

		_, err1 := FindInSlice(nodes, edge.node1, CompareNodes)
		_, err2 := FindInSlice(nodes, edge.node2, CompareNodes)

		if err1 == nil && err2 == nil && !CompareNodes(ego, edge.node1) && !CompareNodes(ego, edge.node2) {
			continue
		}

		ego_graph.AddEdge(edge.Copy())
	}

	return ego_graph
}

func (graph *Graph) GetNodes() []Node {
	return CopySlice(graph.nodes)
}

func (graph *Graph) GetEdges() []Edge {
	return CopySlice(graph.edges)
}
