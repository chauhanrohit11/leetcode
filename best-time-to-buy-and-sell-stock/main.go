package main

import "math"

func maxProfit(prices []int) int {

	/* gives tle
	// diff := 0

	// for i := 0; i < len(prices); i++ {
	// 	for j := i + 1; j < len(prices); j++ {
	// 		if prices[j]-prices[i] > diff {
	// 			diff = prices[j] - prices[i]
	// 		}
	// 	}
	// }
	// return diff
	*/

	maxFromRight := make([]int, 0)
	minFromLeft := make([]int, 0)
	minTillNow := math.MaxInt64
	maxTillNow := math.MinInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < minTillNow {
			minTillNow = prices[i]
		}
		minFromLeft = append(minFromLeft, minTillNow)
	}
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxTillNow {
			maxTillNow = prices[i]
		}
		maxFromRight = append(maxFromRight, maxTillNow)
	}
	for i, j := 0, len(maxFromRight)-1; i < j; i, j = i+1, j-1 {
		maxFromRight[i], maxFromRight[j] = maxFromRight[j], maxFromRight[i]
	}
	res := math.MinInt64
	for i := 0; i < len(prices)-1; i++ {
		if minFromLeft[i] < maxFromRight[i] && maxFromRight[i]-minFromLeft[i] > res {
			res = maxFromRight[i] - minFromLeft[i]
		}
	}
	if res == math.MinInt64 {
		return 0
	}
	return res
}

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}
