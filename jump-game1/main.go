package main

func canJump(nums []int) bool {
	maxPositionFromRight := make([]int, 0)
	if len(nums) == 0 {
		return false
	}
	if len(nums) <= 1 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		maxToGoRight := 0
		if i+nums[i] > maxToGoRight && nums[i] > 0 {
			maxToGoRight = i + nums[i]
		}
		maxPositionFromRight = append(maxPositionFromRight, maxToGoRight)
	}
	currentMaxToRight := maxPositionFromRight[0]
	for i := 0; i <= currentMaxToRight; i++ {
		if maxPositionFromRight[i] >= len(nums)-1 {
			return true
		}
		if maxPositionFromRight[i] > currentMaxToRight {
			currentMaxToRight = maxPositionFromRight[i]
		}
	}
	return false
}
func main() {
	// canJump([]int{3, 2, 1, 0, 4})
	// canJump([]int{2, 3, 1, 1, 4})
	canJump([]int{1, 2, 3})
}
