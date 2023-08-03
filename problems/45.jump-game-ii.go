package problems

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	minStep := 0

	for i := 0; i < len(nums); {
		max := 0
		maxIndex := 0
		minStep++
		if i+nums[i] >= len(nums)-1 {
			break
		}
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			if max < nums[j]+j {
				max = nums[j] + j
				maxIndex = j
			}
		}

		i = maxIndex
	}

	return minStep
}

// @lc code=end
