package lab

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
)

func LoadGraph(filename string) (*simple.UndirectedGraph, error) {
	g := simple.NewUndirectedGraph()

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		fromID, err := strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			continue
		}
		toID, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			continue
		}
		if fromID == toID {
			continue
		}

		from, _ := g.NodeWithID(fromID)
		to, _ := g.NodeWithID(toID)
		g.SetEdge(g.NewEdge(from, to))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return g, nil
}
