package basic

// // 标准库的使用方法
// type MyHeap struct {
// 	Arr []int
// }

// func (h MyHeap) Len() int {
// 	return len(h.Arr)
// }

// // 最小堆
// func (h MyHeap) Less(i, j int) bool {
// 	return h.Arr[i] < h.Arr[j]
// }

// func (h MyHeap) Swap(i, j int) {
// 	h.Arr[i], h.Arr[j] = h.Arr[j], h.Arr[i]
// }

// func (h *MyHeap) Push(x interface{}) {
// 	h.Arr = append(h.Arr, x.(int))
// }

// func (h *MyHeap) Pop() interface{} {
// 	arr := h.Arr
// 	n := len(arr)
// 	x := arr[n-1]
// 	h.Arr = arr[:n-1]
// 	return x
// }

// func main() {
// 	var h = &MyHeap{}
// 	heap.Push(h, 2)
// 	heap.Push(h, 1)
// 	fmt.Println(heap.Pop(h)) // 1
// 	fmt.Println(heap.Pop(h)) // 2
// }

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	up(*h, h.Len()-1)
}

func (h *Heap) Pop() int {
	val := (*h)[0]
	n := h.Len()
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

// up和down中，i都是j的parent
func up(h Heap, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h Heap, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // 注意这里对j1范围的判断
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) { // 注意这里对j2范围的判断
			j = j2 // j取j1, j2中值小的一个
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
