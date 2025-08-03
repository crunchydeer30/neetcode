package main

import (
	"math/rand"
)

// func main() {
// 	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
// }

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k, s, e int) int {
	if s == e {
		return nums[s]
	}

	p := s + rand.Intn(e-s+1)
	nums[p], nums[e] = nums[e], nums[p]

	pivot := e
	left, i := s, s

	for i < pivot {
		if nums[i] < nums[pivot] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
		i++
	}

	nums[pivot], nums[left] = nums[left], nums[pivot]
	pivotKthIndex := len(nums) - left

	if pivotKthIndex == k {
		return nums[left]
	}

	if pivotKthIndex < k {
		return quickSelect(nums, k, s, left-1)
	}

	return quickSelect(nums, k, left+1, e)
}
