package main

import "fmt"

func threeSum(nums []int) [][]int {
	valToIdxToExistsMap := make(map[int]map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := valToIdxToExistsMap[nums[i]]
		if !ok {
			_, ok := valToIdxToExistsMap[nums[i]][i]
			if !ok {
				valToIdxToExistsMap[nums[i]] = make(map[int]bool, 0)
				valToIdxToExistsMap[nums[i]][i] = true
			} else {
				valToIdxToExistsMap[nums[i]][i] = true
			}
		}
	}
	uniqueTriplets := make(map[string]bool)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			_, ok := valToIdxToExistsMap[-nums[i]-nums[j]]
			if ok {
				for k := range valToIdxToExistsMap[-nums[i]-nums[j]] {
					if k != i && k != j {
						a, b, c := sortTriplet(nums[i], nums[j], nums[k])
						tripletStr := fmt.Sprintf("%d%d%d", a, b, c)
						_, ok := uniqueTriplets[tripletStr]
						if !ok {
							res = append(res, []int{nums[i], nums[j], nums[k]})
							uniqueTriplets[tripletStr] = true
						}
						break
					}
				}
			}
		}
	}

	return res
}

func sortTriplet(i int, j int, k int) (int, int, int) {
	if i > j {
		i, j = j, i
	}
	if j > k {
		j, k = k, j
	}
	if i > j {
		i, j = j, i
	}
	return i, j, k
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
}
