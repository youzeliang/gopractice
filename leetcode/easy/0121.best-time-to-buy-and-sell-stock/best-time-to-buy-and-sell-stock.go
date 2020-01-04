package _121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	max, temp := 0, 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}

	return max
}
