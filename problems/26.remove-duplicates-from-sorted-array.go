package problems

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	cnt := 0
	const initT int = -999
	c := 0
	f := 0
	t := initT
	for c < len(nums) {
		if t == initT {
			t = nums[c]
			c++
			f++
			continue
		}
		if t == nums[c] {
			cnt++

		} else {
			nums[f] = nums[c]
			t = nums[c]
			f++
		}
		c++
	}

	return len(nums) - cnt
}

// @lc code=end
