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
