package problems

import "strings"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	var sb strings.Builder

	minL := 200
	for _, s := range strs {
		if len(s) < minL {
			minL = len(s)
		}
	}

	for c := 0; c < minL; c++ {
		pass := true
		tmp := strs[0][c]
		for _, s := range strs {
			if s[c] != tmp {
				pass = false
				break
			}
		}

		if !pass {
			break
		}
		sb.WriteByte(tmp)
	}

	return sb.String()
}

// @lc code=end
