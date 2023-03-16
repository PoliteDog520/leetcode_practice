package main

import "log"

func maxProfit(prices []int) int {
	var maxSalePrice, minBuyPrice int
	if len(prices) < 0 {
		return 0
	} else {
		// 初始化
		minBuyPrice = prices[0]
	}
	for i := range prices {
		if prices[i] < minBuyPrice {
			minBuyPrice = prices[i]
		}
		if prices[i]-minBuyPrice > maxSalePrice {
			maxSalePrice = prices[i] - minBuyPrice
		}
	}
	return maxSalePrice
}

func main() {
	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Println(maxProfit([]int{7, 6, 4, 3, 2, 1}))
	log.Println(maxProfit([]int{2, 2, 2, 2, 2, 2}))
}
