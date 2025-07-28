package main

import "fmt"

type node struct {
	val int
	min int
}

type MinStack struct {
	values []node
}

func Constructor() MinStack {
	return MinStack{
		values: []node{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.values) == 0 {
		this.values = append(this.values, node{val: val, min: val})
	} else {
		this.values = append(this.values, node{val: val, min: min(this.GetMin(), val)})
	}
}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1].val
}

func (this *MinStack) GetMin() int {
	return this.values[len(this.values)-1].min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minStack := Constructor()

	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}
