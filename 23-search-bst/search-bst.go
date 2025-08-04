package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	curr := root

	for curr != nil {
		if val > curr.Val {
			curr = curr.Right
		} else if val < curr.Val {
			curr = curr.Left
		} else {
			return curr
		}
	}

	return nil
}
