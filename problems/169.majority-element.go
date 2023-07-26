package problems

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	candidate := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum == 0 {
			candidate = nums[i]
		}

		if nums[i] == candidate {
			sum++
		} else {
			sum--
		}
	}

	return candidate
}

// @lc code=end
