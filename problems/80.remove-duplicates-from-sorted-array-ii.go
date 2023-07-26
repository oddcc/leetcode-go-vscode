package problems

import "math"

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	cnt := 0
	i := 0
	f := 0
	rc := 0
	const init int = math.MinInt
	t := init

	for i < len(nums) {
		if t == init {
			t = nums[i]
			f++
			i++
			rc++
			continue
		}

		if t == nums[i] {
			if rc < 2 {
				nums[f] = nums[i]
				f++
			} else if rc >= 2 {
				cnt++
			}
			rc++
		} else {
			nums[f] = nums[i]
			t = nums[i]
			f++
			rc = 1
		}
		i++
	}

	return len(nums) - cnt
}

// @lc code=end
