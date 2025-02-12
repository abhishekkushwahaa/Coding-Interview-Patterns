package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(value int) {
	newNode := &Node{data: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	currenNode := ll.head
	for currenNode.next != nil {
		currenNode = currenNode.next
	}

	currenNode.next = newNode
}

// Logic of this Question
func (ll *LinkedList) DetectLoop() bool {
	slow, fast := ll.head, ll.head

	for fast.next != nil && fast != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

// Only for testing
func (ll *LinkedList) CreateLoop(position int) {
	if ll.head == nil {
		return
	}

	loopStart, temp := ll.head, ll.head
	for i := 0; temp.next != nil; i++ {
		if i == position {
			loopStart = temp
		}
		temp = temp.next
	}
	temp.next = loopStart
}

func (ll *LinkedList) Print() {
	currentNode := ll.head
	for currentNode != nil {
		fmt.Print(currentNode.data, "->")
		currentNode = currentNode.next
	}
	fmt.Println("Nil")
}

func main() {
	ll := &LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)

	ll.Print()
	fmt.Println("Loop detected?", ll.DetectLoop())

	ll.CreateLoop(2)
	// ll.Print()
	fmt.Println("Loop detected after creating a loop?", ll.DetectLoop())
}
