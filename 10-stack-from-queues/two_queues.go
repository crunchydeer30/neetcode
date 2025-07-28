package main

// type MyStack struct {
// 	q1 *Queue
// 	q2 *Queue
// }

// func Constructor() MyStack {
// 	return MyStack{
// 		q1: &Queue{},
// 		q2: &Queue{},
// 	}
// }

// func (this *MyStack) Push(x int) {
// 	this.q1.Push(x)
// }

// func (this *MyStack) Pop() int {
// 	var shifted int

// 	for this.q1.IsEmpty() == false {
// 		shifted = this.q1.Shift()
// 		if this.q1.IsEmpty() == false {
// 			this.q2.Push(shifted)
// 		}
// 	}

// 	this.q1 = this.q2
// 	this.q2 = &Queue{}

// 	return shifted
// }

// func (this *MyStack) Top() int {
// 	var shifted int

// 	for this.q1.IsEmpty() == false {
// 		shifted = this.q1.Shift()
// 		this.q2.Push(shifted)
// 	}

// 	this.q1 = this.q2
// 	this.q2 = &Queue{}

// 	return shifted
// }

// func (this *MyStack) Empty() bool {
// 	return this.q1.IsEmpty()
// }

// type Queue struct {
// 	items []int
// }

// func (this *Queue) Shift() int {
// 	shifted := this.items[0]
// 	this.items = this.items[1:]
// 	return shifted
// }

// func (this *Queue) IsEmpty() bool {
// 	return len(this.items) == 0
// }

// func (this *Queue) Front() int {
// 	return this.items[0]
// }

// func (this *Queue) Push(val int) {
// 	this.items = append(this.items, val)
// }

// func main() {
// 	stack := Constructor()
// 	stack.Push(1)
// 	stack.Push(2)
// 	stack.Push(3)

// 	fmt.Println("Top", stack.Top(), "q1", stack.q1.items, "q2", stack.q2.items)
// 	fmt.Println("Popped", stack.Pop(), "q1", stack.q1.items, "q2", stack.q2.items)
// 	fmt.Println("Top", stack.Top(), "q1", stack.q1.items, "q2", stack.q2.items)
// }
