package problems
/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	var profit = make([][]int, len(prices))
	for i :=0; i < len(prices); i++ {
		profit[i] = make([]int, 2)
	}

	profit[0][0] = 0
	profit[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		profitIfNoStock := profit[i-1][0]
		if profitIfNoStock < profit[i-1][1] + prices[i] {
			profitIfNoStock = profit[i-1][1] + prices[i]
		}
		profit[i][0] = profitIfNoStock

		profitIfHoldStock := profit[i-1][0] - prices[i]
		if profitIfHoldStock < profit[i-1][1] {
			profitIfHoldStock = profit[i-1][1]
		}
		profit[i][1] = profitIfHoldStock
	}

	return profit[len(prices) - 1][0]
}
// @lc code=end

