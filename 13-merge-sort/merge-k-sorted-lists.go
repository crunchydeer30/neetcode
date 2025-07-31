package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Next = list2
			break
		}

		if list2 == nil {
			curr.Next = list1
			break
		}

		if list1.Val > list2.Val {
			tmp := list2.Next
			curr.Next = list2
			list2 = tmp
		} else {
			tmp := list1.Next
			curr.Next = list1
			list1 = tmp
		}

		curr = curr.Next
	}

	return head.Next
}
