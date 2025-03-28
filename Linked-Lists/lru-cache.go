package main

import "fmt"

// Doubly Linked List - Node Structure
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache structure
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

// Constructor Function
func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{}, // Dummy head
		tail:     &Node{}, // Dummy tail
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

// Get: Retrieves value and marks as recently used
func (lru *LRUCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.value
	}
	return -1 // Key not found
}

// Put: Adds/Updates key-value pair, removes LRU item if needed
func (lru *LRUCache) Put(key int, value int) {
	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToHead(node)
	} else {
		if len(lru.cache) >= lru.capacity {
			lru.removeLRU()
		}
		newNode := &Node{key: key, value: value}
		lru.cache[key] = newNode
		lru.addToHead(newNode)
	}
}

// moveToHead: Moves a node to the front (most recently used)
func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// addToHead: Adds a node right after the head
func (lru *LRUCache) addToHead(node *Node) {
	node.next = lru.head.next
	node.prev = lru.head
	lru.head.next.prev = node
	lru.head.next = node
}

// removeNode: Removes a node from the list
func (lru *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// removeLRU: Removes the least recently used item (tail.prev)
func (lru *LRUCache) removeLRU() {
	lruNode := lru.tail.prev
	delete(lru.cache, lruNode.key)
	lru.removeNode(lruNode)
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 10)
	lru.Put(2, 20)
	fmt.Println(lru.Get(1)) // Output: 10
	lru.Put(3, 30)          // Removes key 2
	fmt.Println(lru.Get(2)) // Output: -1 (not found)
	lru.Put(4, 40)          // Removes key 1
	fmt.Println(lru.Get(1)) // Output: -1 (not found)
	fmt.Println(lru.Get(3)) // Output: 30
	fmt.Println(lru.Get(4)) // Output: 40
}

// TC - O(1) && SC - O(1)
