package problems

/*
 * @lc app=leetcode.cn id=2938 lang=golang
 *
 * [2938] Separate Black and White Balls
 */

// @lc code=start
func minimumSteps(s string) int64 {
	var res int64
	res = 0
	nextWhiteI := 0
	current := 0

	for current < len(s) {
		if s[current] == '0' {
			if nextWhiteI != current {
				res += int64(current - nextWhiteI)
			}
			nextWhiteI++
		}
		current++
	}

	return res
}

// @lc code=end
