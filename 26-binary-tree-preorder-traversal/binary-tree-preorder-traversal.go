package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	steps := []int{}
	traverse(root, &steps)
	return steps
}

func traverse(root *TreeNode, steps *[]int) {
	if root == nil {
		return
	}

	*steps = append(*steps, root.Val)
	traverse(root.Left, steps)
	traverse(root.Right, steps)
}

func preorderTraversalIterative(root *TreeNode) []int {
	steps := []int{}
	stack := []*TreeNode{}

	curr := root

	for curr != nil || len(stack) > 0 {
		if curr == nil {
			curr = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, curr)
			steps = append(steps, curr.Val)
			curr = curr.Left
		}
	}

	return steps
}
