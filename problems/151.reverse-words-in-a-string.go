package problems

import "strings"

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	var res strings.Builder

	isWord := false
	isFirstWord := true
	var tmp strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			isWord = false
			if tmp.Len() > 0 {
				if isFirstWord {
					isFirstWord = false
				} else {
					res.WriteByte(' ')
				}
				res.WriteString(reverseWord(tmp.String()))
				tmp.Reset()
			}
			continue
		}
		if s[i] != ' ' && !isWord {
			isWord = true
		}
		if i == 0 {
			tmp.WriteByte(s[i])
			if isFirstWord {
				isFirstWord = false
			} else {
				res.WriteByte(' ')
			}
			res.WriteString(reverseWord(tmp.String()))
			tmp.Reset()
		} else {
			tmp.WriteByte(s[i])
		}
	}

	return res.String()
}

func reverseWord(s string) string {
	bytes := []byte(s)

	for i, j := 0, len(bytes)-1; i < len(bytes) && j >= 0; i, j = i+1, j-1 {
		if i >= j {
			break
		}
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return string(bytes)
}

// @lc code=end
