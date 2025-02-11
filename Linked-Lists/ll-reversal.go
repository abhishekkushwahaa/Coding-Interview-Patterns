package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insertion
func (ll *LinkedList) Insert(value int) {
	NewNode := &Node{data: value}
	if ll.head == nil {
		ll.head = NewNode
		return
	}

	CurrentNode := ll.head
	for CurrentNode.next != nil {
		CurrentNode = CurrentNode.next
	}
	CurrentNode.next = NewNode
}

// Reversal
func (ll *LinkedList) Reverse() {
	var PrevNode *Node
	CurrentNode := ll.head

	for CurrentNode != nil {
		NextNode := CurrentNode.next
		CurrentNode.next = PrevNode
		PrevNode = CurrentNode
		CurrentNode = NextNode
	}
	ll.head = PrevNode
}

// Print LinkedList
func (ll *LinkedList) Print() {
	CurrentNode := ll.head
	for CurrentNode != nil {
		fmt.Print(CurrentNode.data, "->")
		CurrentNode = CurrentNode.next
	}
	fmt.Println("Nil")
}

func main() {
	ll := &LinkedList{}
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)
	ll.Insert(50)

	fmt.Println("Original Linked List:")
	ll.Print()

	ll.Reverse()
	fmt.Println("Reverse Linked List:")
	ll.Print()
}
