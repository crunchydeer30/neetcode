package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	steps := []int{}
	stack := []*TreeNode{}
	curr := root
	counter := 0
	solution := root.Val

	for (curr != nil || len(stack) > 0) && len(steps) < k {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		counter++
		if counter == k {
			solution = curr.Val
			break
		}

		curr = curr.Right
	}

	return solution
}
