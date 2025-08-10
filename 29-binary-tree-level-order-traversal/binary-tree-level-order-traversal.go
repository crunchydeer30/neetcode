package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	steps := [][]int{}
	queue := []*TreeNode{root}

	if root == nil {
		return steps
	}

	for len(queue) > 0 {
		levelSize := len(queue)
		curLevelSteps := make([]int, 0, levelSize)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			curLevelSteps = append(curLevelSteps, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if len(curLevelSteps) > 0 {
			steps = append(steps, curLevelSteps)
		}
	}

	return steps
}
