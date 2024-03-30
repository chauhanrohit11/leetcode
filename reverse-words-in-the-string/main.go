package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := ""
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			res += words[i] + " "
		}
	}
	res = strings.TrimRight(res, " ")
	return res
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
