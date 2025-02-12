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

	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

// Logic of this Question
func (ll *LinkedList) RemoveKthFromEnd(k int) {
	// Step-1: Find the length of list
	length := 0
	currentNode := ll.head
	for currentNode != nil {
		length++
		currentNode = currentNode.next
	}

	// Step-2: If K is larger than length, do nothing
	if k > length || k <= 0 {
		fmt.Println("K is out of valid range!")
		return
	}

	// Step-3: If we need to remove head
	if k == length {
		ll.head = ll.head.next
		return
	}

	// Step-4: Traverse again to the (length-k-1)th node
	currentNode = ll.head
	for i := 0; i < length-k-1; i++ {
		currentNode = currentNode.next
	}

	// Remove Kth Last Node
	currentNode.next = currentNode.next.next
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
	ll.Insert(4)
	ll.Insert(7)
	ll.Insert(3)

	fmt.Println("Original Linked List:")
	ll.Print()

	k := 2
	ll.RemoveKthFromEnd(k)
	fmt.Printf("\nLinked List after removing %d-th last node:\n", k)
	ll.Print()
}
