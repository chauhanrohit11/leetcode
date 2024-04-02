package main

import "fmt"

func summaryRanges(nums []int) []string {
	j := 1
	res := make([]string, 0)
	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}
	for i := 0; i < len(nums); i++ {
		k := i
		if i == len(nums)-1 {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else if j < len(nums) && k < len(nums) && nums[k]+1 == nums[j] {
			for j < len(nums) && nums[k]+1 == nums[j] {
				j++
				k++
			}
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
			i = k
			j++
		} else if nums[i] != nums[j] {
			res = append(res, fmt.Sprintf("%d", nums[i]))
			j++
		}
	}
	return res
}

func main() {
	// fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
