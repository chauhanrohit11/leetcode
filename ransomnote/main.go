package main

func canConstruct(ransomNote string, magazine string) bool {
	charToCountMap := make(map[string]int, 0)
	for i := 0; i < len(magazine); i++ {
		_, ok := charToCountMap[magazine[i:i+1]]
		if !ok {
			charToCountMap[magazine[i:i+1]] = 1
		} else {
			charToCountMap[magazine[i:i+1]] += 1
		}
	}
	for i := 0; i < len(ransomNote); i++ {
		val, ok := charToCountMap[ransomNote[i:i+1]]
		if !ok || val == 0 {
			return false
		} else {
			charToCountMap[ransomNote[i:i+1]] -= 1
		}
	}
	return true
}

func main() {
	// println(canConstruct("a", "b"))    // false
	// println(canConstruct("aa", "ab"))  // false
	println(canConstruct("aab", "baa")) // true
}
