package main

import "fmt"

type MyStack struct {
	q *Queue
}

func Constructor() MyStack {
	return MyStack{
		q: &Queue{},
	}
}

func (this *MyStack) Push(x int) {
	this.q.Push(x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < len(this.q.items)-1; i++ {
		shifted := this.q.Shift()
		this.q.Push(shifted)
	}

	return this.q.Shift()
}

func (this *MyStack) Top() int {
	var shifted int

	for i := 0; i <= len(this.q.items)-1; i++ {
		shifted = this.q.Shift()
		this.q.Push(shifted)
	}

	return shifted
}

func (this *MyStack) Empty() bool {
	return this.q.IsEmpty()
}

type Queue struct {
	items []int
}

func (this *Queue) Shift() int {
	shifted := this.items[0]
	this.items = this.items[1:]
	return shifted
}

func (this *Queue) IsEmpty() bool {
	return len(this.items) == 0
}

func (this *Queue) Front() int {
	return this.items[0]
}

func (this *Queue) Push(val int) {
	this.items = append(this.items, val)
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Top", stack.Top(), "q1", stack.q.items)
	fmt.Println("Popped", stack.Pop(), "q1", stack.q.items)
	fmt.Println("Top", stack.Top(), "q1", stack.q.items)
}
