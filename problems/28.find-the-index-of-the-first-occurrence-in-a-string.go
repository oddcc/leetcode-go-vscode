package problems

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		hI := i
		found := false
		for j := 0; j < len(needle) && hI < len(haystack); j++ {
			if haystack[hI] != needle[j] {
				break
			}
			hI++
			if j == len(needle)-1 {
				found = true
			}
		}
		if found {
			return i
		}
	}

	return -1
}

// @lc code=end
