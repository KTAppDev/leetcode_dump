package main

func main() {
	prices := []int{7, 1, 5, 10, 6, 14}
	maxProfit(prices)
}

// maxProfit calculates the maximum profit that can be obtained by buying and selling
// a stock represented by daily prices.
func maxProfit(prices []int) int {
	// Initialize the minimum price to the first element in the prices array
	minPrice := prices[0]

	// Initialize the maximum profit to zero
	maxProfit := 0

	// Go through the prices array starting from the second element
	// We iterate through the array starting from the second element because we are comparing each price with the previous one.
	for i := 1; i < len(prices); i++ {
		// Update the minimum price if the current price is lower
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Calculate the current profit by subtracting the minimum price from the current price
			profit := prices[i] - minPrice

			// Update the maximum profit if the current profit is greater
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	// Return the maximum profit obtained
	println(maxProfit)
	return maxProfit
}
