package graph

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adjList map[int32][]Node
	nodes   []Node
	edges   []Edge
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

	remove_edges := make([]Edge, 0)
	for _, edge := range graph.edges {
		if CompareNodes(node, edge.node1) || CompareNodes(node, edge.node2) {
			remove_edges = append(remove_edges, edge)
		}
	}

	for _, edge := range remove_edges {
		graph.edges = RemoveFromSlice(graph.edges, edge, CompareEdges)
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
	return CopySlice(graph.nodes)
}

func (graph *Graph) GetEdges() []Edge {
	return CopySlice(graph.edges)
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
	return len(graph.adjList[node.id])
}

// removes attributes from edges
func (graph Graph) GetConnectedComponents() []*Graph {
	components := make([]*Graph, 0)

	visited := make(map[int32]bool)

	for _, node := range graph.nodes {
		visited[node.id] = false
	}

	for _, node := range graph.nodes {
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

			for _, neighbour := range graph.adjList[curr_node.id] {
				if !visited[neighbour.id] {
					curr_graph.AddEdge(NewEdge(curr_node, neighbour))
					stack = append(stack, neighbour)
					visited[neighbour.id] = true
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
