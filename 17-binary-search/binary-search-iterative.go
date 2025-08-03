package main

import "fmt"

func main() {
	nums := []int{0, 1, 3}
	fmt.Println(binarySearchIterative(nums, 3))
}

func binarySearchIterative(nums []int, target int) int {
	s, e := 0, len(nums)-1

	for s <= e {
		mid := s + (e-s)/2

		if target < nums[mid] {
			e = mid - 1
		} else if target > nums[mid] {
			s = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
