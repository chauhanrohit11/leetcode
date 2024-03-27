package main

func twoSum(numbers []int, target int) []int {
	var sum int
	p1 := 0
	p2 := len(numbers) - 1
	for p1 < p2 {
		sum = numbers[p1] + numbers[p2]
		if sum == target {
			return []int{p1 + 1, p2 + 1}
		} else if sum < target {
			p1++
		} else {
			p2--
		}
	}
	return []int{-1, -1}
}
func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
