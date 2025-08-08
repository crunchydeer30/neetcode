package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	// curr := root
	// for curr != nil {
	// 	fmt.Println(curr.Val)
	// 	curr = curr.Left
	// }
	// curr = root
	// for curr != nil {
	// 	fmt.Println(curr.Val)
	// 	curr = curr.Right
	// }
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderHash := make(map[int]int)
	for idx, val := range inorder {
		inorderHash[val] = idx
	}

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	curr := root
	stack := []*TreeNode{root}

	for i := 1; i < len(preorder); i++ {
		inorderIdx, _ := inorderHash[preorder[i]]
		inorderRootIdx, _ := inorderHash[curr.Val]
		fmt.Println(preorder[i], curr)
		if inorderIdx < inorderRootIdx {
			curr.Left = &TreeNode{Val: preorder[i]}
			curr = curr.Left
			stack = append(stack, curr)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderRootIdx] {
				curr = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderRootIdx++
			}
			curr.Right = &TreeNode{Val: preorder[i]}
			curr = curr.Right
			stack = append(stack, curr)
		}
	}

	return root
}
