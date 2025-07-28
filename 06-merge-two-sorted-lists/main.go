package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	list1 := createLinkedList([]int{1, 2, 4})
	list2 := createLinkedList([]int{1, 3, 4})
	fmt.Println("list1")
	printLinkedList(list1)
	fmt.Println("list2")
	printLinkedList(list2)

	merged := mergeTwoLists(list1, list2)
	fmt.Println("merged")
	printLinkedList(merged)
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
		fmt.Println(head.Val)
		head = head.Next
	}
}
