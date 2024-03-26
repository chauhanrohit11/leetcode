package main

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	i := 0
	currentElementOccurence := 1
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] || (nums[i] == nums[j] && currentElementOccurence < 2) {
			if nums[i] == nums[j] {
				currentElementOccurence += 1
			} else {
				currentElementOccurence = 1
			}
			i++
			nums[i] = nums[j]
		} else if nums[i] == nums[j] {
			currentElementOccurence++
		}
	}
	return i + 1
}

func main() {
	removeDuplicates([]int{1, 1, 1, 2, 3, 3, 3, 4, 4, 5, 5, 5})
}
