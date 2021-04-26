package array2

func MaxProfit(prices []int) int {
	res := 0
	for i, v := range prices {
		if i > 0 && v > prices[i-1] {
			res += v - prices[i-1]
		}
	}
	return res
}
