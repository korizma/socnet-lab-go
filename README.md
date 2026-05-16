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

This demo shows the shortest distance between every node in the graph. Also shows the average shortedst distance between every node, and the diameter of the graph.

It uses the function `path.DijkstraAllPaths` that returns the object `path.AllShortest`. You can call the method `Between` on that object so you get the information about the shortest path about that path.

## Demo 7

This demo prints out the top 10 nodes for Betweeness Centrality, Closeness Centrality and Eigenvector centrality.

Betweeness centrality is how many times the node appears on the shortest paths.
Closeness centrality represents how close the node is to the rest of the network.
Eigenvector centrality represents how important a node is in a network.

## Demo 8

This demo prints out the shell index for every node in the graph. It also prints out the sizes of the k-core subgraphs.

A k-core graph is a graph where all of the nodes have degree of at least k.
Shell index of a node is k, where k is the highest k-core subgraph the node is in.

## Demo 9

This demo shows algorithms for community detection. You can see Label Propagation, Greedy Modularity Optimization and the Louvain method.

Label propagation is an iterative method. In the start each node has its label, and after every iteration it gets the label that is most popular among his neigbours. This process is done until the labels stop changing.
Greedy Modularity Optimization is also an iterative method. Modularity is the score for how good a certain choice of communities is. Every node starts as their own community, and each iteration two communities are joined if their joining is the best for the modularity. This process is done untill no improvements can be done.
Louvain algorithm has 2 phases. The first phase is similar to GMO, but instead of merging communities it moves nodes between communities. The second phase makes a new graph where each community is a node and the edges between them represent the edges between communities. On that graph the first phase is then done.

## Demo 10

