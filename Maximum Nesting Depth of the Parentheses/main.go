package main

func maxDepth(s string) int {
	max := 0
	current := 0
	for _, c := range s {
		if c == '(' {
			current++
			if current > max {
				max = current
			}
		} else if c == ')' {
			current--
		}
	}
	return max
}

func main() {
	maxDepth("1+(2*3)/(2-1)")
}
