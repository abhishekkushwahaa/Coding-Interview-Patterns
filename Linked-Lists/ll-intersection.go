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
func GetIntersectionNode(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}

	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		ptrA = ptrA.next
		ptrB = ptrB.next
		if ptrA == nil && ptrB != nil {
			ptrA = headB
		}
		if ptrB == nil && ptrA != nil {
			ptrB = headA
		}
	}
	return ptrA
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
	list1 := &LinkedList{}
	list1.Insert(1)
	list1.Insert(2)
	list1.Insert(3)

	list2 := &LinkedList{}
	list2.Insert(4)
	list2.Insert(5)

	intersectingNode := &Node{data: 6}
	list1.head.next.next.next = intersectingNode // 3 → 6

	list2.head.next.next = intersectingNode // 5 → 6

	intersectingNode.next = &Node{data: 7}
	intersectingNode.next.next = &Node{data: 8}

	fmt.Println("List 1:")
	list1.Print()

	fmt.Println("List 2:")
	list2.Print()

	intersection := GetIntersectionNode(list1.head, list2.head)
	if intersection != nil {
		fmt.Println("Intersection at node with value:", intersection.data)
	} else {
		fmt.Println("No intersection found.")
	}

}
