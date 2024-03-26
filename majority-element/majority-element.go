package main

/**
 * @input A : Integer array
 *
 * @Output Integer
 * Date: 2024-03-03
 */
import "math"

func majorityElement(A []int) int {
	arrayElementToCountMap := make(map[int]int, 0)
	for i := 0; i < len(A); i++ {
		_, ok := arrayElementToCountMap[A[i]]
		if !ok {
			arrayElementToCountMap[A[i]] = 1
		} else {
			arrayElementToCountMap[A[i]] += 1
		}
		if float64(arrayElementToCountMap[A[i]]) > math.Floor(float64(len(A))/2.0) {
			return A[i]
		}
	}
	return 0
}

func main() {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 1, 1, 1, 1, 1, 1}
	majorityElement(A)
}
