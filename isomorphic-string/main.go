package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sMap := make(map[string]int, 0)
	tMap := make(map[string]int, 0)
	for i := 0; i < len(s); i++ {
		_, ok := sMap[s[i:i+1]]
		if !ok {
			sMap[s[i:i+1]] = i
		}
		_, ok = tMap[t[i:i+1]]
		if !ok {
			tMap[t[i:i+1]] = i
		}
		if sMap[s[i:i+1]] != tMap[t[i:i+1]] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isIsomorphic("egg", "add")) // true
}
