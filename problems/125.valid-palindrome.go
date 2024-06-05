package problems

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		lc := string(s[l])
		rc := string(s[r])
		if !isAlphanumeric(lc) {
			l++
			continue
		}
		if !isLowercase(lc) {
			lc = strings.ToLower(lc)
		}
		if !isAlphanumeric(rc) {
			r--
			continue

		}
		if !isLowercase(rc) {
			rc = strings.ToLower(rc)
		}
		if rc != lc {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphanumeric(str string) bool {
	return ("a" <= str && str <= "z") || ("A" <= str && str <= "Z") || ("0" <= str && str <= "9")
}

func isLowercase(str string) bool {
	return str >= "a"
}

// @lc code=end
