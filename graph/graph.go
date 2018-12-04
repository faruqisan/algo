package graph

import (
	"fmt"
	"log"
)

type Node struct {
	Value string
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func NewGraph() *Graph {

	e := make(map[Node][]*Node)
	return &Graph{
		edges: e,
	}

}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(n1 Node, n2 *Node) {
	g.edges[n1] = append(g.edges[n1], n2)
}

func (g *Graph) PrintGraph() {
	for _, n := range g.nodes {

		output := fmt.Sprintf("%s -> ", n.Value)

		for _, ne := range g.edges[*n] {
			output += ne.Value + ", "
		}
		log.Println(output)
	}
}

func (g *Graph) GetNeighbour(source Node) string {
	output := fmt.Sprintf("%s -> ", source.Value)
	for _, ne := range g.edges[source] {
		output += ne.Value + ", "
	}
	return output
}
