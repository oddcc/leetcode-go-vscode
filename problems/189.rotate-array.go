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
	cnt := 0
	for i := 0; i < len(nums); i++ {
		next := (i + k) % len(nums)
		replacedInit := false
		tmp := nums[i]
		for !(replacedInit && next == (i+k)%len(nums)) && cnt < len(nums) {
			if next == (i+k)%len(nums) {
				replacedInit = true
			}
			t := nums[next]
			nums[next] = tmp
			tmp = t
			next = (next + k) % len(nums)

			cnt++
		}

		if cnt == len(nums) {
			break
		}
	}
}

// @lc code=end
