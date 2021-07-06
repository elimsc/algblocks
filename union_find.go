package main

type UnionFind struct {
	parent []int
	size   int
}

func NewUnionFind(size int) *UnionFind {
	uf := UnionFind{size: size}
	uf.parent = make([]int, size)
	for i := 0; i < size; i++ {
		uf.parent[i] = i
	}
	return &uf
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
	uf.size--
}

func (uf *UnionFind) Find(p int) int {
	if p >= len(uf.parent) {
		return -1
	}
	for uf.parent[p] != p {
		uf.parent[p] = uf.parent[uf.parent[p]] // 路径压缩
		p = uf.parent[p]
	}

	return p
}
