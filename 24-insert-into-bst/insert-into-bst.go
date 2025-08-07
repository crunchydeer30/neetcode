package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}

	if root == nil {
		return node
	}

	curr := root

	for {
		if node.Val < curr.Val {
			if curr.Left == nil {
				curr.Left = node
				return root
			}
			curr = curr.Left
		} else if node.Val > curr.Val {
			if curr.Right == nil {
				curr.Right = node
				return root
			}
			curr = curr.Right
		} else {
			return root
		}
	}
}

func insertIntoBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insertIntoBSTRecursive(root.Right, val)
	} else if val < root.Val {
		root.Left = insertIntoBSTRecursive(root.Left, val)
	}

	return root
}
