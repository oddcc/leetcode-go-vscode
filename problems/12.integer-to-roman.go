package problems

import "strings"

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	var sb strings.Builder

	thousand := num / 1000
	num -= thousand * 1000
	for i := 0; i < thousand; i++ {
		sb.WriteByte('M')
	}

	hundred := num / 100
	num -= hundred * 100
	if hundred == 9 {
		sb.WriteString("CM")
	} else if hundred == 4 {
		sb.WriteString("CD")
	} else if hundred >= 5 {
		sb.WriteString("D")
		for i := 0; i < hundred-5; i++ {
			sb.WriteByte('C')
		}
	} else {
		for i := 0; i < hundred; i++ {
			sb.WriteByte('C')
		}
	}

	ten := num / 10
	num -= ten * 10
	if ten == 9 {
		sb.WriteString("XC")
	} else if ten == 4 {
		sb.WriteString("XL")
	} else if ten >= 5 {
		sb.WriteString("L")
		for i := 0; i < ten-5; i++ {
			sb.WriteByte('X')
		}
	} else {
		for i := 0; i < ten; i++ {
			sb.WriteByte('X')
		}
	}

	if num == 9 {
		sb.WriteString("IX")
	} else if num == 4 {
		sb.WriteString("IV")
	} else if num >= 5 {
		sb.WriteString("V")
		for i := 0; i < num-5; i++ {
			sb.WriteByte('I')
		}
	} else {
		for i := 0; i < num; i++ {
			sb.WriteByte('I')
		}
	}

	return sb.String()
}

// @lc code=end
