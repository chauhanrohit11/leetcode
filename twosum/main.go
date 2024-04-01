package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numToIdxMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := numToIdxMap[target-nums[i]]
		if ok {
			return []int{numToIdxMap[target-nums[i]], i}
		}
		numToIdxMap[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
}
