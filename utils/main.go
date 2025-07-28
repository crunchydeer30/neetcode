package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	arr := []int{}
	head := createLinkedList(arr)
	if head == nil {
		fmt.Println(nil)
	}
	printLinkedList(head)
}
