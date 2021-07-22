package main

import "math/rand"

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	for {
		a := rand7()
		b := rand7()
		n := a + (b-1)*7
		if n <= 40 {
			return n%10 + 1
		}
		a = n - 40 // rand9
		b = rand7()
		n = a + (b-1)*9 // 1-63, a是rand9，所以*9
		if n <= 60 {
			return n%10 + 1
		}
		a = n - 60 // rand3
		b = rand7()
		n = a + (b-1)*3 // 1-21
		if n <= 20 {
			return n%10 + 1
		}
	}
}
