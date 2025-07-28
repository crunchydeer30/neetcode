package main

import "fmt"

func climbStairs(n int) int {
	m := make(map[int]int)

	return takeSteps(n, m)
}

func takeSteps(n int, m map[int]int) int {
	if n == 1 || n == 2 {
		return n
	}

	if val, exists := m[n]; exists {
		return val
	}

	result := takeSteps(n-1, m) + takeSteps(n-2, m)
	m[n] = result
	return result
}

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}
