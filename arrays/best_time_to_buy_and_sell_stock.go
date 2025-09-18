package main

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Time: O(n^2)
// Intuition: For each day I can try buying and selling on the day with the maximum profit
// now comparing the maximum profit for each combination we can draw the maximum out of them
func MaxProfitBrute(rates []int) int {
	maxProfit := 0
	for i, r := range rates {
		maxProfitIfBoughtOnIthDay := 0
		for j := i + 1; j < len(rates); j++ {
			profit := rates[j] - r
			if profit > maxProfitIfBoughtOnIthDay {
				maxProfitIfBoughtOnIthDay = profit
			}
		}
		if maxProfitIfBoughtOnIthDay > maxProfit {
			maxProfit = maxProfitIfBoughtOnIthDay
		}
	}
	return maxProfit
}

// Intuition: Minimize the buy value from left and maximize the sell value from right while left > right
func MaxProfitOptimal(rates []int) int {
	buy, sell := rates[0], rates[len(rates)-1]
	left, right := 0, len(rates)-1

	for left <= right {
		if buy > rates[left] {
			buy = rates[left]
			left++
		}
		if sell < rates[right] {
			sell = rates[right]
			right--
		}
	}
	return sell - buy
}
