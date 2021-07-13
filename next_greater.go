package main

func NextGreater(nums []int) []int {
	var res = make([]int, len(nums))
	var stack []int

	for i := len(nums) - 1; i >= 0; i-- {
		// 入单调栈
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, nums[i])

		// 取值 next greater
		if len(stack) == 1 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-2]
		}
	}
	return res
}
