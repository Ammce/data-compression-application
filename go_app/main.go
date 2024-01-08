package main

import (
	"container/heap"
	"fmt"
	"strings"
	"time"
)

// Node represents a node in the Huffman tree
type Node struct {
	Value     rune
	Frequency int
	Left      *Node
	Right     *Node
}

// HuffmanHeap is a min-heap of Nodes used to build the Huffman tree
type HuffmanHeap []*Node

func (h HuffmanHeap) Len() int           { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency }
func (h HuffmanHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[0 : n-1]
	return node
}

// BuildHuffmanTree builds the Huffman tree and returns the root node
func BuildHuffmanTree(freqMap map[rune]int) *Node {
	priorityQueue := &HuffmanHeap{}
	heap.Init(priorityQueue)

	for char, freq := range freqMap {
		node := &Node{Value: char, Frequency: freq}
		heap.Push(priorityQueue, node)
	}

	for priorityQueue.Len() > 1 {
		left := heap.Pop(priorityQueue).(*Node)
		right := heap.Pop(priorityQueue).(*Node)
		internalNode := &Node{Frequency: left.Frequency + right.Frequency, Left: left, Right: right}
		heap.Push(priorityQueue, internalNode)
	}

	return (*priorityQueue)[0]
}

// GenerateHuffmanCodes generates Huffman codes for each character in the tree
func GenerateHuffmanCodes(root *Node, code string, codes map[rune]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		codes[root.Value] = code
	}

	GenerateHuffmanCodes(root.Left, code+"0", codes)
	GenerateHuffmanCodes(root.Right, code+"1", codes)
}

// CompressMessage compresses a message using Huffman codes
func CompressMessage(message string, codes map[rune]string) string {
	var compressedMessage strings.Builder

	for _, char := range message {
		compressedMessage.WriteString(codes[char])
	}

	return compressedMessage.String()
}

func main() {
	text := "BCCABBDDAECCBBAEDDCC"

	// Count character frequencies
	freqMap := make(map[rune]int)
	for _, char := range text {
		freqMap[char]++
	}

	// Build Huffman tree
	root := BuildHuffmanTree(freqMap)

	// Generate Huffman codes
	codes := make(map[rune]string)
	GenerateHuffmanCodes(root, "", codes)

	// Compress the message
	startTime := time.Now()
	compressedMessage := CompressMessage(text, codes)

	endTime := time.Now()

	executionTime := endTime.Sub(startTime).Microseconds()

	fmt.Println("Execution Time:", executionTime)

	// Print Huffman codes
	fmt.Println("Huffman Codes:")
	for char, code := range codes {
		fmt.Printf("%c: %s\n", char, code)
	}

	fmt.Println("\nOriginal Message:", text)
	fmt.Println("Compressed Message:", compressedMessage)
}
