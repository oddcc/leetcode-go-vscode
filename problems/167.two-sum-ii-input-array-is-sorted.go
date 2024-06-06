package problems

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	i1 := 0
	i2 := 0

	for i1 < len(numbers)-1 {
		i2 = i1 + 1
		if result, found := biSearch(numbers, i2, len(numbers)-1, target-numbers[i1]); found {
			return []int{i1 + 1, result + 1}
		}
		i1++
	}

	return []int{i1 + 1, i2 + 1}
}

func biSearch(numbers []int, l, r, tar int) (result int, success bool) {
	if l >= r {
		if numbers[l] == tar {
			return l, true
		} else {
			return -1, false
		}
	}

	mid := l + (r - l) / 2
	if numbers[mid] >= tar {
		return biSearch(numbers, l, mid, tar)
	} else {
		return biSearch(numbers, mid+1, r, tar)
	}
}

// @lc code=end
