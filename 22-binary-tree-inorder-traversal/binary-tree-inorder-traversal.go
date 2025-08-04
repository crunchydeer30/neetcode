package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	steps := []int{}
	stack := []*TreeNode{}

	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		steps = append(steps, curr.Val)

		curr = curr.Right
	}

	return steps
}

func inorderTraversalRecursive(root *TreeNode) []int {
	steps := []int{}
	traverseRecursive(root, &steps)
	return steps
}

func traverseRecursive(node *TreeNode, steps *[]int) {
	if node == nil {
		return
	}

	traverseRecursive(node.Left, steps)
	*steps = append(*steps, node.Val)
	traverseRecursive(node.Right, steps)
}
