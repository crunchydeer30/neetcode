package main

import (
	"fmt"
	"math"
)

func main() {
	piles := []int{30, 11, 23, 4, 20}
	hours := 5

	fmt.Println(minEatingSpeed(piles, hours))
}

func minEatingSpeed(piles []int, h int) int {
	low, high := 1, maxPile(piles)

	for low < high {
		mid := low + (high-low)/2
		if isPossible(piles, h, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func maxPile(piles []int) int {
	max := piles[0]
	for _, pile := range piles {
		max = int(math.Max(float64(max), float64(pile)))
	}
	return max
}

func isPossible(piles []int, availableHours, speed int) bool {
	hoursSpent := 0

	for _, pile := range piles {
		hoursSpent += int(math.Ceil(float64(pile) / float64(speed)))
		if hoursSpent > availableHours {
			return false
		}
	}

	return hoursSpent <= availableHours
}
