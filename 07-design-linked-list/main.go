package main

import "fmt"

type MyLinkedList struct {
	left  *Node
	right *Node
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

func Constructor() MyLinkedList {
	list := MyLinkedList{
		left:  &Node{},
		right: &Node{},
	}

	list.left.next = list.right
	list.right.prev = list.left

	return list
}

func (this *MyLinkedList) Get(index int) int {
	val := -1
	curr := this.left.next
	i := 0

	for curr != this.right && i <= index {
		if i == index {
			return curr.val
		}
		i++
		curr = curr.next
	}

	return val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := Node{val: val, prev: this.left, next: this.left.next}
	this.left.next.prev = &node
	this.left.next = &node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := Node{val: val, next: this.right, prev: this.right.prev}
	this.right.prev.next = &node
	this.right.prev = &node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	curr := this.left.next
	i := 0

	for curr != nil && i <= index {
		if i == index {
			node := Node{val: val, next: curr, prev: curr.prev}
			curr.prev.next = &node
			curr.prev = &node
		}
		curr = curr.next
		i++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	curr := this.left.next
	i := 0

	for curr != this.right && i <= index {
		if i == index {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
		}
		curr = curr.next
		i++
	}
}

func printLinkedList(head *Node) {
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}

func main() {
	list := Constructor()

	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	list.AddAtIndex(3, 4)
	fmt.Println("----")
	printLinkedList(list.left)
	fmt.Println("----")

	list.DeleteAtIndex(0)
	list.DeleteAtIndex(2)
	fmt.Println("----")
	printLinkedList(list.left)
	fmt.Println("----")
}
