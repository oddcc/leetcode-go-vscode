package problems

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	max := 0

	for i := 0; i < len(nums); i++ {
		if i > max {
			break
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}

	return max >= len(nums)-1
}

// @lc code=end
