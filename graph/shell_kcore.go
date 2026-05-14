package graph

import "fmt"

func (graph *Graph) CalculateShellIndex(print_flag bool) map[int32]int {
	new_graph := graph.Copy()
	shell_index := make(map[int32]int)

	k := 0

	for {
		k += 1

		for {
			to_be_removed := make([]Node, 0)
			for key := range new_graph.adj {

				node := new_graph.idToNode[key]

				if new_graph.GetDegree(node) < k {
					shell_index[node.id] = k - 1
					to_be_removed = append(to_be_removed, node)
				}
			}

			if len(to_be_removed) == 0 {
				break
			}

			for _, node := range to_be_removed {
				new_graph.RemoveNode(node)
			}
		}

		if len(new_graph.adj) == 0 {
			break
		}

		if print_flag {
			fmt.Println("Core", k, "#Nodes", len(new_graph.adj), "#Links", len(new_graph.GetEdges()))
		}
	}

	return shell_index
}
