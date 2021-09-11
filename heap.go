package main

import (
	"container/heap"
)

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// TODO: 记下源代码

func useHeap() {
	h := &NodeHeap{}
	heap.Init(h)
	// heap.Push(h, node)
	// node := heap.Pop(h).(*ListNode)
}

// Heap------------
type Heap []int

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	up(*h, len(*h)-1)
}

func (h *Heap) Pop() int {
	val := (*h)[0]
	n := len(*h)
	(*h)[0] = (*h)[n-1]
	*h = (*h)[:n-1]
	down(*h, 0, n-1)
	return val
}

func (h *Heap) Fix(i int) {
	if !down(*h, i, len(*h)) {
		up(*h, i)
	}
}

func up(h Heap, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(h[j] < h[i]) {
			break
		}
		h[i], h[j] = h[j], h[i]
		j = i
	}
}

func down(h Heap, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h[j2] < h[j1] {
			j = j2 // j取j1, j2中值小的一个
		}
		if !(h[j] < h[i]) {
			break
		}
		h[i], h[j] = h[j], h[i]
		i = j
	}
	return i > i0
}
