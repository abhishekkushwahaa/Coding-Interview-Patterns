package main

import "fmt"

// Define linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge Sort for Linked List - O(n log n) time & O(log n) space (for recursion stack)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Find the middle of the list
	mid := getMiddle(head)
	rightHead := mid.Next
	mid.Next = nil // Split the list

	// Recursively sort both half
	left := sortList(head)
	right := sortList(rightHead)

	return merge(left, right)
}

// Find the middle node (Slow & Fast Pointer)
func getMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Merge two sorted lists
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}

// Print the linked list
func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}

	fmt.Println("Original Linked List:")
	print(head)

	sortedHead := sortList(head)

	fmt.Println("Sorted Linked List:")
	print(sortedHead)
}

// TC - O(nlogn) (Using Merge Sort)
// SC - O(logn) (Recursive Stack)
