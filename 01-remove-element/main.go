package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	k := removeElement(nums, 3)
	fmt.Println(nums, k)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	p1 := 0
	p2 := len(nums) - 1

	for p1 < p2 {
		if nums[p2] == val {
			p2--
			continue
		}

		if nums[p1] == val {
			nums[p1] = nums[p2]
			p1++
			p2--
			continue
		}

		p1++
	}

	return p1 + 1
}
