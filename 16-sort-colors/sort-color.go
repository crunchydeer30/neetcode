package main

import "fmt"

func main() {
	colors := []int{}
	sortColors(colors)
	fmt.Println(colors)
}

func sortColors(nums []int) {
	buckets := [3]int{0, 0, 0}

	for _, color := range nums {
		if color >= 0 && color <= 2 {
			buckets[color]++
		}
	}

	i := 0
	for n := range buckets {
		for range buckets[n] {
			nums[i] = n
			i++
		}
	}
}
