package problems

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	current := 0
	needToFill := 0
	cnt := 0

	for current < len(nums) {
		if nums[current] != val {
			nums[needToFill] = nums[current]
			needToFill++
		} else {
			cnt++
		}
		current++
	}
	return len(nums) - cnt
}

// @lc code=end
