package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
