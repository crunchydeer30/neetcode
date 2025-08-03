package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 20))
}

func searchMatrix(matrix [][]int, target int) bool {
	lo_r, hi_r := 0, len(matrix)-1

	for lo_r <= hi_r {
		mid_r := lo_r + (hi_r-lo_r)/2
		if target >= matrix[mid_r][0] && target <= matrix[mid_r][len(matrix[mid_r])-1] {
			return searchRow(matrix[mid_r], target)
		} else if target < matrix[mid_r][0] {
			hi_r = mid_r - 1
		} else {
			lo_r = mid_r + 1
		}
	}

	return false
}

func searchRow(row []int, target int) bool {
	s, e := 0, len(row)-1

	for s <= e {
		mid := s + (e-s)/2

		if target < row[mid] {
			e = mid - 1
		} else if target > row[mid] {
			s = mid + 1
		} else {
			return true
		}
	}

	return false
}
