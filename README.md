# algorithm building blocks

BFS
```go
for len(q) != 0 {
    n := len(q)
    for _, node := range q {
        res = append(res, node.Val) // 替换为实际代码

        if node.Left != nil { // 注意判断非空
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
    }
    q = q[n:]
}
```

partition
```go
func Partition(a []int, l, r int) int {
	pivot := a[r]
	k := l
	for i := l; i < r; i++ {
		if a[i] < pivot {
			a[i], a[k] = a[k], a[i]
			k++
		}
	}
	a[k], a[r] = a[r], a[k]
	return k
}
```

add linked list
```go
node1, node2 := targetLink, another // node1链表更长
var prevNode1 *ListNode // 记录node1指针的前一个位置
for node1 != nil || node2 != nil {
    var sum int
    if node2 == nil {
        sum = node1.Val + carry
    } else {
        sum = node1.Val + node2.Val + carry
    }
    carry = 0
    if sum >= 10 {
        sum -= 10
        carry = 1
    }
    prevNode1 = node1
    node1.Val = sum
    node1 = node1.Next
    if node2 != nil {
        node2 = node2.Next
    }
}
if carry == 1 { // 新增尾节点
    newNode := &ListNode{Val: 1}
    prevNode1.Next = newNode
}
```