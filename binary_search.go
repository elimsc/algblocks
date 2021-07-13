package main

// sort.Search
// 返回: 从左到右第一个让f为true的下标，第一个f不为false的下标
// f(mid)为true, 移动右指针, f(mid)为false, 移动左指针, 且+1
// 关键在于控制什么时候左指针右移, 即f为false
func Search(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

// TODO: SearchRight
