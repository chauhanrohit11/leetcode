package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i < len(nums)+1; i++ {
		for j := 0; j < target+1; j++ {
			dp[i][j] = 0
		}
	}
	// initialize the array
	for i := 0; i < len(nums)+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < target+1; j++ {
			if dp[i-1][j] == 1 {
				dp[i][j] = 1
			} else if j >= nums[i-1] && dp[i-1][j-nums[i-1]] == 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[len(nums)][target] == 1
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 4})) // true
}
