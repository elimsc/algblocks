package main

// 回文字符串，中心扩展法，返回该回文串的长度
func ExpandCenter(s string, left int, right int) int {
	var res int
	if left == right {
		res = 1
		right++
		left--
	}
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			res += 2
			left--
			right++
		} else {
			break
		}
	}
	return res
}
