package problems

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		t := nums[start]
		nums[start] = nums[end]
		nums[end] = t
		start++
		end--
	}
}

// @lc code=end
