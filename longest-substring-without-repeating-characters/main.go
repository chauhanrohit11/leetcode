package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charToIdxMap := make(map[byte]int, 0)
	left := 0
	right := 0
	if len(s) == 0 {
		return 0
	}
	res := 1
	for i := 0; i < len(s); i++ {
		_, ok := charToIdxMap[s[i]]
		if ok {
			if charToIdxMap[s[i]] >= left {
				left = charToIdxMap[s[i]] + 1
			}
		}
		charToIdxMap[s[i]] = i
		if res < right-left+1 {
			res = right - left + 1
		}
		right++
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf")) // 3
}
