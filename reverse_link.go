package main

// 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}

var afterN *ListNode

// 递归
func ReverseN1(head *ListNode, n int) *ListNode {
	if n == 1 {
		afterN = head.Next // 记录N个节点后的节点
		return head
	}
	last := ReverseN1(head.Next, n-1)
	head.Next.Next = head
	head.Next = afterN //让这N个节点的尾节点(head)指向后面的节点
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
		// 记录下一个节点
		next := node2.Next
		node2.Next = node1 // node2->node1

		// 指针后移
		node1 = node2
		node2 = next
	}
	head.Next = node2 // 让N个节点的尾节点(head)指向后面的节点
	return node1
}
