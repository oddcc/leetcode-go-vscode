package problems

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	si := 0
	ti := 0

	for si < len(s) && ti < len(t) {
		if string(s[si]) == string(t[ti]) {
			si++
			ti++
		} else {
			ti++
		}
		if si == len(s) {
			return true
		}
	}

	return false
}

// @lc code=end
