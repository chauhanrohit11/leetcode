package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
func main() {
	fmt.Println(climbStairs(2)) // 2
	fmt.Println(climbStairs(5)) // 8
	fmt.Println(climbStairs(6)) // 13
}
