package main

import (
	"container/heap"
	"fmt"
)

// WordFreq stores word and its frequency
type WordFreq struct {
	word string
	freq int
}

// MaxHeap for sorting based count
type MaxHeap []WordFreq

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq } // Max Heap (higher freq first)

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(WordFreq))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequentMaxHeap(words []string, k int) []string {
	// Step-1: Count the Frequency
	freqCount := make(map[string]int)
	for _, word := range words {
		freqCount[word]++
	}

	// Step-2: Add elements to Max Heap
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for word, count := range freqCount {
		heap.Push(maxHeap, WordFreq{word, count})
	}

	// Step 3: Extract top k elements
	var result []string
	for i := 0; i < k && maxHeap.Len() > 0; i++ {
		result = append(result, heap.Pop(maxHeap).(WordFreq).word)
	}
	return result
}

// MinHeap for maintaining the top k frequent words
type MinHeap []WordFreq

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq } // Min Heap (lower freq first)
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(WordFreq))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequentMinHeap(words []string, k int) []string {
	// Step 1: Count frequencies
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}

	// Step 2: Add elements to Min Heap (size k)
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for word, count := range freqMap {
		heap.Push(minHeap, WordFreq{word, count})
		if minHeap.Len() > k {
			heap.Pop(minHeap) // Remove smallest frequency
		}
	}

	// Step 3: Extract top k elements
	var result []string
	for minHeap.Len() > 0 {
		result = append(result, heap.Pop(minHeap).(WordFreq).word)
	}
	return result
}

func main() {
	words := []string{"apple", "banana", "apple", "orange", "banana", "banana", "grape", "apple", "grape", "grape", "grape"}
	k := 2

	fmt.Println("Top K frequent using MaxHeap:", topKFrequentMaxHeap(words, k)) // TC = O(NlogN)
	fmt.Println("Top K frequent using MinHeap:", topKFrequentMinHeap(words, k)) // TC = O(NlogK)
}
