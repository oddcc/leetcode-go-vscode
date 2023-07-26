package problems

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	queue := make([]int, 0)
	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			queue = append(queue, i)
			cnt++
		} else {
			if len(queue) > 0 {
				fill := queue[0]
				queue = queue[1:]
				queue = append(queue, i)
				nums[fill] = nums[i]
			}
		}
	}

	return len(nums) - cnt
}

// @lc code=end
