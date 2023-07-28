package problems

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit_deleteToSubmit(prices []int) int {
	min := 0
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			min = prices[i]
			continue
		}
		if prices[i] > min {
			profit := prices[i] - min
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			min = prices[i]
		}
	}

	return maxProfit
}

// @lc code=end
