package leetCode

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.
Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

// Approach 1:
/*
1. 先求只买卖一次的时候可以获得的最大值，从前->后，profits[i] = max(profits[i-1], prices[i] - buyPrice)
2. 从后->前，求买卖两次的最大值，maxProfit = max(maxProfit, salePrice-price[i] + profits[i-1]).
最终结果为买卖一次还是两次的最大值。
*/

/*
func max(x, y int) int{
	if x > y{
		return x
	}
	return y
}
*/

func maxProfit123(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profits := make([]int, len(prices))
	buyPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		}
		profits[i] = max(profits[i-1], prices[i]-buyPrice)
	}

	maxProfit := profits[len(profits)-1]
	salePrice := prices[len(prices)-1]
	for i := len(prices) - 2; i > 0; i-- {
		if prices[i] > salePrice {
			salePrice = prices[i]
		}
		maxProfit = max(maxProfit, salePrice-prices[i]+profits[i-1])
	}
	return maxProfit
}

// Approach 2:
/*
分别求[0, i), [i, len)的区间最大值，然后求两者和的最大值
Note: 多重循环，速度很慢
*/

func calMaxProfit(prices []int, beg, end int) int {
	buyPrice := prices[beg]
	maxProfit := 0
	for i := beg + 1; i < end; i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-buyPrice)
		}
	}
	return maxProfit
}

func maxProfit_123(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		maxProfit = max(maxProfit, calMaxProfit(prices, 0, i)+calMaxProfit(prices, i, len(prices)))
	}
	return maxProfit
}
