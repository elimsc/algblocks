package main

import "math/rand"

func RandomPartition(a []int, l, r int) int {
	// rand.Seed(time.Now().UnixNano()) in main function
	i := rand.Int()%(r-l+1) + l // random select pivot
	a[i], a[r] = a[r], a[i]
	return Partition(a, l, r)
}

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
