package _441_arranging_coins

func arrangeCoins(n int) int {

	if n == 0 {
		return 0
	}
	for i := 1; i <= n; i++ {
		if (1+i)*i/2 <= n && (2+i)*(i+1)/2 > n {
			return i
		}
	}

	return 1
}
