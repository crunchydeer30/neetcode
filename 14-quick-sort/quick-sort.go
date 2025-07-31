package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{6, 2, 4, 1, 3}, 0, 4))
}

func quickSort(arr []int, s, e int) []int {
	if e-s+1 <= 1 {
		return arr
	}

	pivot := e
	left, i := s, s

	for i < pivot {
		if arr[i] < arr[pivot] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
		i++
	}

	arr[pivot], arr[left] = arr[left], arr[pivot]

	quickSort(arr, s, left-1)
	quickSort(arr, left+1, e)

	return arr
}
