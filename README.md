# Social Network Lab in Go

There are 10 demos in total and 6 problems you will need to solve.

# Content of the Demos

## Demo 1

This demo shows how you can create a Undirected Graph with the `gonum` library. In the `helper.go` file for this demo you can see how `Osoba` and `Veza` was declared. They satisfy the interfaces for the `graph.Node` and `graph.Edge`, and they can be used as nodes and edges in the `gonum` library.

In every other demo we use the implementation of nodes and edges in the `graph/simple` library. That library also contains the implementation of the undirected graph we will be using in all of our demos. `simple.UndirectedGraph` is an implementation of the interface `graph.Graph`.

## Demo 2

This demo shows how you can load some well known social networks. The graphs can be found in the projects root directory as `.txt` files. The format of these files is that in every line you can find one edge, given in the format `ID_Node1` `spacebar` `ID_Node2`. Using the implemented `LoadGraph` function found in `lab/graph.go` you can load any graph you want.

Every demo after this uses the functions declared for loading the graphs given in the root directory.

## Demo 3

This demo presents the Degree Distibution of every node in the graph.

Degree of a node represents how many edges that node has.

## Demo 4

This demo shows the Clustering Coefficient for each node in the graph. 

The Clustering Coefficient shows if the node is in a cluster. A value closer to 1 says that almost all of his neigbours have edges between them, but a value closer to 0 says that the node is not part of a cluster but more of a bridge between clusters.

The formula for the Clustering Coefficient for a node is the following:

```
CC = num_of_edges_between_his_neighbours / num_of_possible_edges_between_his_neighbours
```

## Demo 5

This demo shows how many Connected Components a graph has. First it showcases that on one of the graphs in the root directory, then on a generated Erdos-Renyi graph.

A Connected Component of graph is a subgraph so that you have a path between every pair of nodes.
Erdos-Renyi graph generation is a technique to generate a random graph. You are given 2 parameters, the number of nodes and the probability for every edge to exist. Then using these two parameters you can generate the graph. You can see the implementation in `demo5/helper.go`

## Demo 6

## Demo 7

## Demo 8

## Demo 9

## Demo 10