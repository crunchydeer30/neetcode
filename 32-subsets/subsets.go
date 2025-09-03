package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			newSubset := make([]int, len(res[i]))
			copy(newSubset, res[i])
			newSubset = append(newSubset, num)
			res = append(res, newSubset)
		}
	}

	return res
}
