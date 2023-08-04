package problems

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)

	for i := 0; i < l; i++ {
		if i == 0 {
			res[i] = 1
			continue
		}
		res[i] = res[i-1] * nums[i-1]
	}

	r := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			r = 1
		} else {
			r = r * nums[i+1]
		}
		res[i] = res[i] * r
	}

	return res
}

// @lc code=end
