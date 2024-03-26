package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// shift nums 1
	// merge nums 2
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < len(nums1); i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	index := m - 1
	for i := len(nums1) - 1; index >= 0; i-- {
		nums1[i] = nums1[index]
		nums1[index] = 0
		index -= 1
	}
	j := len(nums1) - m
	k := 0
	for i := 0; i < len(nums1); i++ {
		if j < len(nums1) && k < n && nums1[j] < nums2[k] {
			nums1[i] = nums1[j]
			j += 1
		} else if k < n && j < len(nums1) && nums1[j] >= nums2[k] {
			nums1[i] = nums2[k]
			k += 1
		} else if j >= len(nums1) && k < n {
			nums1[i] = nums2[k]
			k += 1
		} else if k >= n && j < len(nums1) {
			nums1[i] = nums1[j]
			j += 1
		}
	}
}

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 0}
	m := 5
	nums2 := []int{3}
	n := 1
	merge(nums1, m, nums2, n)
	fmt.Printf("%v", nums1)
}
