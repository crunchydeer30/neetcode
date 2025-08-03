package main

func firstBadVersion(n int) int {
	low, high := 1, n
	for low < high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func isBadVersion(num int) bool {
	return false
}
