package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := createLinkedList([]int{1, 2, 0, -1, 5, 88, -9})
	printLinkedList(head)
	head = insertionSortList(head)
	printLinkedList(head)
}

func insertionSortList(head *ListNode) *ListNode {
	curr := head.Next
	prev := head

	for curr != nil {
		if curr.Val < prev.Val {
			node := curr
			curr = curr.Next
			prev.Next = curr

			if node.Val <= head.Val {
				node.Next = head
				head = node
				continue
			}

			currInner := head.Next
			prevInner := head

			for currInner != curr {
				if node.Val <= currInner.Val {
					node.Next = currInner
					prevInner.Next = node
					break
				}
				prevInner = currInner
				currInner = currInner.Next
			}

		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return head
}

func createLinkedList(arr []int) *ListNode {
	var head *ListNode = nil
	var curr *ListNode = head

	for _, val := range arr {
		if head == nil {
			node := &ListNode{Val: val, Next: nil}
			head = node
			curr = head
			continue
		}

		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		curr.Next = node
		curr = curr.Next
	}

	return head
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
