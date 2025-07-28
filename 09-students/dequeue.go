package main

type Deque struct {
	items []int
}

func (this *Deque) Push(val int) {
	this.items = append(this.items, val)
}

func (this *Deque) Pop() (int, bool) {
	if len(this.items) == 0 {
		return 0, false
	}

	idx := len(this.items) - 1
	popped := this.items[idx]
	this.items = this.items[:idx]
	return popped, true
}

func (this *Deque) Unshift(val int) {
	this.items = append([]int{val}, this.items...)
}

func (this *Deque) Shift() (int, bool) {
	if len(this.items) == 0 {
		return 0, false
	}

	shifted := this.items[0]
	this.items = this.items[1:]

	return shifted, true
}

func (this *Deque) Top() (int, bool) {
	if len(this.items) == 0 {
		return 0, false
	}
	return this.items[len(this.items)-1], true
}

func (this *Deque) Front() (int, bool) {
	if len(this.items) == 0 {
		return 0, false
	}
	return this.items[0], true
}

// func countStudents(students []int, sandwiches []int) int {
// 	studentsQueue := Deque{items: students}
// 	sandwichesQueue := Deque{items: sandwiches}
// 	iterations := 0

// 	for len(sandwichesQueue.items) != 0 && iterations < len(sandwichesQueue.items) {
// 		student, _ := studentsQueue.Shift()
// 		sandwich, _ := sandwichesQueue.Front()

// 		if student == sandwich {
// 			sandwichesQueue.Shift()
// 			iterations = 0
// 		} else {
// 			studentsQueue.Push(student)
// 			iterations++
// 		}
// 	}

// 	return len(studentsQueue.items)
// }
