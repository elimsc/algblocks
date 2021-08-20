package main

// https://leetcode-cn.com/problems/basic-calculator/solution/shuang-zhan-jie-jue-tong-yong-biao-da-sh-olym/
func calculate(s string) int {
	var nums []int
	var ops []byte
	n := len(s)
	i := 0
	nums = append(nums, 0)
	needAppendZero := false // 用于将 (+1, (-1这些操作变为 (0+1, (0-1
	for i < n {
		switch s[i] {
		case ' ':
			i++
		case '(':
			ops = append(ops, '(')
			i++
			needAppendZero = true
		case '+':
			for len(ops) > 0 && ops[len(ops)-1] != '(' { // 执行计算
				nums, ops = compute(nums, ops)
			}
			if needAppendZero {
				nums = append(nums, 0)
				needAppendZero = false
			}
			ops = append(ops, '+')
			i++
		case '-':
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				nums, ops = compute(nums, ops)
			}
			if needAppendZero {
				nums = append(nums, 0)
				needAppendZero = false
			}
			ops = append(ops, '-')
			i++
		case ')':
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				nums, ops = compute(nums, ops)
			}
			ops = ops[:len(ops)-1]
			i++
			needAppendZero = false
		default: // 数字
			num := 0
			j := i
			for ; j < n && s[j] >= '0' && s[j] <= '9'; j++ {
				num = num*10 + int(s[j]-'0')
			}
			nums = append(nums, num)
			i = j
			needAppendZero = false
		}
	}
	for len(ops) > 0 {
		nums, ops = compute(nums, ops)
	}

	return nums[len(nums)-1]
}

func compute(nums []int, ops []byte) ([]int, []byte) {
	if len(ops) == 0 || len(nums) < 2 {
		return nums, ops
	}
	op := ops[len(ops)-1]
	ops = ops[:len(ops)-1]
	right := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	left := nums[len(nums)-1]
	nums = nums[:len(nums)-1]

	switch op {
	case '+':
		ans := left + right
		nums = append(nums, ans)
	case '-':
		ans := left - right
		nums = append(nums, ans)
	case '*':
		ans := left * right
		nums = append(nums, ans)
	case '/':
		ans := left / right
		nums = append(nums, ans)
	}

	return nums, ops
}
