package main

func binarySearchRecursive(nums []int, s, e, target int) int {
	if s > e {
		return -1
	}

	mid := (s + e) / 2
	if target < nums[mid] {
		return binarySearchRecursive(nums, s, mid-1, target)
	} else if target > nums[mid] {
		return binarySearchRecursive(nums, mid+1, e, target)
	} else {
		return mid
	}
}
