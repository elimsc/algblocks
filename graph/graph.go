package graph

import "fmt"

// ======= 图定义 ============

func NewGraph() *Graph {
	g := Graph{}
	g.Nodes = []*Node{}
	g.Edges = make(map[*Node][]*Node)
	return &g
}

type Node struct {
	Value int
}

type Graph struct {
	Nodes []*Node
	Edges map[*Node][]*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	// 无向图
	g.Edges[n1] = append(g.Edges[n1], n2) // 设定从节点n1到n2的边
	g.Edges[n2] = append(g.Edges[n2], n1) // 设定从节点n2到n1的边
}

// 节点的相邻节点
func (g *Graph) Adj(n *Node) []*Node {
	return g.Edges[n]
}

func (g *Graph) String() string {
	s := ""

	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + "->"
		near := g.Edges[g.Nodes[i]]

		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}

	return s
}

// ================ DFS ============

func (g *Graph) DFS(visit func(*Node)) {
	if len(g.Nodes) == 0 {
		return
	}
	visited := make(map[*Node]bool)
	start := g.Nodes[0]
	g.dfs(start, visit, visited)
}

func (g *Graph) dfs(start *Node, visit func(*Node), visited map[*Node]bool) {
	if visited[start] {
		return
	}
	visit(start)
	visited[start] = true
	adjs := g.Adj(start)
	for _, node := range adjs {
		g.dfs(node, visit, visited)
	}
}

// ================= BFS ==============
func (g *Graph) BFS(visit func(*Node)) {
	if len(g.Nodes) == 0 {
		return
	}
	visited := make(map[*Node]bool)
	queue := []*Node{}
	queue = append(queue, g.Nodes[0])
	for len(queue) != 0 {
		for _, node := range queue {
			if !visited[node] {
				visit(node)
				visited[node] = true
				queue = append(queue, g.Adj(node)...)
			}
			queue = queue[1:]
		}
	}
}
