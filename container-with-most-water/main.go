package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	res := 0
	if len(height) <= 1 {
		return 0
	}
	var area int
	for i < j {
		area = (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
		if area > res {
			res = area
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
}
