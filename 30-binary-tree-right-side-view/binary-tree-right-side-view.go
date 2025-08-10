package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	dq := []*TreeNode{root}

	if root == nil {
		return result
	}

	for len(dq) > 0 {
		levelSize := len(dq)
		rightMostValue := dq[0].Val

		for i := 0; i < levelSize; i++ {
			node := dq[0]
			dq = dq[1:]
			rightMostValue = node.Val
			if node.Left != nil {
				dq = append(dq, node.Left)
			}
			if node.Right != nil {
				dq = append(dq, node.Right)
			}
		}

		result = append(result, rightMostValue)
	}

	return result
}
