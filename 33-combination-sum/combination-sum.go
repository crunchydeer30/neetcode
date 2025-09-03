package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	path := []int{}
	traverse(candidates, path, &res, target)

	return res
}

func traverse(candidates, path []int, res *[][]int, target int) {
	for i, cand := range candidates {
		if cand > target {
			continue
		}
		if cand == target {
			tmp := make([]int, len(path))
			tmp = append(tmp, cand)
			copy(tmp, path)
			*res = append(*res, tmp)
		}
		if cand < target {
			path = append(path, cand)
			traverse(candidates[i:], path, res, target-cand)
			path = path[:len(path)-1]
		}
	}
}
