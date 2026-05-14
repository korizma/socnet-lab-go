package graph

import "math"

func buildPaths(before map[int32][]int32, startID int32, endID int32) [][]int32 {
	if endID == startID {
		return [][]int32{{startID}}
	}

	paths := make([][]int32, 0)
	for _, parentID := range before[endID] {
		parentPaths := buildPaths(before, startID, parentID)
		for _, path := range parentPaths {
			path = append(path, endID)
			paths = append(paths, path)
		}
	}

	return paths
}

func (graph *Graph) getBetweenessCentralityUpdateMap(node1 Node, node2 Node) map[int32]float64 {
	before := make(map[int32][]int32)
	dist := make(map[int32]int32)

	dist[node1.id] = 0
	nodeQ := []Node{node1}
	levelQ := []int{0}

	path_exists := false
	path_level := -1

	for len(levelQ) > 0 {
		curr_node := nodeQ[0]
		curr_lvl := levelQ[0]
		nodeQ = nodeQ[1:]
		levelQ = levelQ[1:]

		if curr_node.id == node2.id {
			path_exists = true
			path_level = curr_lvl
			continue
		}

		if path_exists && curr_lvl >= path_level {
			continue
		}

		neighbours := graph.GetNeighbours(curr_node)
		for _, neighbour := range neighbours {
			neighbour_dist, ok := dist[neighbour.id]
			if !ok {
				before[neighbour.id] = []int32{curr_node.id}
				dist[neighbour.id] = int32(curr_lvl) + 1
				nodeQ = append(nodeQ, neighbour)
				levelQ = append(levelQ, curr_lvl+1)
			} else if neighbour_dist == int32(curr_lvl)+1 {
				before[neighbour.id] = append(before[neighbour.id], curr_node.id)
			}
		}
	}

	bc_update_map := make(map[int32]float64)

	if !path_exists {
		return bc_update_map
	}

	paths := buildPaths(before, node1.id, node2.id)
	path_num := float64(len(paths))

	for _, path := range paths {
		for _, id := range path[1 : len(path)-1] {
			bc_update_map[id] += 1 / path_num
		}
	}

	return bc_update_map
}

func (graph *Graph) CalculateBetweenessCentrality() map[int32]float64 {
	bc_map := make(map[int32]float64)

	nodes := graph.GetNodes()

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			update_map := graph.getBetweenessCentralityUpdateMap(nodes[i], nodes[j])

			for key, value := range update_map {
				bc_map[key] += value
			}
		}
	}

	if len(nodes) < 3 {
		return bc_map
	}

	normalization := float64((len(nodes) - 1) * (len(nodes) - 2) / 2)
	for id := range bc_map {
		bc_map[id] /= normalization
	}

	return bc_map
}

// needs to be fully connected
func (graph *Graph) CalculateClosenessCentrality() map[int32]float64 {
	cc_map := make(map[int32]float64)

	nodes := graph.GetNodes()

	for _, node := range nodes {
		cc_map[node.id] = 0
	}

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			dist := graph.GetShortestPath(nodes[i], nodes[j])
			cc_map[nodes[i].id] += float64(len(dist))
			cc_map[nodes[j].id] += float64(len(dist))
		}
	}

	for key := range cc_map {
		if cc_map[key] != 0 {
			cc_map[key] = float64(len(nodes)-1) / cc_map[key]
		}
	}

	return cc_map
}

func (graph *Graph) CalculateEigenvectorCentrality() map[int32]float64 {
	ec_map_prev := make(map[int32]float64)
	ec_map_curr := make(map[int32]float64)

	nodes := graph.GetNodes()

	for _, node := range nodes {
		ec_map_prev[node.id] = 1
	}

	iterations := 100

	for iterations > 0 {
		iterations -= 1

		for _, node := range nodes {
			ec_map_curr[node.id] = 0
			neighbours := graph.GetNeighbours(node)

			for _, neighbour := range neighbours {
				ec_map_curr[node.id] += ec_map_prev[neighbour.id]
			}
		}

		norm := 0.0
		for _, val := range ec_map_curr {
			norm += val * val
		}
		norm = math.Sqrt(norm)

		for id := range ec_map_curr {
			ec_map_curr[id] /= norm
		}

		temp := ec_map_curr
		ec_map_curr = ec_map_prev
		ec_map_prev = temp
	}
	return ec_map_prev
}
