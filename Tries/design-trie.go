package main

import "fmt"

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the whole Trie
type Trie struct {
	root *TrieNode
}

// Initialize the NewTrie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root

	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search a word into the Trie
func (t *Trie) Search(word string) bool {
	node := t.root

	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith checks if any word in the Trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}

// Delete a word from the Trie
func (t *Trie) Delete(word string) {
	var deleteHelper func(node *TrieNode, word string, depth int) bool
	deleteHelper = func(node *TrieNode, word string, depth int) bool {
		if depth == len(word) {
			if !node.isEnd {
				return false
			}
			node.isEnd = false
			return len(node.children) == 0
		}
		ch := rune(word[depth])
		child, exists := node.children[ch]
		if !exists {
			return false
		}
		shouldDelete := deleteHelper(child, word, depth+1)
		if shouldDelete {
			delete(node.children, ch)
			return len(node.children) == 0
		}
		return false
	}
	deleteHelper(t.root, word, 0)
}

// Prints all words in the Trie
func (t *Trie) Print() {
	var dfs func(node *TrieNode, prefix string)
	dfs = func(node *TrieNode, prefix string) {
		if node.isEnd {
			fmt.Println(prefix)
		}
		for ch, child := range node.children {
			dfs(child, prefix+string(ch))
		}
	}
	dfs(t.root, "")
}

func main() {
	trie := NewTrie()
	trie.Insert("hii")
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("helicopter")

	fmt.Println("All words in Trie:")
	trie.Print()
	fmt.Println("Search for 'hello':", trie.Search("hello"))
	fmt.Println("Start with 'he':", trie.StartsWith("he"))

	trie.Delete("hii")
	fmt.Println("\nAfter deleting 'hii':")
	trie.Print()
}
