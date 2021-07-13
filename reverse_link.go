package main

// 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}

var next *ListNode

// 递归
func ReverseN1(head *ListNode, n int) *ListNode {
	if n == 1 {
		next = head.Next
		return head
	}
	last := ReverseN1(head.Next, n-1)
	head.Next.Next = head
	head.Next = next
	return last
}

// 迭代
func ReverseN2(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node1 := head
	node2 := head.Next
	for i := 0; i < n-1; i++ {
		next := node2.Next
		node2.Next = node1 // node2->node1
		node1 = node2
		node2 = next
	}
	head.Next = node2
	return node1
}
