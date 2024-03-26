package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	st := nonAlphanumericRegex.ReplaceAllString(s, "")
	st = strings.Replace(st, " ", "", -1)
	pt := ""
	for _, v := range st {
		pt = string(v) + pt
	}
	return strings.ToLower(st) == strings.ToLower(pt)
}

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}
