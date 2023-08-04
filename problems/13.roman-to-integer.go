package problems

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	res := 0

	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	for i := len(s) - 1; i>= 0; i-- {
		switch c:= s[i]; c {
		case 'I':
			res += 1
		case 'V':
			if i-1>=0 && s[i-1] == 'I' {
				res += 4
				i--
			} else {
				res += 5
			}
		case 'X':
			if i-1>=0 && s[i-1] == 'I' {
				res += 9
				i--
			} else {
				res += 10
			}
		case 'L':
			if i-1>=0 && s[i-1] == 'X' {
				res += 40
				i--
			} else {
				res += 50
			}
		case 'C':
			if i-1>=0 && s[i-1] == 'X' {
				res += 90
				i--
			} else {
				res += 100
			}
		case 'D':
			if i-1>=0 && s[i-1] == 'C' {
				res += 400
				i--
			} else {
				res += 500
			}
		case 'M':
			if i-1>=0 && s[i-1] == 'C' {
				res += 900
				i--
			} else {
				res += 1000
			}
		}
	}

	return res
}
// @lc code=end

