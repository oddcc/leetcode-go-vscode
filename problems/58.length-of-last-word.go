package problems

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	res := 0
	metNonSpace := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if !metNonSpace {
				continue
			}
			break
		}
		if s[i] != ' ' && !metNonSpace {
			metNonSpace = true
			res++
			continue
		}
		res++
	}

	return res
}

// @lc code=end
