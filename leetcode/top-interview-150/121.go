package main

func maxProfit(prices []int) int {
	max := 0
	start := 0
	end := len(prices) - 1
	var diff int
	for start < end {
		diff = prices[end] - prices[start]
		if diff >= max {
			max = diff
		}
		if prices[start] > prices[end] {
			start++
		} else {
			end--
		}
	}
	return max
}
