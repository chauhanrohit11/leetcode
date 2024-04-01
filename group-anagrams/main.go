package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	// sort the strings and use the sorted string as the key to group the anagrams
	// the value of the map will be the list of anagrams
	// the time complexity is O(n * k * log(k)) where n is the number of strings and k is the length of the string
	// the space complexity is O(n * k)
	sortedStrToAnagramsMap := make(map[string][]string, 0)
	for i := 0; i < len(strs); i++ {
		sortedStr := []byte(strs[i])
		sort.Slice(sortedStr, func(i, j int) bool {
			return sortedStr[i] < sortedStr[j]
		})
		sStr := string(sortedStr)
		_, ok := sortedStrToAnagramsMap[sStr]
		if !ok {
			sortedStrToAnagramsMap[sStr] = make([]string, 0)
		}
		sortedStrToAnagramsMap[sStr] = append(sortedStrToAnagramsMap[sStr], strs[i])
	}
	res := make([][]string, 0)
	for _, anagrams := range sortedStrToAnagramsMap {
		res = append(res, anagrams)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // [["bat"],["nat","tan"],["ate","eat","tea"]]
}
