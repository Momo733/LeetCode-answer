package _21_买卖股票的最佳时机


func MaxProfit(prices []int) int {
	var res int
	for i, v := range prices { //买
		for j, k := range prices { //卖
			if i > j {
				continue
			}
			if k-v > res {
				res = k - v
			}
		}
	}
	return res
}