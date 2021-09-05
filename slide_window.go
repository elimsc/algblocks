package main

// https://programmercarl.com/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.html#_209-%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	minLen := len(nums) + 1

	l, r := 0, 0 // [l,r]
	for ; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target { // 收缩窗口
			curLen := r - l + 1
			if curLen < minLen {
				minLen = curLen
			}
			sum -= nums[l]
			l++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
