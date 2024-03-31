package main

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	sum := 0
	minLen := len(nums) + 1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func main() {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}) // 2
}
