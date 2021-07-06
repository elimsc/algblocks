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