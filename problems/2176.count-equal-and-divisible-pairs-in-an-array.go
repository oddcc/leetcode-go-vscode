package problems

/*
 * @lc app=leetcode.cn id=2176 lang=golang
 *
 * [2176] Count Equal and Divisible Pairs in an Array
 */

// @lc code=start
func countPairs(nums []int, k int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && ((i*j)%k == 0) {
				res++
			}
		}
	}

	return res
}

// @lc code=end
