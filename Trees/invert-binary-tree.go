package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Invert tree using recursively DFS
func invertTreeRecursively(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeRecursively(root.Left)
	invertTreeRecursively(root.Right)

	return root
}

// Invert tree using iterative DFS
func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	fmt.Print("Original Tree:")
	printTree(root)

	// invertTreeRecursively(root)
	invertTreeIterative(root)

	fmt.Print("\nInverted Tree:")
	printTree(root)
}
