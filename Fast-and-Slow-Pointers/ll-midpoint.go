package main

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	Value int
	Next  *Node
}

// FindMidpoint finds the middle node of the linked list
// TC - 0(n) && SC - 0(1)
func FindMidpoint(head *Node) *Node {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// Create a linked list and return its
func createLinkedList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{Value: values[0]}
	current := head

	for _, v := range values[1:] {
		current.Next = &Node{Value: v}
		current = current.Next
	}

	return head
}

// Print linked list
func printLinkedList(head *Node) {
	for head != nil {
		fmt.Print(head.Value, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

// Main function
func main() {
	values := []int{1, 2, 3, 4, 5}
	head := createLinkedList(values)
	fmt.Print("Linked List: ")
	printLinkedList(head)

	mid := FindMidpoint(head)
	fmt.Println("Midpoint:", mid.Value)
}
