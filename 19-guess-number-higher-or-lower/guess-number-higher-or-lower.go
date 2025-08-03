package main

func guessNumber(n int) int {
	low, high := 0, n

	for low <= high {
		mid := low + (high-low)/2

		switch guess(mid) {
		case -1:
			high = mid - 1
		case 0:
			return mid
		case 1:
			low = mid + 1
		}
	}

	return -1
}

func guess(num int) int {
	return 0
}
