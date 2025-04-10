package main

import "fmt"

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{inStack: []int{}, outStack: []int{}}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) shiftStacks() {
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			n := len(q.inStack)
			// Pop from inStack and Push to outStack
			q.outStack = append(q.outStack, q.inStack[n-1])
			q.inStack = q.inStack[:n-1]
		}
	}
}

func (q *MyQueue) Pop() int {
	q.shiftStacks()
	val := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return val
}

func (q *MyQueue) Peek() int {
	q.shiftStacks()
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println("Peek:", q.Peek())    // Output: 1
	fmt.Println("Pop:", q.Pop())      // Output: 1
	fmt.Println("Empty?:", q.Empty()) // Output: false
}
