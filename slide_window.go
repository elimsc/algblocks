func minSubArrayLen(target int, nums []int) int {
	var windowSum int
	var ans = len(nums) + 1
	var left, right int

	for right < len(nums) {
		for windowSum < target && right < len(nums) { // 扩张window
			windowSum += nums[right] // window为 [left, right]
			if windowSum >= target {
				windowLen := right - left + 1
				if ans > windowLen {
					ans = windowLen
				}
			}
			right++ // window为 [left, right)
		}
		for windowSum >= target && left < right { // 收缩窗口
			windowSum -= nums[left]
			left++
			if windowSum >= target {
				windowLen := right - left
				if ans > windowLen {
					ans = windowLen
				}
			}
		}
	}

	if ans == len(nums)+1 {
		return 0
	}
	return ans
}