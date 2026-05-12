package graph

type Node struct {
	id   int32
	attr map[string]any
}

type Edge struct {
	to   Node
	from Node
	attr map[string]any
}

type Graph struct {
	adjList  map[int32][]Node
	nodes    []Node
	directed bool
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

func NewEdge(nodeTo Node, nodeFrom Node) Edge {
	return Edge{to: nodeTo, from: nodeFrom, attr: make(map[string]any)}
}

func (edge Edge) AddAttr(name string, value any) {
	edge.attr[name] = value
}

func (edge Edge) GetAttr(name string) any {
	return edge.attr[name]
}

func NewGraph(directed bool) Graph {
	return Graph{adjList: make(map[int32][]Node), directed: directed}
}

func (graph Graph) AddNode(node Node) {
	// todo
}

func (graph Graph) AddEdge(edge Edge) {
	graph.adjList[edge.from.id] = append(graph.adjList[edge.from.id], edge.to)

	if !graph.directed {
		graph.adjList[edge.to.id] = append(graph.adjList[edge.to.id], edge.from)
	}
}

func (graph Graph) RemoveNode(node Node) {
	// TODO
}

func (graph Graph) RemoveEdge(edge Edge) {
	// TODO
}

func (graph Graph) GetNeighbours(node Node) []Node {
	return graph.adjList[node.id]
}

func (graph Graph) CreateEgoGraph(node Node) Graph {
	// todo
}

func (graph Graph) GetNodes() []Node {
	return graph.nodes
}
