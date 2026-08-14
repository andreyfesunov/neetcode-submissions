func maxProfit(prices []int) int {
	l, r, m := 0, 1, 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			m = max(m, prices[r] - prices[l])
		} else {
			l = r
		}
		r++
	}

	return m
}
