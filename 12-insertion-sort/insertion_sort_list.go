package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	sortedTail := head
	curr := head.Next

	for curr != nil {
		if curr.Val >= sortedTail.Val {
			sortedTail, curr = curr, curr.Next
			continue
		}

		sortedTail.Next = curr.Next
		node := curr
		curr = sortedTail.Next

		prev := dummy

		for prev.Next.Val < node.Val {
			prev = prev.Next
		}

		node.Next = prev.Next
		prev.Next = node
	}

	return dummy.Next
}
