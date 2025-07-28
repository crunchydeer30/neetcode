package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	k := sort(nums)
	fmt.Println(k)
}

func sort(nums []int) int {
	p1 := 0
	p2 := 0

	for p2 < len(nums) {
		if nums[p1] != nums[p2] {
			p1++
			nums[p1] = nums[p2]
		}
		p2++
	}

	fmt.Println(nums)

	return p1 + 1
}
