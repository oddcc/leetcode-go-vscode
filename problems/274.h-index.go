package problems

import "sort"

/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	sort.Ints(citations)

	h := 0
	for i := len(citations); i > 0; i-- {
		// find the smallest int that greater or equal than i
		t := search(citations, 0, len(citations)-1, i)
		if t == -1 {
			continue
		}
		if len(citations)-1-t+1 >= i {
			h = i
			break
		}
	}

	return h
}

func search(nums []int, start, end, target int) int {
	if start >= end {
		if nums[end] >= target {
			return end
		}
		return -1
	}

	mid := start + ((end - start) / 2)
	if nums[mid] >= target {
		return search(nums, start, mid, target)
	}
	return search(nums, mid+1, end, target)
}

// @lc code=end
