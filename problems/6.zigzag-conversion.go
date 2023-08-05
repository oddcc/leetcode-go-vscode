package problems

import "strings"

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	var sb strings.Builder

	l := len(s)
	oneLoopLengh := 2*numRows - 2
	if l <= numRows {
		return s
	}
	if oneLoopLengh == 0 {
		return s
	}
	for r := 0; r < numRows; r++ {
		if r == 0 {
			for i := 0; i < l; i += oneLoopLengh {
				sb.WriteByte(s[i])
			}
		} else if r == numRows-1 {
			for i := numRows - 1; i < l; i += oneLoopLengh {
				sb.WriteByte(s[i])
			}
		} else {
			x := oneLoopLengh - 2*r
			for i := r; i < l; {
				sb.WriteByte(s[i])
				i = i + x
				x = oneLoopLengh - x
			}
		}
	}

	return sb.String()
}

// @lc code=end
