package main

import "fmt"

type ListNodeGeneric[T any] struct {
	Val  T
	Next *ListNodeGeneric[T]
}

func createGenericLinkedList[T any](arr []T) *ListNodeGeneric[T] {
	var head *ListNodeGeneric[T] = nil
	var curr *ListNodeGeneric[T] = head

	for _, val := range arr {
		if head == nil {
			node := &ListNodeGeneric[T]{Val: val, Next: nil}
			head = node
			curr = head
			continue
		}

		node := &ListNodeGeneric[T]{
			Val:  val,
			Next: nil,
		}
		curr.Next = node
		curr = curr.Next
	}

	return head
}

func printGenericLinkedList[T any](head *ListNodeGeneric[T]) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
