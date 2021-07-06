package graph_test

import (
	"testing"

	"github.com/elimsc/algblocks/graph"
)

// go test -test.v
func TestDFS(t *testing.T) {
	g := graph.NewGraph()
	n1 := graph.Node{Value: 1}
	n2 := graph.Node{Value: 2}
	n3 := graph.Node{Value: 3}
	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddEdge(&n1, &n2)
	g.AddEdge(&n2, &n3)

	visitFn := func(n *graph.Node) { t.Log(n.Value) }

	t.Log(g)

	t.Log("DFS=========")
	g.DFS(visitFn)
	t.Log("BFS=========")
	g.BFS(visitFn)
}
