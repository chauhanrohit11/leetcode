package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			k += 1
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == val && nums[j] != val {
				nums[i], nums[j] = nums[j], nums[i]
			}

		}
	}

	return len(nums) - k
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
	fmt.Println("%v", nums)
}
