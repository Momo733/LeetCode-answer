package _22_买卖股票的最佳时机_II

func MaxProfit(prices []int) int { //贪心算法，一有利润立即卖出
	var res int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res = prices[i] - prices[i-1] +res
		}
	}
	return res
}
