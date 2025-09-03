package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(hasPathSum(root, 3))
}

// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	targetSum -= root.Val

// 	if root.Left == nil && root.Right == nil {
// 		return targetSum == 0
// 	}

// 	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
// }
