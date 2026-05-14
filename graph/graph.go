package graph

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adj      map[int32]map[int32]Edge
	idToNode map[int32]Node
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int32]map[int32]Edge), idToNode: make(map[int32]Node)}
}

// Adds node to the graph, if a node with the same ID exists it does nothing
func (graph *Graph) AddNode(node Node) {
	_, ok := graph.idToNode[node.id]
	if !ok {
		graph.idToNode[node.id] = node
		graph.adj[node.id] = make(map[int32]Edge)
	}
}

func (graph *Graph) AddEdge(edge Edge) {
	if edge.node1.id == edge.node2.id {
		return
	}
	_, ok := graph.adj[edge.node1.id][edge.node2.id]
	if ok {
		return
	}
	graph.AddNode(edge.node1)
	graph.AddNode(edge.node2)

	graph.adj[edge.node1.id][edge.node2.id] = edge
	graph.adj[edge.node2.id][edge.node1.id] = edge
}

func (graph *Graph) RemoveNode(node Node) {
	delete(graph.idToNode, node.id)
	delete(graph.adj, node.id)

	for _, adj_layer2 := range graph.adj {
		delete(adj_layer2, node.id)
	}
}

func (graph *Graph) RemoveEdge(edge Edge) {
	delete(graph.adj[edge.node1.id], edge.node2.id)
	delete(graph.adj[edge.node2.id], edge.node1.id)
}

func (graph *Graph) GetNeighbours(node Node) []Node {
	node_ids := MapKeysToSlice(graph.adj[node.id])
	neighbours := make([]Node, 0)

	for _, id := range node_ids {
		neighbours = append(neighbours, graph.idToNode[id])
	}
	return neighbours
}

func (graph *Graph) CreateEgoGraph(ego Node) *Graph {
	ego_graph := NewGraph()
	ego_graph.AddNode(ego)

	nodes := graph.GetNeighbours(ego)
	all_edges := graph.GetEdges()

	for _, edge := range all_edges {

		_, err1 := FindInSlice(nodes, edge.node1, CompareNodes)
		_, err2 := FindInSlice(nodes, edge.node2, CompareNodes)

		if !(CompareNodes(ego, edge.node1) || CompareNodes(ego, edge.node2)) {
			if err1 != nil || err2 != nil {
				continue
			}
		}

		ego_graph.AddEdge(edge.Copy())
	}

	return ego_graph
}

func (graph *Graph) GetNodes() []Node {
	node_ids := MapKeysToSlice(graph.idToNode)
	neighbours := make([]Node, 0)

	for _, id := range node_ids {
		neighbours = append(neighbours, graph.idToNode[id])
	}
	return neighbours
}

func (graph *Graph) GetEdges() []Edge {
	edges := make([]Edge, 0)

	for _, adj_layer2 := range graph.adj {
		for _, edge := range adj_layer2 {
			edges = append(edges, edge)
		}
	}

	return edges
}

// must be in root dir or full path
func LoadGraph(filename string) *Graph {
	G := NewGraph()

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
		return NewGraph()
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		pair := strings.Fields(line)

		if len(pair) < 2 {
			continue
		}

		node1_id, err := strconv.ParseInt(pair[0], 10, 32)
		if err != nil {
			continue
		}

		node2_id, err := strconv.ParseInt(pair[1], 10, 32)
		if err != nil {
			continue
		}

		G.AddEdge(NewEdge(NewNode(int32(node1_id)), NewNode(int32(node2_id))))
	}
	return G
}

func (graph Graph) GetDegree(node Node) int {
	return len(graph.adj[node.id])
}

// removes attributes from edges
func (graph Graph) GetConnectedComponents() []*Graph {
	components := make([]*Graph, 0)

	visited := make(map[int32]bool)

	nodes := graph.GetNodes()

	for _, node := range nodes {
		visited[node.id] = false
	}

	for _, node := range nodes {
		if visited[node.id] {
			continue
		}

		curr_graph := NewGraph()
		curr_graph.AddNode(node)
		stack := []Node{node}
		visited[node.id] = true

		for len(stack) != 0 {
			curr_node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			for neighbour := range graph.adj[curr_node.id] {
				if !visited[neighbour] {
					neighbour_node, _ := graph.GetNode(neighbour)
					curr_graph.AddEdge(NewEdge(curr_node, neighbour_node))
					stack = append(stack, neighbour_node)
					visited[neighbour] = true
				}
			}
		}

		components = append(components, curr_graph)
	}

	return components
}

// erdos_renyi_graph

func GenerateErdosRenyiGraph(nodes int, p float64) *Graph {
	graph := NewGraph()

	for i := 0; i < nodes; i++ {
		node_i := NewNode(int32(i))
		graph.AddNode(node_i)
		for j := i + 1; j < nodes; j++ {

			if rand.Float64() < p {
				node_j := NewNode(int32(j))
				graph.AddEdge(NewEdge(node_i, node_j))
			}

		}
	}

	return graph
}

func (graph *Graph) GetNode(id int32) (Node, bool) {
	node, ok := graph.idToNode[id]
	return node, ok
}

func (graph *Graph) GetEdge(node_id1 int32, node_id2 int32) (Edge, bool) {
	node, ok := graph.adj[node_id1][node_id2]
	return node, ok
}
