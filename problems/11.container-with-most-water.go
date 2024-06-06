package problems

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0

	for l < r {
		res = max(res, min(height[l], height[r]) * (r - l))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int{
	if a >= b {
		return a
	}
	return b
}

// @lc code=end
