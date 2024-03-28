package main

// two pointer approach
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
func main() {
	isSubsequence("abc", "ahbgdc")
	isSubsequence("axc", "ahbgdc")
}
