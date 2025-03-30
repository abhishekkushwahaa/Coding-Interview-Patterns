package main

import "fmt"

// Linked List Node Structure
type Node struct {
	Val  int
	Next *Node
}

// TC - O(n) && SC - O(1)
// Reverse k nodes in the linked list
func reverseKGroup(head *Node, k int) *Node {
	if head == nil || k == 1 {
		return head
	}

	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}

	// Dummy node to handle head changes
	dummy := &Node{0, head}
	prevGroupEnd := dummy

	for count >= k {
		// Start reversing k nodes
		curr := prevGroupEnd.Next
		next := curr.Next

		for i := 1; i < k; i++ {
			curr.Next = next.Next
			next.Next = prevGroupEnd.Next
			prevGroupEnd.Next = next
			next = curr.Next
		}

		prevGroupEnd = curr
		count -= k
	}

	return dummy.Next
}

// Create a Linked List
func createLinkedList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	head := &Node{values[0], nil}
	curr := head
	for _, val := range values[1:] {
		curr.Next = &Node{val, nil}
		curr = curr.Next
	}
	return head
}

// Print Linked List
func printLinkedList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	k := 3
	fmt.Println("Original List:")
	printLinkedList(head)

	newHead := reverseKGroup(head, k)
	fmt.Println("Reversed in k-Groups:")
	printLinkedList(newHead)
}
