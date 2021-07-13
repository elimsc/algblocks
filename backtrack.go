package main

// 回溯算法
// res: 结果, nums: 待选择路径, track:已选择路径
func Backtrack(res [][]int, nums []int, track []int) [][]int {
	// 1. 确定结束条件
	if len(nums) == 0 {
		res = append(res, append([]int{}, track...))
	}

	for i := range nums {
		track = append(track, nums[i])

		// 2. 确定下一层的待选择路径
		next := append([]int{}, nums[:i]...)
		next = append(next, nums[i+1:]...)
		res = Backtrack(res, next, track)

		track = track[:len(track)-1]
	}
	return res
}
