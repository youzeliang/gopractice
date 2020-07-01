package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

func maxProfit(prices []int, fee int) int {
	empty := 0
	hold := -1 << 63

	for _, p := range prices {
		temp := empty
		empty = max(empty, hold+p)
		hold = max(hold, temp-p-fee)
	}

	return empty
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
