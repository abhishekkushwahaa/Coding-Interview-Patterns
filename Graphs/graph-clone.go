package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

// TC - O(n + e) && SC - O(n)
func cloneGraphDFS(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if clonedNode, found := visited[node]; found {
		return clonedNode
	}

	clone := &Node{Val: node.Val}
	visited[node] = clone

	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneGraphDFS(neighbor, visited))
	}

	return clone
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return cloneGraphDFS(node, visited)
}

// DFS function to print the graph
func printGraphDFS(node *Node, visited map[*Node]bool) {
	if node == nil || visited[node] {
		return
	}

	visited[node] = true
	fmt.Printf("Node %d -> ", node.Val)
	for _, neighbor := range node.Neighbors {
		fmt.Printf("%d ", neighbor.Val)
	}
	fmt.Println()

	for _, neighbor := range node.Neighbors {
		printGraphDFS(neighbor, visited)
	}
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	fmt.Println("Original Graph:")
	visited := make(map[*Node]bool)
	printGraphDFS(node1, visited)

	clonedGraph := cloneGraph(node1)

	fmt.Println("\nCloned Graph:")
	visited = make(map[*Node]bool)
	printGraphDFS(clonedGraph, visited)
}
